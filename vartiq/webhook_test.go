package vartiq

import (
	"context"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// Helper to create a WebhookService with a mock client
func newMockWebhookService() (*WebhookService, *resty.Client) {
	r := resty.New()
	c := &Client{resty: r}
	return &WebhookService{client: c}, r
}

func TestWebhookService_Create(t *testing.T) {
	ws, _ := newMockWebhookService()
	ctx := context.Background()
	_, err := ws.Create(ctx, &CreateWebhookRequest{Name: "Test", URL: "http://example.com", AppID: "appId"})
	assert.Error(t, err)
}

func TestWebhookService_GetAll(t *testing.T) {
	ws, _ := newMockWebhookService()
	ctx := context.Background()
	_, err := ws.GetAll(ctx, "appId")
	assert.Error(t, err)
}

func TestWebhookService_GetOne(t *testing.T) {
	ws, _ := newMockWebhookService()
	ctx := context.Background()
	_, err := ws.GetOne(ctx, "webhookId")
	assert.Error(t, err)
}

func TestWebhookService_Update(t *testing.T) {
	ws, _ := newMockWebhookService()
	ctx := context.Background()
	_, err := ws.Update(ctx, "webhookId", map[string]interface{}{"name": "new"})
	assert.Error(t, err)
}

func TestWebhookService_Delete(t *testing.T) {
	ws, _ := newMockWebhookService()
	ctx := context.Background()
	err := ws.Delete(ctx, "webhookId")
	assert.Error(t, err)
}
