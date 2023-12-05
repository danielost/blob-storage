package handlers

import (
	"net/http"

	"gitlab.com/dl7850949/blob-storage/internal/data"
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
	blobsQ := BlobsQ(r)

	blobs, err := blobsQ.Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blobs")
		ape.Render(w, problems.InternalError())
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
		result[i] = newBlobModel(blob)
	}
	return result
}
