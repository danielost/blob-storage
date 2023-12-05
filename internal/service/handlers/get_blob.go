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

	blob, err := BlobsQ(r).FilterByID(request.BlobID).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if blob == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	result := resources.BlobResponse{
		Data: newBlobModel(*blob),
	}

	ape.Render(w, result)
}
