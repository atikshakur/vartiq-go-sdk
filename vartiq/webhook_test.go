package vartiq

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

func TestWebhookService_Verify(t *testing.T) {
	ws, _ := newMockWebhookService()

	secret := "testsecret"
	payload := []byte("testpayload")

	// Generate a valid signature for testing
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	tests := []struct {
		name                 string
		payload              []byte
		signature            string
		secret               string
		expectedError        bool
		expectedErrorMessage string
	}{
		{
			name:          "Valid signature",
			payload:       payload,
			signature:     expectedSignature,
			secret:        secret,
			expectedError: false,
		},
		{
			name:                 "Missing signature",
			payload:              payload,
			signature:            "",
			secret:               secret,
			expectedError:        true,
			expectedErrorMessage: "signature header is missing",
		},
		{
			name:                 "Invalid signature",
			payload:              payload,
			signature:            "abcdef1234567890",
			secret:               secret,
			expectedError:        true,
			expectedErrorMessage: "signature verification failed",
		},
		{
			name:                 "Invalid signature format",
			payload:              payload,
			signature:            "not-a-hex-string",
			secret:               secret,
			expectedError:        true,
			expectedErrorMessage: "failed to decode signature: encoding/hex: invalid byte: U+006E 'n'",
		},
		{
			name:                 "Mismatched secret",
			payload:              payload,
			signature:            expectedSignature,
			secret:               "wrongsecret",
			expectedError:        true,
			expectedErrorMessage: "signature verification failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifiedPayload, err := ws.Verify(tt.payload, tt.signature, tt.secret)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrorMessage)
				assert.Nil(t, verifiedPayload)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.payload, verifiedPayload)
			}
		})
	}
}
