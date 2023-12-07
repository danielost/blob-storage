package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob, jsonErr := helpers.GetBlobById(r, request.BlobID)
	if jsonErr != nil {
		ape.RenderErr(w, jsonErr)
		return
	}

	ownerId, err := middleware.GetIdFromJWT(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get id from JWT")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if blob.OwnerId != ownerId {
		helpers.Log(r).Warn("operation forbidden")
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	response := resources.BlobResponse{
		Data: newBlobModel(blob, ownerId),
	}

	ape.Render(w, response)
}
