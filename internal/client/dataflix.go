package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetDataflixByAmId(availabilityMachineId string) (*model.TessellAmDataflixDTO, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dataflix/%s", c.APIAddress, availabilityMachineId), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellAmDataflixDTO := model.TessellAmDataflixDTO{}
	err = json.Unmarshal(body, &tessellAmDataflixDTO)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellAmDataflixDTO, statusCode, nil
}

func (c *Client) GetDataflixes(name string, loadAcls bool, owners []string) (*model.TessellDataflixResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dataflix", c.APIAddress), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	q.Add("name", fmt.Sprintf("%v", name))
	q.Add("load-acls", fmt.Sprintf("%v", loadAcls))
	q.Add("owners", strings.Join(owners, ","))
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellDataflixResponse := model.TessellDataflixResponse{}
	err = json.Unmarshal(body, &tessellDataflixResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellDataflixResponse, statusCode, nil
}
