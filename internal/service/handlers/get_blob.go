package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	byteResponse, err := helpers.BlobsOps(r).GetById(request.BlobID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get blob")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if len(byteResponse) == 0 {
		helpers.Log(r).Info("blob not found")
		ape.RenderErr(w, problems.NotFound())
		return
	}

	response, err := helpers.Unmarshal(byteResponse)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to unmarshal TokenD response")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, response)
}
