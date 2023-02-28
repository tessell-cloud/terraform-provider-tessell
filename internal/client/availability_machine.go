package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetAvailabilityMachine(id string) (*model.TessellDMMServiceConsumerDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/availability-machines/%s", c.APIAddress, id), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellDMMServiceConsumerDTO := model.TessellDMMServiceConsumerDTO{}
	err = json.Unmarshal(body, &tessellDMMServiceConsumerDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellDMMServiceConsumerDTO, statusCode, nil
}

func (c *Client) GetAvailabilityMachines(name string, status string, engineType string, loadAcls bool, owners []string) (*model.GetDMMsServiceView, int, error) {
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

	getDMMsServiceView := model.GetDMMsServiceView{}
	err = json.Unmarshal(body, &getDMMsServiceView)
	if err != nil {
		return nil, statusCode, err
	}

	return &getDMMsServiceView, statusCode, nil
}
