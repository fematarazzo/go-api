package authentication

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"api/src/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userID

	tokenString, error := token.SignedString([]byte(config.SecretKey))
	if error != nil {
		return "", error
	}

	return tokenString, nil
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, error := jwt.Parse(tokenString, returnVerificationKey)

	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}

func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnVerificationKey)
	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if error != nil {
			return 0, error
		}

		return userID, nil
	}

	return 0, errors.New("Invalid Token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 && strings.Split(token, " ")[0] == "Bearer" {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
