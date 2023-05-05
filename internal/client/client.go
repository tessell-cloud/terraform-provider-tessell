package client

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	APIAddress         string
	HTTPClient         *http.Client
	AuthorizationToken string
	TenantId           string
}

func NewClient(apiAddress *string, apiKey *string, tenantId *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		APIAddress: *apiAddress,
		TenantId:   *tenantId,
	}

	ar, err := c.SignIn(*apiKey)
	if err != nil {
		return nil, err
	}
	c.AuthorizationToken = ar.AccessToken

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, int, error) {
	req.Header.Set("tenant-id", c.TenantId)
	if c.AuthorizationToken != "" {
		req.Header.Set("Authorization", c.AuthorizationToken)
	}
	if req.Method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		if res != nil {
			return nil, res.StatusCode, err
		}
		return nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode, fmt.Errorf("%s", body)
	}

	return body, res.StatusCode, err
}
