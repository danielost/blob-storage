package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	waitForIngest := true
	status, _, err := helpers.BlobsOps(r).Delete(request.BlobID, &waitForIngest)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if status == http.StatusBadRequest {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
