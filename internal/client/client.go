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
	TerraformVersion   string
}

func (c *Client) renewTokenInBackground(apiKey *string) {
	time.Sleep(15 * time.Minute)
	for start := time.Now(); time.Since(start) < 10*time.Minute; {
		ar, err := c.SignIn(*apiKey)
		if err == nil {
			c.AuthorizationToken = ar.AccessToken
			go c.renewTokenInBackground(apiKey)
			break
		} else {
			time.Sleep(15 * time.Second)
		}
	}
}

func NewClient(apiAddress *string, apiKey *string, tenantId *string, terraformVersion *string) (*Client, error) {
	c := Client{
		HTTPClient:       &http.Client{Timeout: 30 * time.Second},
		APIAddress:       *apiAddress,
		TenantId:         *tenantId,
		TerraformVersion: *terraformVersion,
	}

	ar, err := c.SignIn(*apiKey)
	if err != nil {
		return nil, err
	}
	c.AuthorizationToken = ar.AccessToken
	go c.renewTokenInBackground(apiKey)

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, int, error) {
	req.Header.Set("tenant-id", c.TenantId)
	req.Header.Set("client-version", c.TerraformVersion)
	req.Header.Set("client-type", "terraform")
	if c.AuthorizationToken != "" {
		req.Header.Set("Authorization", c.AuthorizationToken)
	}
	if req.Method == "POST" || req.Method == "PATCH" {
		req.Header.Set("Content-Type", "application/json")
	}
	if req.Method == "DELETE" && req.ContentLength > 0 {
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
