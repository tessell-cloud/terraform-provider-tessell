package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetDataflixCatalog(availabilityMachineId string) (*model.TessellDmmDataflixServiceView, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dataflix/%s/catalog", c.APIAddress, availabilityMachineId), nil)
	if err != nil {
		return nil, 0, err
	}

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellDmmDataflixServiceView := model.TessellDmmDataflixServiceView{}
	err = json.Unmarshal(body, &tessellDmmDataflixServiceView)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellDmmDataflixServiceView, statusCode, nil
}
