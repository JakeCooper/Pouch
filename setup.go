package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
)

func checkAndFailure(err error) {
	if err != nil {
		fmt.Println("Logging failure: " + err.Error())
	}
}

func setup() {
	cmd := exec.Command("/bin/sh", "-c", "sh install.sh")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Setup failed")
	}
}

// Configuration : Config for loaded shit
type Configuration struct {
	S3Root    string
	PouchRoot string
}

var s3Bucket *s3.Bucket

var config Configuration

func loadSettings() Configuration {
	configuration := Configuration{}
	file, err := os.Open("./.settings.json")
	checkAndFailure(err)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if configuration.PouchRoot == "" {
		// User picks pouchroot
	}
	s3Bucket = nil
	if configuration.S3Root == "" {
		configuration.S3Root = generateBucketName()
		fmt.Printf("+%v", configuration)
		content, err := json.Marshal(configuration)
		checkAndFailure(err)
		err = ioutil.WriteFile("./.settings.json", content, 0644)
		checkAndFailure(err)
		s3Bucket = createS3Bucket(configuration.S3Root)
	} else {
		s3Bucket = getS3Bucket(configuration.S3Root)
	}

	return configuration
}

func pullFromCloud() {
	argsWithoutProg := os.Args[1:]
	file := argsWithoutProg[0]
	fmt.Println(file)
	// Get that file from the S3 bucket. Location probably dumped into the tombstone so it's ez to read.

}

func createS3Bucket(bucketName string) *s3.Bucket {
	fmt.Println(bucketName)

	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	fmt.Println(auth)

	conn := s3.New(auth, aws.USWest2)
	bucket := conn.Bucket(bucketName)

	// Need a new bucket
	err = bucket.PutBucket(s3.BucketOwnerFull)
	if err != nil {
		panic(err)
	}

	resp, err := bucket.List("", "", "", 100)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success!", resp)
	return bucket
}

func generateBucketName() string {
	return RandStringRunes(16)
}

func getS3Bucket(bucketName string) *s3.Bucket {
	// test if setup has completed
	// if _, err := os.Stat("~/.pouchconfig"); os.IsExist(err) {
	// 	fmt.Printf("Setup appears to be ready remove pouchconfig to continue")
	// 	return
	// }

	// Create a bucket using bound parameters and put something in it.
	// Make sure to change the bucket name from "myBucket" to something unique.

	fmt.Println(bucketName)

	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err)
	}

	fmt.Println(auth)

	conn := s3.New(auth, aws.USWest2)
	bucket := conn.Bucket(bucketName)

	// resp, err := bucket.List("", "", "", 100)
	// if err != nil {
	// 	panic(err)
	// }

	//fmt.Println("Success!", resp)
	return bucket
}

func init() {
	rand.Seed(time.Now().UnixNano())
	config = loadSettings()
}

func main() {
	// setup()
	fmt.Println(s3Bucket)
	pullFromCloud()
	daemon()
}
