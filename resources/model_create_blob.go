/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateBlob struct {
	Key
	Attributes CreateBlobAttributes `json:"attributes"`
}
type CreateBlobResponse struct {
	Data     CreateBlob `json:"data"`
	Included Included   `json:"included"`
}

type CreateBlobListResponse struct {
	Data     []CreateBlob `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustCreateBlob - returns CreateBlob from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreateBlob(key Key) *CreateBlob {
	var createBlob CreateBlob
	if c.tryFindEntry(key, &createBlob) {
		return &createBlob
	}
	return nil
}
