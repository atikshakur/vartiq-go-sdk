// Package vartiq provides a Go SDK for the Vartiq API.
// You must provide an API key to use this client.
package vartiq

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	baseURL string
	apiKey  string
	resty   *resty.Client

	Project        *ProjectService
	App            *AppService
	Webhook        *WebhookService
	WebhookMessage *WebhookMessageService
}

// New creates a new Vartiq API client. If baseURL is not provided, it defaults to https://api.us.vartiq.com
func New(apiKey string, baseURL ...string) *Client {
	url := "https://api.us.vartiq.com"
	if len(baseURL) > 0 && baseURL[0] != "" {
		url = baseURL[0]
	}
	r := resty.New().SetBaseURL(url).SetHeader("x-api-key", apiKey)
	c := &Client{
		baseURL: url,
		apiKey:  apiKey,
		resty:   r,
	}
	c.Project = &ProjectService{client: c}
	c.App = &AppService{client: c}
	c.Webhook = &WebhookService{client: c}
	c.WebhookMessage = &WebhookMessageService{client: c}
	return c
}
