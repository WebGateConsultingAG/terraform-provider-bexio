package bexioclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// this function searches the bexio database with the given values in the searchField, searchValue and returns what fits the criteria
// @searchValue the value to search with
// @searchField the name of the field to be searched
// @searchCriteria the criteria to be used to filter the response, to see what criterias there are click on the link(https://docs.bexio.com/#section/API-basics/Search)
// @apiPath is the path after the HostURL (for it to work you need to add a trailing slash)
func (api *API) Search(searchValue string, searchField string, searchCriteria string, apiPath string) ([]map[string]interface{}, error) {
	var jsonData = []byte(`[
		{
		  "field": "` + searchField + `",
		  "value": "` + searchValue + `",
		  "criteria": "` + searchCriteria + `"
		}
	  ]`)
	response := make([]map[string]interface{}, 1)

	req, err := http.NewRequest("POST", api.HostURL+apiPath, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	body, _, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if len(response) > 1 {
		return nil, fmt.Errorf("ERROR: More than one Id was returned, contact_email " + searchValue + " is not unique")
	}

	return response, nil
}
