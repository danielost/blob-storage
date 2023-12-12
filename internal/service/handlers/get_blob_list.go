package handlers

import (
	"net/http"

	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetBlobsList(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobsListRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	byteResponse, err := helpers.BlobsOps(r).Get(request.OffsetPageParams)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get blobs")
		ape.RenderErr(w, problems.InternalError())
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
