package main

import "net/http"

// CloudStorage : Interface for CloudStorage
type CloudStorage interface {
	Create() HTTPStatus
	Read() HTTPStatus
	Update() HTTPStatus
	Delete() HTTPStatus
}

// HTTPStatus : Status for HTTP, just an int
type HTTPStatus int

// AWS : Struct for Amazon Web Services
type AWS struct {
	Creds string
}

func (aws AWS) Create() HTTPStatus {
	return http.StatusOK
}

func (aws AWS) Read() HTTPStatus {
	return http.StatusOK
}

func (aws AWS) Update() HTTPStatus {
	return http.StatusOK
}

func (aws AWS) Delete() HTTPStatus {
	return http.StatusOK
}
