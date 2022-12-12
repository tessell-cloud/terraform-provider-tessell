package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type APIKeyPayload struct {
	APIKey string `json:"apiKey"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
}

func (c *Client) SignIn(apiKey string) (*AuthResponse, error) {
	rb, err := json.Marshal(APIKeyPayload{APIKey: apiKey})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/iam/authorize", c.APIAddress), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *Client) SignOut() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signout", c.APIAddress), strings.NewReader(string("")))
	if err != nil {
		return err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Signed out user" {
		return errors.New(string(body))
	}

	return nil
}
