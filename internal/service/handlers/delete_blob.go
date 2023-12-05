package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = BlobsQ(r).Delete(request.BlobID)
	if err != nil {
		Log(r).WithError(err).Error("failed to delete blob from DB")
		ape.Render(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
