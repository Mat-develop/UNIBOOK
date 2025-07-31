package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	config "v1/monorepo/util/route_config"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["expire"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey)) // secret key
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return err
	}
	fmt.Println(token)

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalid")

}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
