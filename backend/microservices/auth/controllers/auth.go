package controllers

import (
	//"encoding/json"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"govueadmin/framework/cookie"
	"govueadmin/framework/jwt"
	"govueadmin/framework/request"
	"govueadmin/framework/response"
	"govueadmin/microservices/auth/pb"

	//"govueadmin/framework/response"
	"govueadmin/microservices/auth/models"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	JWT "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	//"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string
	Password string
}

var lr LoginRequest
var (
	LOGIN_FAILED                   = "Email or password is incorrect"
	LOGIN_INTERNAL_ERR_JSON_DECODE = "Invalid request body"
	LOGIN_EMAIL_DONT_EXIST         = "Account doesn't exist with the email specfied, please create an account then try again to login."
	LOGIN_WRONG_PASSWORD           = "You have provided wrong password"
	LOGIN_CANT_GEN_TOKEN           = "Unable to generate jwt token"
	LOGIN_SUCCESS                  = "You have successfully logged in"
	LOGOUT_SUCCESS                 = "You have successfully logged out"
)

/*
  Parse the login request into struct
*/
func ParseLoginRequest(body io.ReadCloser) error {

	return json.NewDecoder(body).Decode(&lr)
}

/*
  Set JWT claims options of user model
*/
func SetJwtClaims(fields ...string) *JWT.MapClaims {
	return &JWT.MapClaims{
		"user_id":  fields[0],
		"token_id": fields[1],
		"expires":  time.Now().Add(time.Minute * 60),
	}
}

/*
  Authenticate user by email & password and returns
  jwt token in cookie response.
*/
func LoginHandler(c *gin.Context) {
	// set default http 200 status
	status := http.StatusOK
	var message string

	req, err := request.Factory(c.Request)

	// decode json body and populate LoginRequest
	//err := ParseLoginRequest(r.Body)

	if err != nil {
		status = http.StatusInternalServerError
		message = LOGIN_INTERNAL_ERR_JSON_DECODE
	} else {
		req.Parse()

		email := req.Get("Email")
		password := req.Get("Password")

		var user models.User

		if err = user.GetByEmail(email); err != nil && errors.Is(err, gorm.ErrRecordNotFound) { // fetch the user information by the email address
			status = http.StatusUnauthorized
			message = LOGIN_EMAIL_DONT_EXIST
		} else {
			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // match the request password with the user's saved password
				status = http.StatusUnauthorized
				message = LOGIN_WRONG_PASSWORD
			} else {
				uid := user.Uid
				tkn, _ := uuid.NewRandom()

				ch := make(chan bool)

				// Save auth tokens in database
				go func(userId string, ch chan bool) {
					authToken := &models.AuthToken{
						UserId:  uid,
						TokenId: tkn.String(),
					}

					err1 := authToken.Create()

					if err1 != nil {
						ch <- false
					} else {
						ch <- true
					}
				}(uid, ch)

				// create jwt token claims map
				claims := SetJwtClaims(uid, tkn.String())

				// generate jwt token from the claims map
				jwtToken, err := jwt.GenerateJWTToken(claims, os.Getenv("SECRET_KEY"))

				tokenId := <-ch

				if err != nil || !tokenId {
					status = http.StatusInternalServerError
					message = LOGIN_CANT_GEN_TOKEN
				} else {
					c.Header("Authorization", jwtToken)
					// generate an encoded cookie string for the jwt token
					encoded, err := cookie.EncodeCookieString(&jwtToken, "auth-token", os.Getenv("SECRET_KEY_COOKIE"))

					if err != nil {
						status = http.StatusInternalServerError
						message = err.Error()
					} else {

						// save encoded jwt token in http cookie
						message = LOGIN_SUCCESS
						// cookie := &http.Cookie{
						// 	Name:     "auth-token",
						// 	Value:    encoded,
						// 	Path:     "/",
						// 	HttpOnly: true,
						// }

						// save user
						c.SetCookie("auth-token", encoded, 0, "/", "/", false, true)
					}
				}
			}
		}
	}

	// send response back to the client
	c.JSON(status, gin.H{
		"message": message,
	})
}

func ExtractToken(c *gin.Context) string {
	var token string
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		// Extract auth-token cookie from request
		authCookie, err := c.Cookie("auth-token")

		if err == nil {
			mu.Lock()
			cookie.DecodeCookieString(authCookie, "auth-token", &token, os.Getenv("SECRET_KEY_COOKIE"))
			mu.Unlock()
		}

		wg.Done()
	}()

	go func() {
		mu.Lock()
		token = c.GetHeader("Authorization")
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()

	return token
}

func GetJwtClaims(token string) JWT.MapClaims {

	jwtToken, _ := jwt.GetOriginalToken(&token)

	return jwt.GetJWTClaims(jwtToken)
}

// check if user is logged or not
// this call should pass through the middleware
// if middleware okays request then it should
// just send details of the logged in user
func MeHandler(c *gin.Context, pc pb.AuthServiceClient) {
	// parse gin context parameters such as auth cookie, authorization headers
	// to Me Request and that MeRequest should be passed to subsequent calls
	res, err := pc.Me(context.Background(), &pb.MeRequest{
		Token: ExtractToken(c),
	})

	fmt.Println("Error:", err)
	if err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, res.GetMessage(), res.GetData())
}

func (s *GrpcServer) Me(c context.Context, mr *pb.MeRequest) (*pb.MeResponse, error) {
	fmt.Println("I am calling from grpc api call")
	// fetch the claims of the user
	claims := GetJwtClaims(mr.Token)

	user := &models.User{}

	err := user.Find(claims["user_id"].(string))

	if err != nil {
		return &pb.MeResponse{
			Status:  http.StatusUnauthorized,
			Message: http.StatusText(http.StatusUnauthorized),
			Data:    nil,
		}, nil
	}

	data := &pb.User{
		Uid:        user.Uid,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Phone:      user.Phone,
		Address:    user.Address,
		IsAdmin:    user.IsAdmin,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
	}

	return &pb.MeResponse{
		Status:  http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    data,
	}, nil
}

// logout user
func LogoutHandler(c *gin.Context) {
	var wg sync.WaitGroup

	wg.Add(2)

	// remove jwt token from cookie
	go func() {
		// // fetch auth cookie from request
		// cookie := &http.Cookie{
		// 	Name:     "auth-cookie",
		// 	Value:    "",
		// 	Path:     "/",
		// 	HttpOnly: true,
		// 	Expires:  time.Unix(0, 0),
		// }

		// set cookie in http response
		c.SetCookie("auth-cookie", "", 0, "/", "/", false, true)
		wg.Done()
	}()

	// remove jwt token from authorization header
	go func() {
		claims := GetJwtClaims(ExtractToken(c))

		if claims != nil {
			authToken := &models.AuthToken{}

			authToken.Delete(claims["user_id"].(string), claims["token_id"].(string))
		}
		wg.Done()
	}()

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"message": LOGOUT_SUCCESS,
	})

}
