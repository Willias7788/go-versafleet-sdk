package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Willias7788/go-versafleet-sdk/client"
	"github.com/Willias7788/go-versafleet-sdk/config"
	"github.com/Willias7788/go-versafleet-sdk/jobs"
	"github.com/Willias7788/go-versafleet-sdk/tasks"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %s\n", err)
		os.Exit(1)
	}

	// 2. Initialize Client
	c := client.New(cfg)

	// 3. Authenticate
	ctx := context.Background()
	if err := c.Authenticate(ctx); err != nil {
		fmt.Printf("Authentication failed: %s\n", err)
		// handling for demo purposes; in real app, might retry or exit
		os.Exit(1)
	}
	fmt.Println("Successfully authenticated!")

	// 4. Use Jobs Service
	jobsService := jobs.New(c)

	// Create a job (Example)
	/*
	newJob := &jobs.Job{
		JobNumber: "JOB-12345",
		Status:    "active",
	}
	createdJob, err := jobsService.Create(ctx, newJob)
	if err != nil {
		log.Printf("Failed to create job: %v", err)
	} else {
		fmt.Printf("Created Job ID: %s\n", createdJob.ID)
	}
	*/

	// List Jobs
	fmt.Println("Listing Jobs:")
	iter := jobsService.List(ctx, client.ListOptions{PerPage: 5})
	for iter.Next() {
		job := iter.Value()
		fmt.Printf("- %s (ID: %s)\n", job.JobNumber, job.ID)
	}
	if err := iter.Err(); err != nil {
		log.Printf("Error during iteration: %v", err)
	}

	// 5. Use Tasks Service
	tasksService := tasks.New(c)
	fmt.Println("\nListing Tasks:")
	taskIter := tasksService.List(ctx, client.ListOptions{PerPage: 5})
	for taskIter.Next() {
		task := taskIter.Value()
		fmt.Printf("- Task %s (Type: %s)\n", task.ID, task.Type)
	}
}
