package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
)

var jwtSecret string

func ValidateJWT(secret string) func(http.Handler) http.Handler {
	jwtSecret = secret

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := ParseJWT(r)
			if err != nil || !token.Valid {
				helpers.Log(r).WithError(err).Warn("token is invalid")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				helpers.Log(r).Error("couldn't process jwt claims")
				ape.RenderErr(w, problems.InternalError())
				return
			}

			exp, ok := claims["expiresAt"]
			if !ok {
				helpers.Log(r).Warn("no expiration date in jwt claims")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			layout := "2006-01-02 15:04:05 -0700 MST"
			expTime, err := time.Parse(layout, exp.(string))
			if err != nil {
				helpers.Log(r).WithError(err).Warn("error parsing expiration date of jwt")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			if time.Now().After(expTime) {
				helpers.Log(r).Warn("jwt expired")
				ape.RenderErr(w, problems.Unauthorized())
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func ParseJWT(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")[len("Bearer "):]

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
}

func GenerateJWT(user data.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":     user.Login,
		"id":        user.ID,
		"expiresAt": time.Now().Add(time.Minute * 5).UTC().String(),
	})
	return token.SignedString([]byte(jwtSecret))
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
