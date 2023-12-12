package helpers

import "encoding/json"

func Unmarshal(b []byte) (map[string]interface{}, error) {
	var response map[string]interface{}
	err := json.Unmarshal(b, &response)
	return response, err
}
