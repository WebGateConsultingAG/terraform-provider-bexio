package bexioclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//this function takes data and creates a new Element
// @data is the data to create the element with
// @apiPath is the path after the HostURL (for it to work you need to add a trailing slash)
func (api *API) Create(data interface{}, apiPath string) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)

	c, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", api.HostURL+apiPath, bytes.NewBuffer(c))
	if err != nil {
		return nil, err
	}
	resp, _, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
