package requests

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
)

type GetBlobRequest struct {
	BlobID int64 `url:"-"`
}

func NewGetBlobRequest(r *http.Request) (GetBlobRequest, error) {
	request := GetBlobRequest{}
	request.BlobID = cast.ToInt64(chi.URLParam(r, "id"))

	return request, nil
}
