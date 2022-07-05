package bexioclient

import (
	"encoding/json"
	"net/http"
)

//gets a list of contacts
func (api *API) ListContacts() ([]map[string]interface{}, error) {
	contacts := make([]map[string]interface{}, 0)
	req, err := http.NewRequest("GET", api.HostURL+"/2.0/contact", nil)
	if err != nil {
		return nil, err
	}

	body, _, err := api.doRequest(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

//gets a single contact with the corresponding Id
func (api *API) ReadContact(contactId string) (map[string]interface{}, error) {
	result := make(map[string]interface{}, 0)

	req, err := http.NewRequest("GET", api.HostURL+"/2.0/contact/"+contactId, nil)
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
