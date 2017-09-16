package main

import "net/http"
import "fmt"

// CloudStorage : Interface for CloudStorage
type CloudStorage interface {
	Create(path string) HTTPStatus
	Read(path string) HTTPStatus
	Update(path string) HTTPStatus
	Delete(path string) HTTPStatus
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
