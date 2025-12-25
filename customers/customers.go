package customers

import (
	"context"
	"fmt"

	"github.com/Willias7788/go-versafleet-sdk/client"
	"github.com/Willias7788/go-versafleet-sdk/model"
)

type Service struct {
	client *client.Client
}

func New(c *client.Client) *Service {
	return &Service{client: c}
}

type CustomerListResponse struct {
	Customers []model.Customer `json:"customers"`
	Meta      *model.Meta      `json:"meta"`
}

// List returns an iterator to list all customers
func (s *Service) List(ctx context.Context, opts *model.CustomerListOptions) *client.Iterator[model.Customer, *model.CustomerListOptions] {
	return client.NewIterator(ctx, s.client, "/customers", opts, func(ctx context.Context, path string, opts *model.CustomerListOptions) ([]model.Customer, *model.Meta, error) {
		var resp CustomerListResponse
		path = fmt.Sprintf("%s?page=%d&per_page=%d", path, opts.Page, opts.PerPage)
		if err := s.client.Get(ctx, path, &resp); err != nil {
			return nil, nil, err
		}
		return resp.Customers, resp.Meta, nil
	})
}

// Get retrieves a single customer by ID
func (s *Service) Get(ctx context.Context, id string) (*model.CustomerDetail, error) {
	var customer model.CustomerDetail
	path := fmt.Sprintf("/customers/%s", id)
	err := s.client.Get(ctx, path, &customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// Create creates a new customer
func (s *Service) Create(ctx context.Context, customer *model.Customer) (*model.Customer, error) {
	var createdCustomer model.Customer
	err := s.client.Post(ctx, "/customers", customer, &createdCustomer)
	if err != nil {
		return nil, err
	}
	return &createdCustomer, nil
}

// Update updates an existing customer
func (s *Service) Update(ctx context.Context, id string, customer *model.Customer) (*model.Customer, error) {
	var updatedCustomer model.Customer
	path := fmt.Sprintf("/customers/%s", id)
	err := s.client.Put(ctx, path, customer, &updatedCustomer)
	if err != nil {
		return nil, err
	}
	return &updatedCustomer, nil
}

// Delete deletes a customer
func (s *Service) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/customers/%s", id)
	return s.client.Delete(ctx, path)
}
