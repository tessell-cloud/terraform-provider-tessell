package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func (c *Client) CreateDatabaseSnapshotRequest(availabilityMachineId string, payload model.CreateDatabaseSnapshotTaskPayload) (*model.TaskSummary, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/availability-machines/%s/snapshots", c.APIAddress, availabilityMachineId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	taskSummary := model.TaskSummary{}
	err = json.Unmarshal(body, &taskSummary)
	if err != nil {
		return nil, statusCode, err
	}

	return &taskSummary, statusCode, nil
}

func (c *Client) DeleteDatabaseSnapshotRequest(availabilityMachineId string, id string) (*model.APIStatus, int, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/availability-machines/%s/snapshots/%s", c.APIAddress, availabilityMachineId, id), nil)
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

func (c *Client) GetDatabaseSnapshot(availabilityMachineId string, id string) (*model.DatabaseSnapshot, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s/snapshots/%s", c.APIAddress, availabilityMachineId, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	databaseSnapshot := model.DatabaseSnapshot{}
	err = json.Unmarshal(body, &databaseSnapshot)
	if err != nil {
		return nil, statusCode, err
	}

	return &databaseSnapshot, statusCode, nil
}

func (c *Client) GetDatabaseSnapshots(availabilityMachineId string, name *string, manual *bool) (*model.GetDatabaseSnapshotsResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s/snapshots", c.APIAddress, availabilityMachineId), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	if !helper.IsNilString(name) {
		q.Add("name", fmt.Sprintf("%v", *name))
	}
	if !helper.IsNilBool(manual) {
		q.Add("manual", fmt.Sprintf("%v", *manual))
	}
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	getDatabaseSnapshotsResponse := model.GetDatabaseSnapshotsResponse{}
	err = json.Unmarshal(body, &getDatabaseSnapshotsResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &getDatabaseSnapshotsResponse, statusCode, nil
}

func (c *Client) DBSnapshotPollForStatus(availabilityMachineId string, id string, value string, timeout int, interval int) error {

	loopCount := 0
	sleepCycleDurationSmall, err := time.ParseDuration("10s")
	if err != nil {
		return err
	}
	sleepCycleDuration, err := time.ParseDuration(fmt.Sprintf("%ds", interval))
	if err != nil {
		return err
	}

	loops := timeout/int(sleepCycleDuration.Seconds()) + 5

	for {
		response, _, err := c.GetDatabaseSnapshot(availabilityMachineId, id)
		if err != nil {
			return fmt.Errorf("error while polling: %s", err.Error())
		}

		switch *response.Status {
		case value:
			return nil
		case "FAILED":
			return fmt.Errorf("received status FAILED while polling")
		}

		loopCount = loopCount + 1
		if loopCount > loops {
			return fmt.Errorf("timed out while polling")
		}
		if loopCount > 6 {
			time.Sleep(sleepCycleDuration)
		} else {
			time.Sleep(sleepCycleDurationSmall)
		}
	}
}
