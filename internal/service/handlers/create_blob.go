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
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("wrong request")
		ape.RenderErr(w, problems.BadRequest(err)...)
	}

	var resultBlob *data.Blob

	err = BlobsQ(r).Transaction(func(q data.BlobsQ) error {
		blob := data.Blob{
			CreatedAt: time.Now().UTC(),
			Value:     request.Data.Value,
		}
		resultBlob, err = q.Insert(blob)
		if err != nil {
			return errors.Wrap(err, "failed to insert blob")
		}
		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	result := resources.BlobResponse{
		Data: newBlobModel(*resultBlob),
	}
	ape.Render(w, result)
}

func newBlobModel(blob data.Blob) resources.Blob {
	return resources.Blob{
		Key: resources.NewKeyInt64(blob.ID, resources.BLOB),
		Attributes: resources.BlobAttributes{
			CreatedAt: blob.CreatedAt,
			Value:     blob.Value,
		},
	}
}
