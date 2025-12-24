# VersaFleet Go SDK

Unofficial Golang SDK for the [VersaFleet API](https://versafleet.docs.apiary.io/).

## Installation

```bash
go get github.com/Willias7788/go-versafleet-sdk
```

## Configuration

The SDK uses `viper` for configuration. You can configure it via environment variables or a `.env` file.

### Environment Variables

*   `VERSAFLEET_BASE_URL`: API Base URL (default: `https://api.versafleet.co/api`)
*   `VERSAFLEET_CLIENT_ID`: OAuth2 Client ID
*   `VERSAFLEET_CLIENT_SECRET`: OAuth2 Client Secret
*   `VERSAFLEET_DEBUG`: Enable debug logging (true/false)

### .env Example

```ini
VERSAFLEET_CLIENT_ID=your_client_id
VERSAFLEET_CLIENT_SECRET=your_client_secret
```

## Quickstart

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Willias7788/go-versafleet-sdk/client"
	"github.com/Willias7788/go-versafleet-sdk/config"
	"github.com/Willias7788/go-versafleet-sdk/jobs"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize client
	c := client.New(cfg)

	// Authenticate
	ctx := context.Background()
	if err := c.Authenticate(ctx); err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	// access Jobs service
	jobsService := jobs.New(c)

	// List jobs
	iterator := jobsService.List(ctx, client.ListOptions{PerPage: 10})
	for iterator.Next() {
		job := iterator.Value()
		fmt.Printf("Job: %s - %s\n", job.JobNumber, job.Status)
	}

	if err := iterator.Err(); err != nil {
		log.Printf("Error iterating jobs: %v", err)
	}
}
```

## Features

### Rate Limiting

The SDK automatically adheres to the 100 requests/minute limit using a token bucket algorithm.

### Pagination

List endpoints return an `Iterator` helper to easily traverse pages.

```go
iter := tasksService.List(ctx, client.ListOptions{Page: 1, PerPage: 50})
for iter.Next() {
    task := iter.Value()
    // process task
}
```

### Webhooks

Helper to validate and parse webhooks.

```go
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    event, err := webhooks.Parse(r, "your_webhook_secret")
    if err != nil {
        http.Error(w, "Invalid webhook", http.StatusBadRequest)
        return
    }
    
    fmt.Printf("Received event: %s\n", event.Type)
}
```

### Error Handling

API errors are returned as `*client.APIError` structs containing the status code, message, and request ID.

```go
if err != nil {
    if apiErr, ok := err.(*client.APIError); ok {
        if apiErr.IsNotFound() {
            // handle 404
        }
    }
}
```

## Modules Covered

*   Jobs
*   Tasks
*   Drivers (placeholder)
*   Vehicles (placeholder)
*   Webhooks
*   (Add others as implemented)
