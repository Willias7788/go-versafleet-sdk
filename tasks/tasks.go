package tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/Willias7788/go-versafleet-sdk/client"
)

type Service struct {
	client *client.Client
}

func New(c *client.Client) *Service {
	return &Service{client: c}
}

type Task struct {
	ID          string    `json:"id"`
	JobID       string    `json:"job_id"`
	Type        string    `json:"type"` // e.g., "delivery", "collection"
	Status      string    `json:"status"`
	Address     string    `json:"address"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

type TaskListResponse struct {
	Tasks []Task       `json:"tasks"`
	Meta  *client.Meta `json:"meta"`
}

// List returns an iterator to list all tasks
func (s *Service) List(ctx context.Context, opts client.ListOptions) *client.Iterator[Task] {
	return client.NewIterator(ctx, s.client, "/tasks", opts, func(ctx context.Context, path string, opts client.ListOptions) ([]Task, *client.Meta, error) {
		var resp TaskListResponse
		path = fmt.Sprintf("%s?page=%d&per_page=%d", path, opts.Page, opts.PerPage)
		
		if err := s.client.Get(ctx, path, &resp); err != nil {
			return nil, nil, err
		}
		return resp.Tasks, resp.Meta, nil
	})
}

// Get retrieves a single task by ID
func (s *Service) Get(ctx context.Context, id string) (*Task, error) {
	var task Task
	path := fmt.Sprintf("/tasks/%s", id)
	err := s.client.Get(ctx, path, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
