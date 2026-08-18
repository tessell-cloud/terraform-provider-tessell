package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"terraform-provider-tessell/internal/model"
)

func (c *Client) GetDatabaseOptionProfilesForConsumption(status string, engineType string, version string) (*model.TessellDatabaseOptionProfileConsumptionListResponse, int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/databases/profiles/options-profiles", c.APIAddress), nil)
	if err != nil {
		return nil, 0, err
	}
	q := req.URL.Query()
	if status != "" {
		q.Add("status", status)
	}
	if engineType != "" {
		q.Add("engine-type", engineType)
	}
	if version != "" {
		q.Add("version", version)
	}
	q.Add("page-size", "1000")
	req.URL.RawQuery = q.Encode()

	body, statusCode, err := c.doRequest(req)
	if err != nil {
		return nil, statusCode, err
	}

	tessellDatabaseOptionProfileConsumptionListResponse := model.TessellDatabaseOptionProfileConsumptionListResponse{}
	err = json.Unmarshal(body, &tessellDatabaseOptionProfileConsumptionListResponse)
	if err != nil {
		return nil, statusCode, err
	}

	return &tessellDatabaseOptionProfileConsumptionListResponse, statusCode, nil
}
