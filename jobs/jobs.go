package jobs

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

// Job represents a VersaFleet job
type Job struct {
	ID             string `json:"id"`
	JobNumber      string `json:"job_number"`
	Status         string `json:"status"`
	DriverID       string `json:"driver_id,omitempty"`
	VehicleID      string `json:"vehicle_id,omitempty"`
	Customer       string `json:"customer,omitempty"`
	Note           string `json:"note,omitempty"`
	// Add other fields as per API spec
}

type JobListResponse struct {
	Jobs []Job        `json:"jobs"`
	Meta *client.Meta `json:"meta"`
}

// List returns an iterator to list all jobs
func (s *Service) List(ctx context.Context, opts client.ListOptions) *client.Iterator[Job] {
	return client.NewIterator(ctx, s.client, "/jobs", opts, func(ctx context.Context, path string, opts client.ListOptions) ([]Job, *client.Meta, error) {
		var resp JobListResponse
		// Construct query params manually or use library (resty supports SetQueryParam)
		// Simulating simplistic approach here.
		// Detailed implementation would map opts to query params.
		
		path = fmt.Sprintf("%s?page=%d&per_page=%d", path, opts.Page, opts.PerPage)
		
		if err := s.client.Get(ctx, path, &resp); err != nil {
			return nil, nil, err
		}
		return resp.Jobs, resp.Meta, nil
	})
}

// Get retrieves a single job by ID
func (s *Service) Get(ctx context.Context, id string) (*Job, error) {
	var job Job
	path := fmt.Sprintf("/jobs/%s", id)
	err := s.client.Get(ctx, path, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// Create creates a new job
func (s *Service) Create(ctx context.Context, job *Job) (*Job, error) {
	var createdJob Job
	err := s.client.Post(ctx, "/jobs", job, &createdJob)
	if err != nil {
		return nil, err
	}
	return &createdJob, nil
}
