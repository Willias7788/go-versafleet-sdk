package jobs

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

type JobListResponse struct {
	Jobs []model.Job `json:"jobs"`
	Meta *model.Meta `json:"meta"`
}

// List returns an iterator to list all jobs
func (s *Service) List(ctx context.Context, opts *model.JobListOptions) *client.Iterator[model.Job, *model.JobListOptions] {
	return client.NewIterator(ctx, s.client, "/v2/jobs", opts, func(ctx context.Context, path string, opts *model.JobListOptions) ([]model.Job, *model.Meta, error) {
		var resp JobListResponse
		// Construct query params manually or use library (resty supports SetQueryParam)
		// Simulating simplistic approach here.
		// Detailed implementation would map opts to query params.

		path = fmt.Sprintf("%s?page=%d&per_page=%d", path, opts.Page, opts.PerPage)
		fmt.Println("List Job Path:", path)
		if err := s.client.Get(ctx, path, &resp); err != nil {
			return nil, nil, err
		}
		return resp.Jobs, resp.Meta, nil
	})
}

// Get retrieves a single job by ID
func (s *Service) Get(ctx context.Context, id string) (*model.Job, error) {
	var job model.Job
	path := fmt.Sprintf("/v2/jobs/%s", id)
	err := s.client.Get(ctx, path, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

// Create creates a new job
func (s *Service) Create(ctx context.Context, job *model.Job) (*model.Job, error) {
	var createdJob model.Job
	err := s.client.Post(ctx, "/v2/jobs", job, &createdJob)
	if err != nil {
		return nil, err
	}
	return &createdJob, nil
}
