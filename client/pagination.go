package client

import (
	"context"

	"github.com/Willias7788/go-versafleet-sdk/model"
)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptionsooo struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// Iterator is a helper to iterate over pages of results
type Iterator[T any, O model.Paginatable] struct {
	ctx          context.Context
	client       *Client
	path         string
	listOptions  O
	items        []T
	currentIndex int
	totalItems   int
	err          error
	fetchFunc    func(context.Context, string, O) ([]T, *Meta, error)
	meta         *Meta
}

type Meta struct {
	TotalEntries int `json:"total_entries"`
	TotalPages   int `json:"total_pages"`
	CurrentPage  int `json:"current_page"`
	PerPage      int `json:"per_page"`
}

// NewIterator creates a new iterator.
// fetchFunc is a closure that calls the specific API endpoint.
func NewIterator[T any, O model.Paginatable](
	ctx context.Context,
	client *Client,
	path string,
	opts O,
	fetchFunc func(context.Context, string, O) ([]T, *Meta, error),
) *Iterator[T, O] {
	if opts.GetPage() == 0 {
		opts.SetPage(1)
	}
	if opts.GetPerPage() == 0 {
		opts.SetPerPage(20)
	}

	return &Iterator[T, O]{
		ctx:         ctx,
		client:      client,
		path:        path,
		listOptions: opts,
		fetchFunc:   fetchFunc,
	}
}

// Next returns the next item in the iteration.
// It returns true if there is a next item, false otherwise.
func (it *Iterator[T, O]) Next() bool {
	if it.err != nil {
		return false
	}

	// detailed logic:
	// if we have items and index is within range, use it.
	// if index is at end of items, check if we can fetch more.

	if it.currentIndex < len(it.items) {
		it.currentIndex++
		return true
	}

	// Need to fetch more?
	if it.meta != nil && it.listOptions.GetPage() >= it.meta.TotalPages {
		return false // No more pages
	}

	// If it's not the first run, increment page
	if it.meta != nil {
		it.listOptions.SetPage(it.listOptions.GetPage() + 1)
	}

	items, meta, err := it.fetchFunc(it.ctx, it.path, it.listOptions)
	if err != nil {
		it.err = err
		return false
	}

	if len(items) == 0 {
		return false
	}

	it.items = items
	it.meta = meta
	it.currentIndex = 1 // 1-based "current" but we access with -1
	return true
}

// Value returns the current item.
func (it *Iterator[T, O]) Value() T {
	if len(it.items) == 0 || it.currentIndex-1 < 0 || it.currentIndex-1 >= len(it.items) {
		var zero T
		return zero
	}
	return it.items[it.currentIndex-1]
}

// Err returns any error that occurred during iteration.
func (it *Iterator[T, O]) Err() error {
	return it.err
}
