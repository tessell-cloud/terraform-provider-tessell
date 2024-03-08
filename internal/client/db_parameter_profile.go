package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetDatabaseParameterProfilesById(id string) (*model.DatabaseParameterProfileResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tessell-internal/parameter-profiles/%s", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	databaseParameterProfileResponse := model.DatabaseParameterProfileResponse{}
	err = json.Unmarshal(body, &databaseParameterProfileResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &databaseParameterProfileResponse, statusCode, nil
}

func (c *Client) GetDatabaseParameterProfilesForConsumers(status *string, engineType *string, name *string) (*model.DatabaseParameterProfileListResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/databases/parameter-profiles", c.APIAddress), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	if !helper.IsNilString(name) {
		q.Add("name", fmt.Sprintf("%v", *name))
	}
	if !helper.IsNilString(engineType) {
		q.Add("engineType", fmt.Sprintf("%v", *engineType))
	}
	if !helper.IsNilString(status) {
		q.Add("status", fmt.Sprintf("%v", *status))
	}
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	databaseParameterProfileListResponse := model.DatabaseParameterProfileListResponse{}
	err = json.Unmarshal(body, &databaseParameterProfileListResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &databaseParameterProfileListResponse, statusCode, nil
}
