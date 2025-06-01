package vartiq

import (
	"context"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// Helper to create a WebhookMessageService with a mock client
func newMockWebhookMessageService() (*WebhookMessageService, *resty.Client) {
	r := resty.New()
	c := &Client{resty: r}
	return &WebhookMessageService{client: c}, r
}

func TestWebhookMessageService_Create(t *testing.T) {
	wms, _ := newMockWebhookMessageService()
	ctx := context.Background()
	_, err := wms.Create(ctx, &CreateWebhookMessageRequest{AppID: "appId", Payload: map[string]interface{}{}})
	assert.Error(t, err)
}
