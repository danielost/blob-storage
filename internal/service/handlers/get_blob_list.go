package handlers

import (
	"net/http"

	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBlobsList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobsListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	ownerId, err := middleware.GetIdFromJWT(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get id from JWT")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	blobsQ := helpers.BlobsQ(r)
	blobsQ.Page(request.OffsetPageParams)

	blobs, err := blobsQ.FilterByOwnerId(ownerId).Select()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get blobs")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	response := resources.BlobListResponse{
		Data:  newBlobsList(blobs),
		Links: GetOffsetLinks(r, request.OffsetPageParams),
	}

	ape.Render(w, response)
}

func newBlobsList(blobs []data.Blob) []resources.Blob {
	result := make([]resources.Blob, len(blobs))
	for i, blob := range blobs {
		result[i] = newBlobModel(&blob, blob.OwnerId)
	}
	return result
}
