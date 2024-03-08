package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) CreateServiceStartStopSchedule(serviceId string, payload model.CreateStartStopSchedulePayload) (*model.StartStopScheduleDTO, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/%s/schedules/start-stop", c.APIAddress, serviceId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	startStopScheduleDTO := model.StartStopScheduleDTO{}
	err = json.Unmarshal(body, &startStopScheduleDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &startStopScheduleDTO, statusCode, nil
}

func (c *Client) DeleteServiceStartStopSchedule(serviceId string, id string) (*model.APIStatus, int, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/services/%s/schedules/start-stop/%s", c.APIAddress, serviceId, id), nil)
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

func (c *Client) GetServiceStartStopSchedule(serviceId string, id string) (*model.StartStopScheduleDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/%s/schedules/start-stop/%s", c.APIAddress, serviceId, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	startStopScheduleDTO := model.StartStopScheduleDTO{}
	err = json.Unmarshal(body, &startStopScheduleDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &startStopScheduleDTO, statusCode, nil
}

func (c *Client) UpdateServiceStartStopSchedule(serviceId string, id string, payload model.UpdateStartStopSchedulePayload) (*model.StartStopScheduleDTO, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/%s/schedules/start-stop/%s", c.APIAddress, serviceId, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	startStopScheduleDTO := model.StartStopScheduleDTO{}
	err = json.Unmarshal(body, &startStopScheduleDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &startStopScheduleDTO, statusCode, nil
}
