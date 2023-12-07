package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
	"gitlab.com/dl7850949/blob-storage/resources"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob, ok := getBlobById(w, r, request.BlobID)
	if !ok {
		return
	}

	ownerId, allowed := isAllowed(w, r, blob)
	if !allowed {
		return
	}

	response := resources.BlobResponse{
		Data: newBlobModel(blob, ownerId),
	}

	ape.Render(w, response)
}
