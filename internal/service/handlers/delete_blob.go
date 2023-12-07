package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/helpers"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
	"gitlab.com/dl7850949/blob-storage/internal/service/requests"
)

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetBlobRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob, ok := getBlobById(w, r, request.BlobID)
	if !ok {
		return
	}

	_, allowed := isAllowed(w, r, blob)
	if !allowed {
		return
	}

	err = helpers.BlobsQ(r).Delete(request.BlobID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to delete blob from DB")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getBlobById(w http.ResponseWriter, r *http.Request, blobId int64) (*data.Blob, bool) {
	blob, err := helpers.BlobsQ(r).FilterByID(blobId).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get blob from DB")
		ape.RenderErr(w, problems.InternalError())
		return blob, false
	}

	if blob == nil {
		helpers.Log(r).Info("blob not found")
		ape.RenderErr(w, problems.NotFound())
		return blob, false
	}

	return blob, true
}

func isAllowed(w http.ResponseWriter, r *http.Request, blob *data.Blob) (int64, bool) {
	ownerId, err := middleware.GetIdFromJWT(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get id from JWT")
		ape.RenderErr(w, problems.InternalError())
		return ownerId, false
	}

	if blob.ID != ownerId {
		helpers.Log(r).Info("operation forbidden")
		ape.RenderErr(w, problems.Forbidden())
		return ownerId, false
	}

	return ownerId, true
}
