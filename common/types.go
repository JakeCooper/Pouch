package common

import "net/http"
import "fmt"

// CloudStorage : Interface for CloudStorage
type CloudStorage interface {
	// Create uploads a new file to remote storage
	Create(path string) error
	// Get s a file from remote store and writes it into the supplied path
	Get(path string) error
	// Update a file that exists remote, path is local path to updated file
	Update(path string) error
	// Delete a file on remote fs
	Delete(path string) error
	// List all files on remote fs
	List() ([]string, error)
}

// HTTPStatus : Status for HTTP, just an int
type HTTPStatus int

// AWS : Struct for Amazon Web Services
type AWS struct {
	Creds string
}

// Create : Create for AWS S3
func (aws AWS) Create(path string) HTTPStatus {
	// myBucket, err := getS3Bucket(config.S3Root)
	// checkAndFailure(err)
	// myBucket.Put()
	fmt.Println(path)
	return http.StatusOK
}

// Read : Read for AWS S3
func (aws AWS) Read(path string) HTTPStatus {
	fmt.Println(path)
	return http.StatusOK
}

// Update : Update for AWS S3
func (aws AWS) Update(path string) HTTPStatus {
	fmt.Println(path)
	return http.StatusOK
}

// Delete : Delete for AWS S3
func (aws AWS) Delete(path string) HTTPStatus {
	fmt.Println(path)
	return http.StatusOK
}
