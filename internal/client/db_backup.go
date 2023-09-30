package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetBackupRequest(availabilityMachineId string, id string) (*model.DatabaseBackup, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s/backups/%s", c.APIAddress, availabilityMachineId, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	databaseBackup := model.DatabaseBackup{}
	err = json.Unmarshal(body, &databaseBackup)
	if err != nil {
		return nil, statusCode, err
	}

	return &databaseBackup, statusCode, nil
}

func (c *Client) GetDatabaseBackups(availabilityMachineId string, name string, manual bool) (*model.GetDatabaseBackupsResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s/backups", c.APIAddress, availabilityMachineId), nil)
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

	getDatabaseBackupsResponse := model.GetDatabaseBackupsResponse{}
	err = json.Unmarshal(body, &getDatabaseBackupsResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &getDatabaseBackupsResponse, statusCode, nil
}
