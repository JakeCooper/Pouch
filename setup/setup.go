package main

import (
	"fmt"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"

	"os"
)

func setup() {
	// test if setup has completed
	if _, err := os.Stat("~/.pouchconfig"); os.IsExist(err) {
		fmt.Printf("Setup appears to be ready remove pouchconfig to continue")
		return
	}

	// Create a bucket using bound parameters and put something in it.
	// Make sure to change the bucket name from "myBucket" to something unique.
	bucketName := "Pouch123" + os.Getenv("USER")
	fmt.Println(bucketName)

	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	fmt.Println(auth)

	conn := s3.New(auth, aws.USWest2)
	bucket := conn.Bucket(bucketName)
	err = bucket.PutBucket(s3.BucketOwnerFull)
	if err != nil {
		panic(err)
	}

	resp, err := bucket.List("", "", "", 100)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success!", resp)
}

func main() {
	setup()
	// daemon()
}
