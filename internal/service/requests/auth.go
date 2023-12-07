package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	. "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/dl7850949/blob-storage/resources"
)

type AuthRequest struct {
	Data resources.UserData
}

func NewAuthRequest(r *http.Request) (AuthRequest, error) {
	var request AuthRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, validateAuthRequest(request)
}

func validateAuthRequest(request AuthRequest) error {
	attrs := &request.Data.Attributes

	return ValidateStruct(attrs,
		Field(&attrs.Login, Required),
		Field(&attrs.Password, Required),
	)
}
