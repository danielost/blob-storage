/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"encoding/json"
	"time"
)

type BlobAttributes struct {
	CreatedAt time.Time `json:"created_at"`
	// arbitrary text
	Value json.RawMessage `json:"value"`
}
