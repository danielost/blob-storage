package helpers

import (
	"fmt"
	"net/http"

	"gitlab.com/dl7850949/blob-storage/internal/data"
	"gitlab.com/dl7850949/blob-storage/internal/middleware"
)

func IsAllowed(r *http.Request, blob data.Blob) error {
	userId, err := middleware.GetIdFromJWT(r)
	if err != nil {
		return err
	}

	if blob.OwnerId != userId {
		return fmt.Errorf("user is not the owner of a blob")
	}

	return nil
}
