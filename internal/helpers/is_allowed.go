package helpers

import (
	"fmt"

	"gitlab.com/dl7850949/blob-storage/internal/data"
)

func IsAllowed(userId int64, blob *data.Blob) error {
	if blob.OwnerId != userId {
		return fmt.Errorf("user is not the owner of a blob")
	}

	return nil
}
