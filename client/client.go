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

// Authenticate performs the client credentials flow to obtain an access token
func (c *Client) Authenticate(ctx context.Context) error {
	// Simple implementation of Client Credentials flow
	// Adjust endpoint and payload as per specific API docs if different, 
	// standard oauth2 usually involves /oauth/token
	
	type authRequest struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
		Scope        string `json:"scope,omitempty"`
	}
	
	type authResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	payload := authRequest{
		ClientID:     c.config.ClientID,
		ClientSecret: c.config.ClientSecret,
		GrantType:    "client_credentials",
	}

	var authResp authResponse
	// Note: We bypass the OnAfterResponse error check for Auth if structure differs, 
	// or we can handle it. For now, we assume standard behavior.
	
	// Waiting for rate limiter
	if err := c.limiter.Wait(ctx); err != nil {
		return err
	}

	resp, err := c.http.R().
		SetContext(ctx).
		SetBody(payload).
		SetResult(&authResp).
		Post("/oauth/token")

	if err != nil {
		return err
	}
	
	// Check specifically for auth failure if not caught by hook
	if resp.IsError() {
		return fmt.Errorf("authentication failed: %s", resp.String())
	}

	c.Token = authResp.AccessToken
	c.ExpiresAt = time.Now().Add(time.Duration(authResp.ExpiresIn) * time.Second)
	
	// Set the token for future requests
	c.http.SetAuthToken(c.Token)
	
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
