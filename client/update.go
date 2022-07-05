package bexioclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// update a dataset, using the datasetId
// @id is the dataset id
// @data is the updated data
// @apiPath is the path after the HostURL (for it to work you need to add a trailing slash)
func (api *API) Update(id string, data interface{}, method string, apiPath string) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)

	c, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, api.HostURL+apiPath+id, bytes.NewBuffer(c))
	if err != nil {
		return nil, err
	}
	//return nil, fmt.Errorf("ERROR:  s%", req)
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
