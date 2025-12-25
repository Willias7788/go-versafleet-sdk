package upload

import (
	"context"
	"fmt"
	"os"

	"github.com/Willias7788/go-versafleet-sdk/client"
	"github.com/go-resty/resty/v2"
)

type Service struct {
	client *client.Client
}

func New(c *client.Client) *Service {
	return &Service{client: c}
}

type FileResponse struct {
	ID        string `json:"id"`
	URL       string `json:"url"`
	Filename  string `json:"filename"`
	MediaType string `json:"media_type"`
	Size      int64  `json:"size"`
}

type PresignedURLResponse struct {
	URL       string `json:"url"`
	Method    string `json:"method"` // usually PUT
	ExpiresIn int    `json:"expires_in"`
}

// GetPresignedURL helps to get the pre-signed URL for uploading a file
func (s *Service) GetPresignedURL(ctx context.Context, taskId, filename string) (*PresignedURLResponse, error) {
	var resp PresignedURLResponse
	// Note: Endpoint path is inferred. Please verify with official documentation.
	// Common patterns: /attachments/new, /files/storage_request
	path := fmt.Sprintf("tasks/%s/presigned_url", taskId)

	err := s.client.Get(ctx, path, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadToS3 uploads the file content to the presigned URL
func (s *Service) UploadToS3(ctx context.Context, url string, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// Create a clean Resty client to avoid inheriting auth headers
	// S3 will reject the request if unbeknownst Auth headers are present
	client := resty.New()

	// S3 often requires Content-Length. Resty sets it automatically for Files/Readers usually,
	// but for raw body with io.Reader we might want to be explicit or let Resty handle it.
	// We'll set the Content-Length header manually to be safe if Resty doesn't pick it up from File

	resp, err := client.R().
		SetContext(ctx).
		SetBody(f).
		SetHeader("Content-Type", "application/octet-stream").
		SetContentLength(true). // Force Content-Length header
		Put(url)

	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("s3 upload failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	return nil
}

// Upload performs the full upload flow: Get Presigned URL -> Upload to S3
func (s *Service) Upload(ctx context.Context, taskId, filePath string) (string, error) { // returns the S3 URL or Key
	// 1. Get Presigned URL
	// Extract filename from path
	// default to some logic

	presigned, err := s.GetPresignedURL(ctx, taskId, "filename.txt") // Simplify for now
	if err != nil {
		return "", fmt.Errorf("failed to get presigned url: %w", err)
	}

	// 2. Upload
	if err := s.UploadToS3(ctx, presigned.URL, filePath); err != nil {
		return "", err
	}

	return presigned.URL, nil
}
