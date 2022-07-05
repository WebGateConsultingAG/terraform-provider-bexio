package bexioclient

import (
	"encoding/json"
	"net/http"
)

//Gets a list of fictionalUsers
func (api *API) ListFictionalUsers() ([]map[string]interface{}, error) {
	fictionalUsers := make([]map[string]interface{}, 0)
	req, err := http.NewRequest("GET", api.HostURL+"/3.0/fictional_users", nil)
	if err != nil {
		return nil, err
	}

	body, _, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &fictionalUsers)
	if err != nil {
		return nil, err
	}

	return fictionalUsers, nil
}

func (api *API) ReadFictionalUser(fictionalUserId string) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)

	req, err := http.NewRequest("GET", api.HostURL+"/3.0/fictional_users/"+fictionalUserId, nil)
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
