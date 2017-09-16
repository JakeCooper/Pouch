package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
	"strings"
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

func generateBucketName() string {
	return RandStringRunes(16)
}

func createS3Bucket(bucketName string) *s3.Bucket {
	fmt.Println(bucketName)

	auth, err := aws.EnvAuth()
	checkAndFailure(err)

	conn := s3.New(auth, aws.USWest2)
	bucket := conn.Bucket(bucketName)

	// Need a new bucket
	err = bucket.PutBucket(s3.BucketOwnerFull)
	checkAndFailure(err)

	return bucket
}

func getS3Bucket(bucketName string) *s3.Bucket {
	fmt.Println(bucketName)

	auth, err := aws.EnvAuth()
	checkAndFailure(err)

	conn := s3.New(auth, aws.USWest2)
	bucket := conn.Bucket(bucketName)
	return bucket
}

func createPouch() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.PouchRoot)
	config.PouchRoot = strings.Replace(config.PouchRoot, "~", usr.HomeDir, -1) + "/"
	os.MkdirAll(config.PouchRoot, os.ModePerm)
	loadPouch()
}

func loadPouch() {
	myFiles, err := s3Bucket.GetBucketContents()
	checkAndFailure(err)
	for file := range *myFiles {
		fmt.Printf("%s\n", config.PouchRoot+file)
		if string(file[len(file)-1]) == "/" {
			// directory
			os.MkdirAll(config.PouchRoot+file, os.ModePerm)
		} else {
			// file
			filePtr, err := os.Create(config.PouchRoot + file + ".pouch")
			checkAndFailure(err)
			filePtr.WriteString(config.PouchRoot + file)
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	config = loadSettings()
	createPouch()
}

func main() {
	// setup()
	fmt.Println(s3Bucket)
	pullFromCloud()
	daemon()
}
