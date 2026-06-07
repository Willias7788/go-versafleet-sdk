package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	customerId := -1
	billingAccountId := -1

	// 3. Verify Credentials (Login check)
	ctx := context.Background()
	if err := c.Verify(ctx); err != nil {
		fmt.Printf("Authentication failed: %s\n", err)
		// handling for demo purposes; in real app, might retry or exit
		os.Exit(1)
	}
	fmt.Println("Successfully authenticated (verified via /jobs)!")
	runJob := false
	runTask := false
	runDriver := false
	runCustomer := true
	runUpload := false
	runCreateJob := false
	runUpdateTask := false
	jobService := jobs.New(c)
	tasksService := tasks.New(c)
	driversService := drivers.New(c)
	customersService := customers.New(c)

	// 4. Use Jobs Service
	if runJob {
		fmt.Println("Listing Jobs:")
		jobOpt := model.JobListOptions{}
		jobOpt.PerPage = 5
		iter := jobService.List(ctx, &jobOpt)
		for iter.Next() {
			job := iter.Value()
			taskId = fmt.Sprintf("%v", job.BaseTask.ID)
			fmt.Printf("- %s (ID: %v), TaskId: %v\n", job.JobType, job.ID, job.BaseTask.ID)
			break
		}
		if err := iter.Err(); err != nil {
			log.Printf("Error during iteration: %v", err)
		}
	}

	// 5. Use Tasks Service
	if runTask {
		fmt.Println("\nListing Tasks:")
		opt := model.TaskListOptions{}
		opt.PerPage = 5
		taskIter := tasksService.List(ctx, &opt)
		for taskIter.Next() {
			task := taskIter.Value()
			// taskId = fmt.Sprintf("%v", task.ID)
			fmt.Printf("- Task %d %s (State: %s)\n", task.ID, task.StateUpdatedAt, task.State)
			break
		}
	}

	// 6. Use Drivers Service
	if runDriver {
		fmt.Println("\nListing Drivers:")
		drvOpt := model.ListOptions{PerPage: 5}
		driverIter := driversService.List(ctx, &drvOpt)
		for driverIter.Next() {
			drv := driverIter.Value()
			fmt.Printf("- Driver %s (Phone: %s)\n", drv.Name, drv.Phone)
			break
		}
	}

	// 7. Use Customers Service
	if runCustomer {
		// Test Create Customer
		fmt.Println("\nCreating Customer:")
		newCust := model.Customer{
			Name:          fmt.Sprintf("Test Customer %d", time.Now().Unix()),
			Email:         fmt.Sprintf("test_customer_%d@example.com", time.Now().Unix()),
			ContactPerson: "Test Contact",
			ContactNumber: "91234567",
		}
		createdCust, err := customersService.Create(ctx, &newCust)
		if err != nil {
			log.Printf("Error creating customer: %v", err)
		} else {
			fmt.Printf("Successfully created customer: %s (ID: %d)\n", createdCust.Name, createdCust.ID)
		}

		fmt.Println("\nListing Customers:")
		custOpt := model.CustomerListOptions{}
		custOpt.PerPage = 5
		custIter := customersService.List(ctx, &custOpt)
		for custIter.Next() {
			cust := custIter.Value()
			fmt.Printf("- Customer %s (%s)\n", cust.Name, cust.Email)
			customerId = cust.ID
			break

		}
	}

	if runCreateJob {
		customerId = 30293
		billingAccountId = 54459

		timeFrom := time.Now().Format("2006-01-02 15:04:05")
		timeTo := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
		task := model.TaskParams{
			Price:       20,
			ExpectedCod: 20,
			TimeType:    model.TimeTypeAllDay,
			TrackingID:  "xxsxsxsxs123456",
			TimeFrom:    &timeFrom,
			TimeTo:      &timeTo,
			AddressAttributes: &model.Address{
				Line1:   "Anywhere",
				City:    "Singapore",
				Country: "Singapore",
				Zip:     "610254",
			},
		}
		fmt.Println("\nCreating Job:")
		timeType := model.TimeTypeAllDay
		job := model.JobParams{
			BaseTaskAttributes: &model.BaseTaskParams{
				TimeFrom:         &timeFrom,
				TimeTo:           &timeTo,
				TimeType:         &timeType,
				BillingAccountID: &billingAccountId,
				AddressAttributes: &model.Address{
					Line1:   "2 Pandan Road",
					City:    "Singapore",
					Country: "Singapore",
					Zip:     "609254",
				},
			},

			JobType:    "delivery",
			CustomerID: customerId,
			TasksAttributes: []model.TaskParams{
				task,
			},
		}

		_ = billingAccountId
		jobCreated, err := jobService.Create(ctx, &job)
		if err != nil {
			fmt.Printf("Error creating job: %v\n", err)
		} else {
			fmt.Printf("Job created: %s (ID: %v)\n", jobCreated.JobType, jobCreated.ID)
		}
	}

	if runUpdateTask {
		taskId := "28943974"
		billingAccountId = 54459
		fmt.Println("\nUpdating Job:")
		// inv := "123466"
		task := model.TaskParams{
			InvoiceNumber:    "123466",
			Remarks:          "Updated Remarks",
			BillingAccountID: &billingAccountId,
		}
		taskUpdated, err := tasksService.Update(ctx, taskId, &task)
		if err != nil {
			fmt.Printf("Error updating job: %v\n", err)
		} else {
			fmt.Printf("Job updated: %s (ID: %v)\n", taskUpdated.InvoiceNumber, taskUpdated.ID)
		}
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
	if runUpload {
		uploadService := upload.New(c)
		fmt.Println("\nUpload file:")
		fileURL, err := uploadService.Upload(ctx, taskId, "./test.csv")
		if err != nil {
			fmt.Printf("Upload failed: %v\n", err)
		} else {
			fmt.Printf("Uploaded file URL: %s\n", fileURL)
		}
	}
}
