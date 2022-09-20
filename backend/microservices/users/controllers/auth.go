package controllers

import (
	//"encoding/json"
	"encoding/json"
	"errors"
	"govueadmin/framework/cookie"
	"govueadmin/framework/jwt"
	"govueadmin/framework/request"
	"govueadmin/framework/response"
	"govueadmin/microservices/users/models"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	JWT "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
func Login(w http.ResponseWriter, r *http.Request) {
	// set default http 200 status
	status := http.StatusOK
	var message string

	req, err := request.Factory(r)

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

		if err = models.FindByEmail(&user, email); err != nil && errors.Is(err, mongo.ErrNoDocuments) { // fetch the user information by the email address
			status = http.StatusUnauthorized
			message = LOGIN_EMAIL_DONT_EXIST
		} else {
			if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil { // match the request password with the user's saved password
				status = http.StatusUnauthorized
				message = LOGIN_WRONG_PASSWORD
			} else {
				uid := user.GetHaxUserId()
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
					w.Header().Set("Authorization", jwtToken)
					// generate an encoded cookie string for the jwt token
					encoded, err := cookie.EncodeCookieString(&jwtToken, "auth-token", os.Getenv("SECRET_KEY_COOKIE"))

					if err != nil {
						status = http.StatusInternalServerError
						message = err.Error()
					} else {

						// save encoded jwt token in http cookie
						message = LOGIN_SUCCESS
						cookie := &http.Cookie{
							Name:     "auth-token",
							Value:    encoded,
							Path:     "/",
							HttpOnly: true,
						}

						// save user
						http.SetCookie(w, cookie)
					}
				}
			}
		}
	}

	// send response back to the client
	response.Json(&w, status, map[string]string{
		"message": message,
	})
}

func GetJwtClaims(r *http.Request) JWT.MapClaims {
	var token string
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		// Extract auth-token cookie from request
		authCookie, err := r.Cookie("auth-token")

		if err == nil {
			mu.Lock()
			cookie.DecodeCookieString(authCookie.Value, "auth-token", &token, os.Getenv("SECRET_KEY_COOKIE"))
			mu.Unlock()
		}

		wg.Done()
	}()

	go func() {
		mu.Lock()
		token = r.Header.Get("Authorization")
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()

	if token == "" {
		return nil
	}

	jwtToken, _ := jwt.GetOriginalToken(&token)

	return jwt.GetJWTClaims(jwtToken)
}

// check if user is logged or not
// this call should pass through the middleware
// if middleware okays request then it should
// just send details of the logged in user
func Me(w http.ResponseWriter, r *http.Request) {
	// fetch the claims of the user
	claims := GetJwtClaims(r)

	user := &models.User{}

	objectId, err := user.GetObjectIdFromHex(claims["user_id"].(string))

	if err != nil {
		response.Json(&w, http.StatusUnauthorized, map[string]string{
			"message": http.StatusText(http.StatusUnauthorized),
		})
		return
	}

	user.Find(objectId)

	response.Json(&w, http.StatusOK, user)
}

// logout user
func Logout(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	wg.Add(2)

	// remove jwt token from cookie
	go func() {
		// fetch auth cookie from request
		cookie := &http.Cookie{
			Name:     "auth-cookie",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Unix(0, 0),
		}

		// set cookie in http response
		http.SetCookie(w, cookie)
		wg.Done()
	}()

	// remove jwt token from authorization header
	go func() {
		claims := GetJwtClaims(r)

		if claims != nil {
			authToken := &models.AuthToken{
				UserId:  claims["user_id"].(string),
				TokenId: claims["token_id"].(string),
			}

			authToken.Delete()
		}
		wg.Done()
	}()

	wg.Wait()

	response.Json(&w, http.StatusOK, map[string]string{
		"message": LOGOUT_SUCCESS,
	})

}
