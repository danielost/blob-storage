package handlers

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewAuthRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
	}

	user, err := UsersQ(r).FilterByLogin(request.Data.Attributes.Login).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to fetch user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if user == nil {
		Log(r).WithError(err).Error("user not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Data.Attributes.Password)) != nil {
		Log(r).WithError(err).Error("invalid password")
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	token, err := generateJWT(*user)
	if err != nil {
		Log(r).WithError(err).Error("failed to generate JWT")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.TokenResponse{
		Data: resources.Token{
			Key: resources.NewKeyInt64(0, resources.BEARER_TOKEN),
			Attributes: resources.TokenAttributes{
				Token: token,
			},
		},
	}

	ape.Render(w, result)
}

func generateJWT(user data.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login":     user.Login,
		"id":        user.ID,
		"expiresAt": 15000,
	})
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}
