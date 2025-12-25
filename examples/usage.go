package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// "github.com/Willias7788/go-versafleet-sdk/account"
	"github.com/Willias7788/go-versafleet-sdk/client"
	"github.com/Willias7788/go-versafleet-sdk/config"
	"github.com/Willias7788/go-versafleet-sdk/customers"
	"github.com/Willias7788/go-versafleet-sdk/drivers"
	"github.com/Willias7788/go-versafleet-sdk/jobs"
	"github.com/Willias7788/go-versafleet-sdk/model"
	"github.com/Willias7788/go-versafleet-sdk/tasks"
	"github.com/Willias7788/go-versafleet-sdk/upload"
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
	taskId := ""

	// 3. Verify Credentials (Login check)
	ctx := context.Background()
	if err := c.Verify(ctx); err != nil {
		fmt.Printf("Authentication failed: %s\n", err)
		// handling for demo purposes; in real app, might retry or exit
		os.Exit(1)
	}
	fmt.Println("Successfully authenticated (verified via /jobs)!")

	// 4. Use Jobs Service
	jobsService := jobs.New(c)

	// List Jobs
	fmt.Println("Listing Jobs:")
	jobOpt := model.JobListOptions{ListOptions: model.ListOptions{PerPage: 5}}
	iter := jobsService.List(ctx, &jobOpt)
	for iter.Next() {
		job := iter.Value()
		fmt.Printf("- %s (ID: %v)\n", job.JobType, job.ID)
	}
	if err := iter.Err(); err != nil {
		log.Printf("Error during iteration: %v", err)
	}

	// // 5. Use Tasks Service
	tasksService := tasks.New(c)
	fmt.Println("\nListing Tasks:")
	opt := model.TaskListOptions{ListOptions: model.ListOptions{PerPage: 5}}
	taskIter := tasksService.List(ctx, &opt)
	for taskIter.Next() {
		task := taskIter.Value()
		fmt.Printf("- Task %s (Type: %s)\n", task.ID, task.Job.JobType)
		taskId = fmt.Sprintf("%d", task.ID)
	}

	// 6. Use Drivers Service
	driversService := drivers.New(c)
	fmt.Println("\nListing Drivers:")
	drvOpt := model.ListOptions{PerPage: 5}
	driverIter := driversService.List(ctx, &drvOpt)
	for driverIter.Next() {
		drv := driverIter.Value()
		fmt.Printf("- Driver %s (Phone: %s)\n", drv.Name, drv.Phone)
	}

	// 7. Use Customers Service
	customersService := customers.New(c)
	fmt.Println("\nListing Customers:")
	custOpt := model.ListOptions{PerPage: 5}
	custIter := customersService.List(ctx, &custOpt)
	for custIter.Next() {
		cust := custIter.Value()
		fmt.Printf("- Customer %s (%s)\n", cust.Name, cust.Email)

	}

	// // 8. Use Account Service
	// accountService := account.New(c)
	// fmt.Println("\nShow Account:")
	// acc, err := accountService.Get(ctx, "")
	// if err != nil {
	// 	fmt.Printf("Error getting account: %v\n", err)
	// } else {
	// 	fmt.Printf("Account: %s (Email: %s)\n", acc.Name, acc.Email)
	// }
	// 8. Use Upload Service (Example)
	/*
		uploadService := upload.New(c)
		fileURL, err := uploadService.Upload(ctx, "./test.jpg")
		if err != nil {
			fmt.Printf("Upload failed: %v\n", err)
		} else {
			fmt.Printf("Uploaded file URL: %s\n", fileURL)
		}
	*/
	uploadService := upload.New(c)
	fmt.Println("\nUpload file:")
	fileURL, err := uploadService.Upload(ctx, taskId, "./test.txt")
	if err != nil {
		fmt.Printf("Upload failed: %v\n", err)
	} else {
		fmt.Printf("Uploaded file URL: %s\n", fileURL)
	}
}
