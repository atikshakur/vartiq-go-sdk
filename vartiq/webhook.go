package vartiq

import (
	"context"
)

type WebhookService struct {
	client *Client
}

type Webhook struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	URL          string         `json:"url"`
	AppID        string         `json:"app"`
	Secret       string         `json:"secret"`
	CustomHeaders []Header      `json:"customHeaders"`
	Headers      []Header       `json:"headers"`
	CreatedAt    string         `json:"createdAt"`
	UpdatedAt    string         `json:"updatedAt"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateWebhookRequest struct {
	Name         string   `json:"name"`
	URL          string   `json:"url"`
	AppID        string   `json:"appId"`
	CustomHeaders []Header `json:"customHeaders,omitempty"`
}

type WebhookResponse struct {
	Data    Webhook `json:"data"`
	Message string  `json:"message"`
	Success bool    `json:"success"`
}

type WebhookListResponse struct {
	Data    []Webhook `json:"data"`
	Message string    `json:"message"`
	Success bool      `json:"success"`
}

func (s *WebhookService) Create(ctx context.Context, req *CreateWebhookRequest) (*WebhookResponse, error) {
	resp := &WebhookResponse{}
	_, err := s.client.resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(resp).
		Post("/webhooks")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *WebhookService) GetAll(ctx context.Context, appID string) (*WebhookListResponse, error) {
	resp := &WebhookListResponse{}
	_, err := s.client.resty.R().
		SetContext(ctx).
		SetQueryParam("appId", appID).
		SetResult(resp).
		Get("/webhooks")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *WebhookService) GetOne(ctx context.Context, webhookID string) (*WebhookResponse, error) {
	resp := &WebhookResponse{}
	_, err := s.client.resty.R().
		SetContext(ctx).
		SetResult(resp).
		Get("/webhooks/" + webhookID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *WebhookService) Update(ctx context.Context, webhookID string, req map[string]interface{}) (*WebhookResponse, error) {
	resp := &WebhookResponse{}
	_, err := s.client.resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(resp).
		Put("/webhooks/" + webhookID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *WebhookService) Delete(ctx context.Context, webhookID string) error {
	_, err := s.client.resty.R().
		SetContext(ctx).
		Delete("/webhooks/" + webhookID)
	return err
} 