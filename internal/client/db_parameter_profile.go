package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func (c *Client) GetDatabaseParameterProfileById(id string) (*model.DatabaseParameterProfileResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/governance/parameter-profiles/%s", c.APIAddress, id), nil)
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

func (c *Client) CreateDatabaseParameterProfile(payload model.DatabaseParameterProfileRequest) (*model.DatabaseParameterProfileResponse, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/governance/parameter-profiles", c.APIAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

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

func (c *Client) UpdateDatabaseParameterProfile(id string, payload model.DatabaseParameterProfilePatchRequest) (*model.DatabaseParameterProfileResponse, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/governance/parameter-profiles/%s", c.APIAddress, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

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

func (c *Client) DeleteDatabaseParameterProfile(id string) (*model.APIStatus, int, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/governance/parameter-profiles/%s", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	apiStatus := model.APIStatus{}
	err = json.Unmarshal(body, &apiStatus)
	if err != nil {
		return nil, statusCode, err
	}

	return &apiStatus, statusCode, nil
}

func (c *Client) ResetDatabaseParameterProfile(id string) (*model.DatabaseParameterProfileResponse, int, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/governance/parameter-profiles/%s/reset", c.APIAddress, id), nil)
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

// UpdateDatabaseParameterProfileMaturityStatus changes the maturity status of a parameter profile.
// action must be one of: "publish", "unpublish", "draft"
func (c *Client) UpdateDatabaseParameterProfileMaturityStatus(name string, action string) (*model.DatabaseParameterProfileResponse, int, error) {
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/databases/profiles/governance/parameter-profiles/%s/%s", c.APIAddress, name, action), nil)
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

func (c *Client) GetParameterProfileUsages(id string) (*model.ParameterProfileUsageResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/governance/parameter-profiles/%s/usages", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	usageResponse := model.ParameterProfileUsageResponse{}
	err = json.Unmarshal(body, &usageResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &usageResponse, statusCode, nil
}
