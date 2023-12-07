package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	userId, err := middleware.GetIdFromJWT(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get id from JWT")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	blob := data.Blob{
		CreatedAt: time.Now().UTC(),
		Value:     request.Data.Attributes.Value,
		OwnerId:   userId,
	}

	resultBlob, err := helpers.BlobsQ(r).Insert(blob)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to insert blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := resources.BlobResponse{
		Data: newBlobModel(resultBlob, userId),
	}

	ape.Render(w, response)
}

func newBlobModel(blob *data.Blob, userId int64) resources.Blob {
	relationshipKey := resources.NewKeyInt64(userId, resources.USER)
	return resources.Blob{
		Key: resources.NewKeyInt64(blob.ID, resources.BLOB),
		Attributes: resources.BlobAttributes{
			CreatedAt: blob.CreatedAt,
			Value:     blob.Value,
		},
		Relationships: resources.BlobRelationships{
			Owner: resources.Relation{
				Data: &relationshipKey,
			},
		},
	}
}
