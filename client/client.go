package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Willias7788/go-versafleet-sdk/config"
	"github.com/Willias7788/go-versafleet-sdk/rate"
	"github.com/go-resty/resty/v2"
	ratelimit "golang.org/x/time/rate"
)

type Client struct {
	http      *resty.Client
	config    *config.Config
	limiter   *ratelimit.Limiter
	Token     string
	ExpiresAt time.Time
}

// New creates a new VersaFleet API client
func New(cfg *config.Config) *Client {
	r := resty.New()
	r.SetBaseURL(cfg.BaseURL)
	r.SetTimeout(time.Minute) // Default timeout

	if cfg.Debug {
		r.SetDebug(true)
	}

	// User indicated that auth might be via query params
	r.SetQueryParam("client_id", cfg.ClientID)
	r.SetQueryParam("client_secret", cfg.ClientSecret)

	c := &Client{
		http:    r,
		config:  cfg,
		limiter: rate.Default(),
	}

	r.SetRetryCount(3)
	r.SetRetryWaitTime(500 * time.Millisecond)
	r.SetRetryMaxWaitTime(2000 * time.Millisecond)

	// Error hook to parse API errors
	r.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		if resp.IsError() {
			apiErr := &APIError{
				StatusCode: resp.StatusCode(),
				RequestID:  resp.Header().Get("X-Request-Id"),
				// Message:    resp.Request.URL,
			}
			// Attempt to unmarshal body into error
			_ = json.Unmarshal(resp.Body(), apiErr)

			// If message is empty, use status text
			if apiErr.Message == "" {
				apiErr.Message = http.StatusText(resp.StatusCode())
			}
			return apiErr
		}
		return nil
	})

	return c
}

// Verify checks the credentials by making a lightweight API call (e.g. List Jobs with limit 1)
func (c *Client) Verify(ctx context.Context) error {
	// Attempt to list jobs with a small limit to verify auth
	// We use raw http request here to avoid circular dependency if we used jobs package
	// But simply checking if we get a 200 OK from an endpoint is enough.

	// Assuming /jobs is a valid endpoint that requires auth.
	resp, err := c.R(ctx).
		SetQueryParam("per_page", "1").
		Get("/v2/jobs")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("verification failed: %s", resp.String())
	}

	return nil
}

// R creates a new request with the context and limiter wait
func (c *Client) R(ctx context.Context) *resty.Request {
	_ = c.limiter.Wait(ctx)
	return c.http.R().SetContext(ctx)
}

// REST methods helpers

func (c *Client) Get(ctx context.Context, path string, result interface{}) error {
	_, err := c.R(ctx).SetResult(result).Get(path)
	return err
}

func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	if err := c.validatePayloadSize(body); err != nil {
		return err
	}
	_, err := c.R(ctx).SetBody(body).SetResult(result).Post(path)
	return err
}

func (c *Client) Put(ctx context.Context, path string, body interface{}, result interface{}) error {
	if err := c.validatePayloadSize(body); err != nil {
		return err
	}
	_, err := c.R(ctx).SetBody(body).SetResult(result).Put(path)
	return err
}

func (c *Client) validatePayloadSize(body interface{}) error {
	if body == nil {
		return nil
	}
	// Marshal to check size
	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal body to check size: %w", err)
	}
	if len(data) > 3*1024*1024 { // 3MB
		return fmt.Errorf("payload size %d bytes exceeds 3MB limit", len(data))
	}
	return nil
}

func (c *Client) Delete(ctx context.Context, path string) error {
	_, err := c.R(ctx).Delete(path)
	return err
}
