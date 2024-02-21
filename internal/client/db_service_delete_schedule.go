package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) CreateServiceDeletionSchedule(serviceId string, payload model.DeletionSchedulePayload) (*model.DeletionScheduleDTO, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/services/%s/schedules/delete", c.APIAddress, serviceId), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	deletionScheduleDTO := model.DeletionScheduleDTO{}
	err = json.Unmarshal(body, &deletionScheduleDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &deletionScheduleDTO, statusCode, nil
}

func (c *Client) DeleteServiceDeletionScheduleTFP(serviceId string, id string) (*model.APIStatus, int, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/services/%s/schedules/delete/%s", c.APIAddress, serviceId, id), nil)
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

func (c *Client) GetServiceDeletionScheduleTFP(serviceId string, id string) (*model.DeletionScheduleDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/services/%s/schedules/delete/%s", c.APIAddress, serviceId, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	deletionScheduleDTO := model.DeletionScheduleDTO{}
	err = json.Unmarshal(body, &deletionScheduleDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &deletionScheduleDTO, statusCode, nil
}

func (c *Client) UpdateServiceDeletionScheduleTFP(serviceId string, id string, payload model.DeletionSchedulePayload) (*model.DeletionScheduleDTO, int, error) {
	rb, err := json.Marshal(payload)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/services/%s/schedules/delete/%s", c.APIAddress, serviceId, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, 0, err
	}

	defer req.Body.Close()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	deletionScheduleDTO := model.DeletionScheduleDTO{}
	err = json.Unmarshal(body, &deletionScheduleDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &deletionScheduleDTO, statusCode, nil
}
