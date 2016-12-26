package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/* Set up a global string for our secret */
var mySigningKey = []byte("KillTeamRepresent")

//GetTokenHandler handler to get token
func GetTokenHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* Create the token */

		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "testMicroserviceAPI",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			panic(err)
		}
		/* Finally, write the token to the browser window */
		cookie := http.Cookie{Name: "pouet", Value: ss}
		http.SetCookie(w, &cookie)

		inner.ServeHTTP(w, r)
	})
}
