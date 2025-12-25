package upload

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

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
	URL string `json:"signed_url"`
}

// GetPresignedURL helps to get the pre-signed URL for uploading a file
func (s *Service) GetPresignedURL(ctx context.Context, taskId, filename string) (*PresignedURLResponse, error) {
	var resp PresignedURLResponse
	// Note: Endpoint path is inferred. Please verify with official documentation.
	// Common patterns: /attachments/new, /files/storage_request
	path := fmt.Sprintf("/tasks/%s/presigned_url", taskId)

	_, err := s.client.R(ctx).SetQueryParam("filename", filename).SetResult(&resp).Get(path)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *Service) UploadBinaryToS3(ctx context.Context, url string, filename string, binary io.Reader) error {
	client := resty.New()

	// S3 often requires Content-Length. Resty sets it automatically for Files/Readers usually,
	resp, err := client.R().
		SetContext(ctx).
		SetFileReader("file", filename, binary).
		SetHeader("Content-Type", "application/octet-stream").
		Put(url)

	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("s3 upload failed with status: %d, body: %s", resp.StatusCode(), resp.String())
	}

	return nil
}

// UploadToS3 uploads the file content to the presigned URL
func (s *Service) UploadToS3(ctx context.Context, url string, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	fileName := filepath.Base(filePath)
	if fileName == "." || fileName == "/" {
		return fmt.Errorf("invalid file path: %s", filePath)
	}
	return s.UploadBinaryToS3(ctx, url, fileName, f)
}

// Upload performs the full upload flow: Get Presigned URL -> Upload to S3
func (s *Service) Upload(ctx context.Context, taskId, filePath string) (string, error) { // returns the S3 URL or Key
	// 1. Get Presigned URL
	// Extract filename from path
	// default to some logic
	fileName := filepath.Base(filePath)
	if fileName == "." || fileName == "/" {
		return "", fmt.Errorf("invalid file path: %s", filePath)
	}

	presigned, err := s.GetPresignedURL(ctx, taskId, fileName) // Simplify for now
	if err != nil {
		return "", fmt.Errorf("failed to get presigned url: %w", err)
	}

	// 2. Upload
	if err := s.UploadToS3(ctx, presigned.URL, filePath); err != nil {
		return "", err
	}

	return presigned.URL, nil
}
