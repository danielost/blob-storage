package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewAuthRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
	}

	var resultUser *data.User

	err = UsersQ(r).Transaction(func(q data.UsersQ) error {
		existingUser, err := q.FilterByLogin(request.Data.Attributes.Login).Get()
		if err != nil {
			return errors.Wrap(err, "failed to fetch user by login")
		}

		if existingUser != nil {
			return errors.Wrap(err, "login is already taken")
		}

		saltedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Data.Attributes.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.Wrap(err, "failed to hash password")
		}

		user := data.User{
			CreatedAt: time.Now().UTC(),
			Login:     request.Data.Attributes.Login,
			Password:  string(saltedPassword),
		}
		
		resultUser, err = q.Insert(user)
		if err != nil {
			return errors.Wrap(err, "failed to insert user")
		}
		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create user")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.UserResponse{
		Data: newUserModel(*resultUser),
	}
	ape.Render(w, result)
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
