package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	custom_errors "github.com/maths-lover/chatgpt/utils"
)

const (
	chatgpt_api_url = "https://api.openai.com/v1"
)

type Client struct {
	client *http.Client
	Config *Config
}

type Config struct {
	// Base URL for API requests.
	BaseURL string

	// API Key (Required)
	APIKey string

	// Organization ID (Optional)
	OrganizationID string
}

func NewClient(apikey string) (*Client, error) {
	if apikey == "" {
		return nil, custom_errors.ErrNoAPI
	}

	return &Client{
		client: &http.Client{},
		Config: &Config{
			BaseURL: chatgpt_api_url,
			APIKey:  apikey,
		},
	}, nil
}

func NewClientWithConfig(config *Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, custom_errors.ErrNoAPI
	}

	return &Client{
		client: &http.Client{},
		Config: config,
	}, nil
}

func (c *Client) SendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Config.APIKey))
	if c.Config.OrganizationID != "" {
		req.Header.Set("OpenAI-Organization", c.Config.OrganizationID)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		// Parse body
		var errMessage interface{}
		if err := json.NewDecoder(res.Body).Decode(&errMessage); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("api request failed: status Code: %d %s %s Message: %+v", res.StatusCode, res.Status, res.Request.URL, errMessage)
	}

	return res, nil
}
