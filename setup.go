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
	"github.com/pkg/errors"
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
	if configuration.S3Root == "" {
		configuration.S3Root = generateBucketName()
		fmt.Printf("+%v", configuration)

		content, err := json.Marshal(configuration)
		checkAndFailure(err)

		err = ioutil.WriteFile("./.settings.json", content, 0644)
		_, err = createS3Bucket(configuration.S3Root)
		if err != nil {
			panic(err)
		}
	} else {
		_, err := getS3Bucket(configuration.S3Root)
		if err != nil {
			panic(err)
		}
	}

	return configuration
}

func pullFromCloud() {
	argsWithoutProg := os.Args[1:]
	file := argsWithoutProg[0]
	fmt.Println(file)
	// Get that file from the S3 bucket. Location probably dumped into the tombstone so it's ez to read.

}

func generateBucketName() string {
	return RandStringRunes(16)
}

func createS3Bucket(bucketName string) (*s3.Bucket, error) {
	bucket, err := getS3Bucket(bucketName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create bucket")
	}
	// Need a new bucket
	err = bucket.PutBucket(s3.BucketOwnerFull)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create bucket")
	}
	return bucket, nil
}

func getS3Bucket(bucketName string) (*s3.Bucket, error) {
	fmt.Println(bucketName)

	auth, err := getAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to bucket")
	}

	conn := s3.New(*auth, aws.USWest2)
	bucket := conn.Bucket(bucketName)
	return bucket, nil
}

func getAuth() (*aws.Auth, error) {
	if os.Getenv("AWS_ACCESS_KEY") == "" || os.Getenv("AWS_SECRET_KEY") == "" {
		auth, err := aws.SharedAuth()
		if err != nil {
			return nil, errors.Wrap(err, "could not authenticate through evars or through creds file")
		}
		return &auth, nil
	}
	auth, err := aws.EnvAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not auth from envars")
	}
	return &auth, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
	config = loadSettings()
}

func main() {
	// setup()
	pullFromCloud()
	daemon()
}
