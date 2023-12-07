package helpers

import (
	"net/http"

	"github.com/google/jsonapi"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/data"
)

func GetBlobById(r *http.Request, blobId int64) (*data.Blob, *jsonapi.ErrorObject) {
	blob, err := BlobsQ(r).FilterByID(blobId).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob from DB")
		return blob, problems.InternalError()
	}

	if blob == nil {
		Log(r).Info("blob not found")
		return blob, problems.NotFound()
	}

	return blob, nil
}
