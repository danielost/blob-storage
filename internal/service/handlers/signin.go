package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
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

	user, jsonErr := helpers.GetUserByLogin(r, request.Data.Attributes.Login)
	if jsonErr != nil {
		ape.RenderErr(w, jsonErr)
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

	token, err := middleware.GenerateJWT(*user)
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
