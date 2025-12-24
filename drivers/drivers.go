package drivers

import (
	"github.com/Willias7788/go-versafleet-sdk/client"
)

type Service struct {
	client *client.Client
}

func New(c *client.Client) *Service {
	return &Service{client: c}
}

// Add methods here
