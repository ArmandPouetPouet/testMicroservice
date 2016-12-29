package main

import (
	"net/http"

	"testMicroservice/Data"

	jwt "github.com/dgrijalva/jwt-go"
)

/* Set up a global string for our secret */
var mySigningKey = []byte("KillTeamRepresent")

//CheckTokenInCookieHandler handler to get token
func CheckTokenInCookieHandler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, cookie := range r.Cookies() {
			if checkToken(*cookie) {
				//User has the right claims
				BuildJSONResponse(w, 200)
				inner.ServeHTTP(w, r)
				break
			}
		}
		BuildJSONResponse(w, 403)
	})
}

func checkToken(cookie http.Cookie) bool {
	if cookie.Name == "security" {
		token, _ := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		return token.Valid
	}
	return false
}

//TODO : Find a way to successfully parse Standard claims ! WTF ???
func checkClaims(tokenString string) bool {

	/*
		token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

		if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
			fmt.Printf("%v", claims.ExpiresAt)
		} else {
			fmt.Println(err)
		}
	*/
	return false
}

//CheckCredentials return true if a user was found with these credentials
func CheckCredentials(w http.ResponseWriter, r *http.Request, login, pwd string) bool {
	user := Data.LoginUser(login, pwd)
	if user.Firstname != "" {
		setToken(w, r, user)
		return true
	}
	return false
}

func setToken(w http.ResponseWriter, r *http.Request, user Data.User) {
	/*
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "testMicroserviceAPI",
			Subject:   "Admin",
			Id:        user.Login,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	*/
	token := jwt.New(jwt.SigningMethodHS256)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		panic(err)
	}

	cookie := http.Cookie{Name: "security", Value: ss, Domain: "localhost", Path: "/"}
	http.SetCookie(w, &cookie)
}
