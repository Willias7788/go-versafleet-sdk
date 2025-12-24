package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// EventType represents the type of webhook event
type EventType string

const (
	EventTypeJobCreated   EventType = "job.created"
	EventTypeJobUpdated   EventType = "job.updated"
	EventTypeTaskCompleted EventType = "task.completed"
	// Add others
)

type Event struct {
	ID        string          `json:"id"`
	Type      EventType       `json:"type"`
	CreatedAt string          `json:"created_at"`
	Data      json.RawMessage `json:"data"`
}

// Parse reads the request body, validates the signature, and returns the event
func Parse(req *http.Request, secret string) (*Event, error) {
	signature := req.Header.Get("X-Versafleet-Signature")
	if signature == "" {
		return nil, errors.New("missing signature header")
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if !VerifySignature(body, signature, secret) {
		return nil, errors.New("invalid signature")
	}

	var event Event
	if err := json.Unmarshal(body, &event); err != nil {
		return nil, err
	}

	return &event, nil
}

// VerifySignature checks the HMAC-SHA256 signature
func VerifySignature(payload []byte, signature, secret string) bool {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(payload)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}
