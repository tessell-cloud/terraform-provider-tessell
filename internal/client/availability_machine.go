package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetAvailabilityMachine(id string) (*model.TessellDmmServiceConsumerDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellDmmServiceConsumerDTO := model.TessellDmmServiceConsumerDTO{}
	err = json.Unmarshal(body, &tessellDmmServiceConsumerDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellDmmServiceConsumerDTO, statusCode, nil
}

func (c *Client) GetAvailabilityMachines(name string, status string, engineType string, loadAcls bool, owners []string) (*model.GetDmmsServiceView, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines", c.APIAddress), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("name", fmt.Sprintf("%v", name))
	q.Add("status", fmt.Sprintf("%v", status))
	q.Add("engine-type", fmt.Sprintf("%v", engineType))
	q.Add("load-acls", fmt.Sprintf("%v", loadAcls))
	q.Add("owners", strings.Join(owners, ","))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	getDmmsServiceView := model.GetDmmsServiceView{}
	err = json.Unmarshal(body, &getDmmsServiceView)
	if err != nil {
		return nil, statusCode, err
	}

	return &getDmmsServiceView, statusCode, nil
}
