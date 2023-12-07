package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewAuthRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	existingUser, ok := getUserByLogin(w, r, request.Data.Attributes.Login)
	if !ok {
		return
	}

	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to fetch user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if existingUser != nil {
		helpers.Log(r).WithError(err).Warn("login is taken")
		ape.RenderErr(w, problems.Conflict())
		return
	}

	saltedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Data.Attributes.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to hash password")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	user := data.User{
		CreatedAt: time.Now().UTC(),
		Login:     request.Data.Attributes.Login,
		Password:  string(saltedPassword),
	}

	resultUser, err := helpers.UsersQ(r).Insert(user)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to insert user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := resources.UserResponse{
		Data: newUserModel(*resultUser),
	}

	ape.Render(w, response)
}

func newUserModel(user data.User) resources.User {
	return resources.User{
		Key: resources.NewKeyInt64(user.ID, resources.USER),
		Attributes: resources.UserAttributes{
			CreatedAt: user.CreatedAt,
			Login:     user.Login,
			Password:  user.Password,
		},
	}
}
