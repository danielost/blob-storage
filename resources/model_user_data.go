/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UserData struct {
	Key
	Attributes UserDataAttributes `json:"attributes"`
}
type UserDataResponse struct {
	Data     UserData `json:"data"`
	Included Included `json:"included"`
}

type UserDataListResponse struct {
	Data     []UserData `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustUserData - returns UserData from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustUserData(key Key) *UserData {
	var userData UserData
	if c.tryFindEntry(key, &userData) {
		return &userData
	}
	return nil
}
