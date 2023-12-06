package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := ParseJWT(r)
		if err != nil || !token.Valid {
			ape.RenderErr(w, problems.Unauthorized())
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ParseJWT(r *http.Request) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	tokenString := r.Header.Get("x-jwt-token")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func GetIdFromJWT(r *http.Request) (int64, error) {
	token, err := ParseJWT(r)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse JWT")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.Wrap(err, "failed to retrive JWT claims")
	}

	userId, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.Wrap(err, "JWT claims doesn't contain id")
	}

	return int64(userId), nil
}
