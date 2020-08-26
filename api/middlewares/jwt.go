package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/leomirandadev/golang-api/services/httpResponse"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

type Output struct {
	TOKEN string `json:"token"`
}

func JwtVerify(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] != nil {
			bearerSplited := strings.Split(r.Header["Authorization"][0], " ")

			if len(bearerSplited) == 2 {
				bearerToken := bearerSplited[1]

				token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return mySigningKey, nil
				})

				if err != nil {
					httpResponse.RenderError(w, "UNAUTHORIZED", http.StatusUnauthorized)
				}

				if token != nil {
					if token.Valid {
						endpoint(w, r)
					}
				}
			} else {
				httpResponse.RenderError(w, "INVALID_AUTHORIZATION", http.StatusUnauthorized)
			}

		} else {
			httpResponse.RenderError(w, "WITHOUT_AUTHORIZATION", http.StatusUnauthorized)
		}
	})
}

func GenerateJwt(data interface{}) (Output, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	tokenOutput := Output{TOKEN: ""}

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return tokenOutput, err
	}
	tokenOutput.TOKEN = tokenString
	return tokenOutput, nil
}
