package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/shipdeo/shipdeo-go/pkg/core"
)

type HttpClient struct {
	BaseURL      string
	ClientId     string
	ClientSecret string
	HttpClient   *http.Client
}

// Authenticate implements the APIClient interface in core.APIClient.
func (c *HttpClient) Authenticate() (*core.TokenResponse, error) {
	data := strings.NewReader("client_id=" + c.ClientId + "&client_secret=" + c.ClientSecret + "&grant_type=client_credentials")
	req, err := http.NewRequest("POST", c.BaseURL+"/oauth2/connect/token", data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tokenResponse core.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
