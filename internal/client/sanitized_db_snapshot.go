package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetSanitizedDatabaseSnapshot(availabilityMachineId string, id string) (*model.SanitizedDatabaseSnapshot, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s/sanitized-snapshots/%s", c.APIAddress, availabilityMachineId, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	sanitizedDatabaseSnapshot := model.SanitizedDatabaseSnapshot{}
	err = json.Unmarshal(body, &sanitizedDatabaseSnapshot)
	if err != nil {
		return nil, statusCode, err
	}

	return &sanitizedDatabaseSnapshot, statusCode, nil
}

func (c *Client) GetSanitizedDatabaseSnapshots(availabilityMachineId string, name string, manual bool) (*model.GetSanitizedDatabaseSnapshotsResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s/sanitized-snapshots", c.APIAddress, availabilityMachineId), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("name", fmt.Sprintf("%v", name))
	q.Add("manual", fmt.Sprintf("%v", manual))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	getSanitizedDatabaseSnapshotsResponse := model.GetSanitizedDatabaseSnapshotsResponse{}
	err = json.Unmarshal(body, &getSanitizedDatabaseSnapshotsResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &getSanitizedDatabaseSnapshotsResponse, statusCode, nil
}
