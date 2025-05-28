# vartiq-go-sdk

A Go SDK for interacting with the Vartiq API. Supports Project, App, Webhook, and Webhook Message resources.

## Installation

```sh
go get github.com/yourusername/vartiq-go-sdk
```

## Usage

### Import and Initialize

```go
import (
	"github.com/yourusername/vartiq-go-sdk/vartiq"
)

client := vartiq.New("YOUR_API_KEY")
// Optionally, override the API URL:
// client := vartiq.New("YOUR_API_KEY", "https://api.eu.vartiq.com")
```

### Go Types

You can import types for strong typing:

```go
import (
	"github.com/yourusername/vartiq-go-sdk/vartiq"
)
// vartiq.Project, vartiq.App, vartiq.Webhook, vartiq.WebhookMessage
```

## API

### Project

```go
// Create a project
projectResp, err := client.Project.Create(ctx, &vartiq.CreateProjectRequest{
	Name:        "Test",
	Description: "desc",
})

// Get all projects
projects, err := client.Project.List(ctx)

// Get a single project
project, err := client.Project.Get(ctx, "PROJECT_ID")

// Update a project
updated, err := client.Project.Update(ctx, "PROJECT_ID", &vartiq.UpdateProjectRequest{
	Name: "New Name",
})

// Delete a project
err := client.Project.Delete(ctx, "PROJECT_ID")
```

### App

```go
// Create an app
appResp, err := client.App.Create(ctx, &vartiq.CreateAppRequest{
	Name:        "App Name",
	Environment: "development",
})

// Get all apps
apps, err := client.App.List(ctx)

// Get a single app
app, err := client.App.Get(ctx, "APP_ID")

// Update an app
updated, err := client.App.Update(ctx, "APP_ID", &vartiq.UpdateAppRequest{
	Name: "New App Name",
})

// Delete an app
err := client.App.Delete(ctx, "APP_ID")
```

### Webhook

```go
// Create a webhook
webhookResp, err := client.Webhook.Create(ctx, &vartiq.CreateWebhookRequest{
	Name:   "Webhook",
	URL:    "https://your-webhook-url.com",
	AppID:  "APP_ID",
	CustomHeaders: []vartiq.Header{{Key: "x-app", Value: "x-value"}}, // optional
})

// Get all webhooks for an app
webhooks, err := client.Webhook.GetAll(ctx, "APP_ID")

// Get a single webhook
webhook, err := client.Webhook.GetOne(ctx, "WEBHOOK_ID")

// Update a webhook
updated, err := client.Webhook.Update(ctx, "WEBHOOK_ID", map[string]interface{}{
	"name": "New Webhook Name",
})

// Delete a webhook
err := client.Webhook.Delete(ctx, "WEBHOOK_ID")
```

### Webhook Message

```go
// Create a webhook message
msgResp, err := client.WebhookMessage.Create(ctx, &vartiq.CreateWebhookMessageRequest{
	WebhookID: "WEBHOOK_ID",
	Payload:   map[string]interface{}{ "hello": "world" },
})
```

## Supported Resources
- Project
- App
- Webhook
- Webhook Message

## API Key
You must provide your API key when creating the client:

```go
client := vartiq.New("https://api.vartiq.com", "your_api_key")
``` 