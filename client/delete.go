package bexioclient

import (
	"fmt"
	"net/http"
)

// this function deletes an dataset with the dataid
// @id is the id of the account to be deleted.
// @apiPath is the path after the HostURL (for it to work you need to add a trailing slash)
func (api *API) Delete(id string, apiPath string) error {
	req, err := http.NewRequest("DELETE", api.HostURL+apiPath+"/"+id, nil)
	if err != nil {
		return err
	}

	resp, _, err := api.doRequest(req)
	if err != nil {
		return err
	}

	if string(resp) != "{\"success\":true}" {
		return fmt.Errorf(string(resp))
	}

	return nil
}
