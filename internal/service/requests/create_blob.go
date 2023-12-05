package requests

import (
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"

	"gitlab.com/dl7850949/blob-storage/resources"
)

type CreateBlobRequest struct {
	Data resources.CreateBlob
}

func NewCreateBlobRequest(r *http.Request) (*CreateBlobRequest, error) {
	request := resources.CreateBlob{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}

	return &CreateBlobRequest{request}, nil
}
