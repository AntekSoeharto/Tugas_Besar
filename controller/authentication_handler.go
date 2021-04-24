package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("bebasapasaja")
var tokenName = "token"

type Claims struct {
	ID       int    `json:id`
	Name     string `json:name`
	UserType int    `json:user_type`
	jwt.StandardClaims
}

func generateToken(w http.ResponseWriter, id int, name string, userType int) {
	tokenExpiryTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		ID:       id,
		Name:     name,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    signedToken,
		Expires:  tokenExpiryTime,
		Secure:   false,
		HttpOnly: true,
	})
}

func resetUsersToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     tokenName,
		Value:    "",
		Expires:  time.Now(),
		Secure:   false,
		HttpOnly: true,
	})
}

func Authenticate(next http.HandlerFunc, accessType int) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidToken := validateUserToken(w, r, accessType)
		fmt.Println(accessType)
		if !isValidToken {
			sendUnAuthorizedResponse(w)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func validateUserToken(w http.ResponseWriter, r *http.Request, accessType int) bool {
	isAccessTokenValid, id, email, userType := validateTokenFromCookies(r)
	fmt.Println(id, email, userType, accessType, isAccessTokenValid)

	if isAccessTokenValid {
		isUserValid := userType == accessType
		if isUserValid {
			return true
		}
	}
	return false
}

func validateTokenFromCookies(r *http.Request) (bool, int, string, int) {
	if coockie, err := r.Cookie(tokenName); err == nil {
		accessToken := coockie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return true, accessClaims.ID, accessClaims.Name, accessClaims.UserType
		}
	}
	return false, -1, "", -1
}

func getid(r *http.Request) int {
	if coockie, err := r.Cookie(tokenName); err == nil {
		accessToken := coockie.Value
		accessClaims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(accessToken, accessClaims, func(accessToken *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err == nil && parsedToken.Valid {
			return accessClaims.ID
		}
	}
	return -1
}
