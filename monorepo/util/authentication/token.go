package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

	return errors.New("invalid token")

}
func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)
	if err != nil {
		return 0, err
	}

	if permission, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permission["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}
	return 0, errors.New("invalid token")
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
