package tasks

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

type TaskListResponse struct {
	Tasks []model.Task `json:"tasks"`
	Meta  *model.Meta  `json:"meta"`
}

// List returns an iterator to list all tasks
func (s *Service) List(ctx context.Context, opts *model.TaskListOptions) *client.Iterator[model.Task, *model.TaskListOptions] {
	return client.NewIterator(ctx, s.client, "/tasks", opts, func(ctx context.Context, path string, opts *model.TaskListOptions) ([]model.Task, *model.Meta, error) {
		var resp TaskListResponse
		path = fmt.Sprintf("%s?page=%d&per_page=%d", path, opts.Page, opts.PerPage)

		if err := s.client.Get(ctx, path, &resp); err != nil {
			return nil, nil, err
		}
		return resp.Tasks, resp.Meta, nil
	})
}

// Get retrieves a single task by ID
func (s *Service) Get(ctx context.Context, id string) (*model.Task, error) {
	var task model.Task
	path := fmt.Sprintf("/tasks/%s", id)
	err := s.client.Get(ctx, path, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
