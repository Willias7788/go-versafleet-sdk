package account

import (
	"context"
	"fmt"

	"github.com/Willias7788/go-versafleet-sdk/client"
)

type Service struct {
	client *client.Client
}

func New(c *client.Client) *Service {
	return &Service{client: c}
}

type Account struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone,omitempty"`
	Address string `json:"address,omitempty"`
	// Add other fields based on the actual API response
}

// Get retrieving a specific account information by ID
// Reference: https://versafleet.docs.apiary.io/#reference/0/account-api/view-a-account
func (s *Service) Get(ctx context.Context, id string) (*Account, error) {
	var account Account
	path := fmt.Sprintf("/accounts/id")
	err := s.client.Get(ctx, path, &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Create creates a new account
func (s *Service) Create(ctx context.Context, account *Account) (*Account, error) {
	var createdAccount Account
	err := s.client.Post(ctx, "/accounts", account, &createdAccount)
	if err != nil {
		return nil, err
	}
	return &createdAccount, nil
}

// Update updates an existing account
func (s *Service) Update(ctx context.Context, id string, account *Account) (*Account, error) {
	var updatedAccount Account
	path := fmt.Sprintf("/accounts/%s", id)
	err := s.client.Put(ctx, path, account, &updatedAccount)
	if err != nil {
		return nil, err
	}
	return &updatedAccount, nil
}

// Delete deletes an account
func (s *Service) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/accounts/%s", id)
	return s.client.Delete(ctx, path)
}
