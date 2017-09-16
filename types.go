package main

import "net/http"

// CloudStorage : Interface for CloudStorage
type CloudStorage interface {
	create() HTTPStatus
	read() HTTPStatus
	update() HTTPStatus
	delete() HTTPStatus
}

// HTTPStatus : Status for HTTP, just an int
type HTTPStatus int

// AWS : Struct for Amazon Web Services
type AWS struct {
	Creds string
}

func (aws AWS) create() HTTPStatus {
	return http.StatusOK
}

func (aws AWS) read() HTTPStatus {
	return http.StatusOK
}

func (aws AWS) update() HTTPStatus {
	return http.StatusOK
}

func (aws AWS) delete() HTTPStatus {
	return http.StatusOK
}

func initCloudFS(storage CloudStorage) {

}
