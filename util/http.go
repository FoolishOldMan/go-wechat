package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response2Struct response转struct
func Response2Struct(resp *http.Response, obj interface{}) error {
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(result, obj)
}
