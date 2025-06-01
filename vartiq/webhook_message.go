package vartiq

import (
	"context"
)

type WebhookMessageService struct {
	client *Client
}

type WebhookMessage struct {
	ID        string      `json:"id"`
	AppID     string      `json:"app"`
	Payload   interface{} `json:"payload"`
	Signature string      `json:"signature"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt string      `json:"updatedAt"`
}

type CreateWebhookMessageRequest struct {
	AppID   string      `json:"appId"`
	Payload interface{} `json:"payload"`
}

type CreateWebhookMessageResponse struct {
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
	Success bool                   `json:"success"`
}

func (s *WebhookMessageService) Create(ctx context.Context, req *CreateWebhookMessageRequest) (*CreateWebhookMessageResponse, error) {
	resp := &CreateWebhookMessageResponse{}
	_, err := s.client.resty.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(resp).
		Post("/webhook-messages")
	if err != nil {
		return nil, err
	}
	return resp, nil
}
