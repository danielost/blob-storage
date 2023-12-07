package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewAuthRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	user, ok := getUserByLogin(w, r, request.Data.Attributes.Login)
	if !ok {
		return
	}

	if user == nil {
		helpers.Log(r).WithError(err).Warn("user not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Data.Attributes.Password)) != nil {
		helpers.Log(r).WithError(err).Warn("invalid password")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	token, err := generateJWT(*user)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to generate JWT")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := resources.TokenResponse{
		Data: resources.Token{
			Key: resources.NewKeyInt64(0, resources.BEARER_TOKEN),
			Attributes: resources.TokenAttributes{
				Token: token,
			},
		},
	}

	ape.Render(w, response)
}

func generateJWT(user data.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":     user.Login,
		"id":        user.ID,
		"expiresAt": time.Now().Add(time.Minute * 5).UTC(),
	})
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

func getUserByLogin(w http.ResponseWriter, r *http.Request, login string) (*data.User, bool) {
	user, err := helpers.UsersQ(r).FilterByLogin(login).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch user")
		ape.RenderErr(w, problems.InternalError())
		return nil, false
	}

	return user, true
}
