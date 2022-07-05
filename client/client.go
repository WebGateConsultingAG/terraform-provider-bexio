package bexioclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//Client -
type API struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

//NewClient -
func NewClient(url, token string) (*API, error) {
	api := API{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    url,
		Token:      token,
	}
	return &api, nil
}

//creates a new request with all the required Headers
func (api *API) doRequest(req *http.Request) ([]byte, int, error) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+api.Token)

	res, err := api.HTTPClient.Do(req)
	if err != nil {
		return nil, -1, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		return body, res.StatusCode, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	if err != nil {
		return nil, res.StatusCode, err
	}

	return body, res.StatusCode, err
}
