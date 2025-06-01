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

type webhookMessageResponse struct {
	Data    WebhookMessage `json:"data"`
	Message string         `json:"message"`
	Success bool           `json:"success"`
}

// Create sends a message to a webhook. The payload can be any JSON-serializable value.
// Example:
//
//	message, err := client.WebhookMessage.Create(ctx, "APP_ID", map[string]interface{}{
//	    "hello": "world",
//	})
func (s *WebhookMessageService) Create(ctx context.Context, appID string, payload interface{}) (*WebhookMessage, error) {
	resp := &webhookMessageResponse{}
	_, err := s.client.resty.R().
		SetContext(ctx).
		SetBody(map[string]interface{}{
			"appId":   appID,
			"payload": payload,
		}).
		SetResult(resp).
		Post("/webhook-messages")
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
