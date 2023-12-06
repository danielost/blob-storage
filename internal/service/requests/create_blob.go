package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	. "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/dl7850949/blob-storage/resources"
)

type CreateBlobRequest struct {
	Data resources.CreateBlob
}

func NewCreateBlobRequest(r *http.Request) (CreateBlobRequest, error) {
	var request CreateBlobRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return CreateBlobRequest{}, errors.Wrap(err, "failed to unmarshal")
	}

	return request, validateCreateBlobRequest(request)
}

func validateCreateBlobRequest(request CreateBlobRequest) error {
	attrs := &request.Data.Attributes

	return ValidateStruct(attrs,
		Field(&attrs.Value, Required),
	)
}
