package drivers

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

type Driver struct {
	ID           int64  `json:"id"`
	GUID         string `json:"guid,omitempty"`
	Name         string `json:"name"`
	Phone        string `json:"phone"` // API uses contact_number usually, but kept Phone if mapped manually or common field
	Username     string `json:"username,omitempty"`
	VehicleID    string `json:"vehicle_id,omitempty"` // Note: API might send object or ID. Kept as string for now, but watch out.
	LicensePlate string `json:"license_plate,omitempty"`
	Status       string `json:"status,omitempty"`
	ExternalID   string `json:"external_id,omitempty"`
}

type DriverListResponse struct {
	Drivers []Driver    `json:"drivers"`
	Meta    *model.Meta `json:"meta"`
}

// List returns an iterator to list all drivers
func (s *Service) List(ctx context.Context, opts *model.ListOptions) *client.Iterator[Driver, *model.ListOptions] {
	return client.NewIterator(ctx, s.client, "/drivers", opts, func(ctx context.Context, path string, opts *model.ListOptions) ([]Driver, *model.Meta, error) {
		var resp DriverListResponse
		path = fmt.Sprintf("%s?page=%d&per_page=%d", path, opts.Page, opts.PerPage)
		if err := s.client.Get(ctx, path, &resp); err != nil {
			return nil, nil, err
		}
		return resp.Drivers, resp.Meta, nil
	})
}

// Get retrieves a single driver by ID
func (s *Service) Get(ctx context.Context, id string) (*Driver, error) {
	var driver Driver
	path := fmt.Sprintf("/drivers/%s", id)
	err := s.client.Get(ctx, path, &driver)
	if err != nil {
		return nil, err
	}
	return &driver, nil
}

// Create creates a new driver
func (s *Service) Create(ctx context.Context, driver *Driver) (*Driver, error) {
	var createdDriver Driver
	err := s.client.Post(ctx, "/drivers", driver, &createdDriver)
	if err != nil {
		return nil, err
	}
	return &createdDriver, nil
}

// Update updates an existing driver
func (s *Service) Update(ctx context.Context, id string, driver *Driver) (*Driver, error) {
	var updatedDriver Driver
	path := fmt.Sprintf("/drivers/%s", id)
	err := s.client.Put(ctx, path, driver, &updatedDriver)
	if err != nil {
		return nil, err
	}
	return &updatedDriver, nil
}

// Delete deletes a driver
func (s *Service) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/drivers/%s", id)
	return s.client.Delete(ctx, path)
}
