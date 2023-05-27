package controllers

import (
	//"encoding/json"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"app/microservices/auth/pb"
	"app/microservices/common/feat"

	"github.com/Keyur1991/go-shreeva/cookie"
	"github.com/Keyur1991/go-shreeva/jwt"
	"github.com/Keyur1991/go-shreeva/request"
	"github.com/Keyur1991/go-shreeva/response"

	//"github.com/Keyur1991/go-shreeva/response"
	"app/microservices/auth/models"
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
	LOGIN_FAILED                   = "login_failed"
	LOGIN_INTERNAL_ERR_JSON_DECODE = "login_internal_err_json_decode"
	LOGIN_EMAIL_DONT_EXIST         = "login_email_dont_exist"
	LOGIN_WRONG_PASSWORD           = "login_wrong_password"
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

func GetJwtClaims(token string) JWT.MapClaims {

	jwtToken, _ := jwt.GetOriginalToken(&token)

	return jwt.GetJWTClaims(jwtToken)
}

// this is a grpc service client method
// which invokes grpc service server method
// returns the response in json format
func MeHandler(c *gin.Context, pc pb.AuthServiceClient) {
	// parse gin context parameters such as auth cookie, authorization headers
	// to Me Request and that MeRequest should be passed to subsequent calls
	res, err := pc.Me(context.Background(), &pb.MeRequest{
		Token: feat.ExtractToken(c),
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
		claims := GetJwtClaims(feat.ExtractToken(c))

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
