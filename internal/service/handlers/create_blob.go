package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	horizon2 "gitlab.com/dl7850949/blob-storage/internal/data/horizon"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
)

func CreateBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob := horizon2.Blob{
		Value: request.Data.Attributes.Value,
	}

	waitForIngest := true
	_, byteResponse, err := helpers.BlobsOps(r).Create(blob, &waitForIngest)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to insert blob")
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
