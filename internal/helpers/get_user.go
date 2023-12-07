package helpers

import (
	"net/http"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/data"
)

func GetUserByLogin(r *http.Request, login string) (*data.User, *jsonapi.ErrorObject) {
	user, err := UsersQ(r).FilterByLogin(login).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to fetch user")
		return nil, problems.InternalError()
	}

	return user, nil
}
