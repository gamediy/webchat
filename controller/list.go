package controller

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func List(w http.ResponseWriter, request *http.Request) error {

	tokenString:=request.Header.Get("x-token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("rJXUCzdnN8mf"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"], claims["exp"])
	} else {
		fmt.Println(err)
	}


	return nil
}