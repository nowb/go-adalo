package adalo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

var DefaultBaseURL = "https://api.adalo.com"

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("apiKey not provided")
	}

	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    DefaultBaseURL,
		APIKey:     apiKey,
	}, nil
}

func (c *Client) SendNotification(input *SendNotificationInput) (*SendNotificationOutput, error) {
	if input == nil {
		return nil, errors.New("input not provided")
	}

	body, _ := json.Marshal(mapSendNotificationInput(input))
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/notifications", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("preparing new HTTP request failed: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending HTTP request failed: %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading HTTP response body failed: %w", err)
	}

	var output SendNotificationOutput
	if err := json.Unmarshal(respBody, &output); err != nil {
		return nil, fmt.Errorf("unmarshaling HTTP response body failed: %w", err)
	}

	return &output, nil
}
