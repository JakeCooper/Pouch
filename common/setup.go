package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/pkg/errors"
)

var GlobalWatcher *fsnotify.Watcher

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

// LoadSettings returns the Pouch config settings
func LoadSettings() Configuration {
	configuration := Configuration{}
	base := path.Join(os.Getenv("HOME"), ".pouch")
	p := path.Join(base, "settings.json")

	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		err := os.Mkdir(base, os.ModePerm)
		if err != nil {
			panic(err)
		}
		configuration.S3Root = generateBucketName()
		configuration.PouchRoot = path.Join(os.Getenv("HOME"), "Pouch")
		fmt.Printf("+%v", configuration)

		content, err := json.Marshal(configuration)
		checkAndFailure(err)

		err = ioutil.WriteFile(p, content, 0644)
		_, err = CreateS3Bucket(configuration.S3Root)
		if err != nil {
			panic(err)
		}
	} else {
		file, err := os.Open(p)
		checkAndFailure(err)
		defer file.Close()

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&configuration)
	}
	return configuration
}

func generateBucketName() string {
	return RandStringRunes(16)
}

// CreateS3Bucket ..
func CreateS3Bucket(bucketName string) (*s3.Bucket, error) {
	bucket, err := GetS3Bucket(bucketName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create bucket")
	}
	// Need a new bucket
	err = bucket.PutBucket(s3.AuthenticatedRead)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create bucket")
	}
	return bucket, nil
}

// GetS3Bucket returns a new instance of a connection to an S3 bucket
func GetS3Bucket(bucketName string) (*s3.Bucket, error) {
	auth, err := GetAuth()
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to bucket")
	}

	conn := s3.New(*auth, aws.USWest2)
	bucket := conn.Bucket(bucketName)
	return bucket, nil
}

// GetAuth returns an AWS authenication instance. Note that most of the time
// functions expect a struct not a pointer so you'll need to de-reference the pointer
func GetAuth() (*aws.Auth, error) {
	if os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == "" {
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

// CreatePouch : Create Pouch
func CreatePouch(config *Configuration) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	config.PouchRoot = strings.Replace(config.PouchRoot, "~", usr.HomeDir, -1) + "/"
	os.MkdirAll(config.PouchRoot, os.ModePerm)
	loadPouch(config)
}

func loadPouch(config *Configuration) error {
	myBucket, err := GetS3Bucket(config.S3Root)
	if err != nil {
		return err
	}
	fmt.Println(myBucket.Name)
	myFiles, err := myBucket.GetBucketContents()
	if err != nil {
		return err
	}
	for file := range *myFiles {
		fmt.Printf("%s\n", config.PouchRoot+file)
		if string(file[len(file)-1]) == "/" {
			// directory
			os.MkdirAll(config.PouchRoot+file, os.ModePerm)
			GlobalWatcher.Add(config.PouchRoot + file)
		} else {
			// file
			filePtr, err := os.Create(config.PouchRoot + file + ".pouch")
			if err != nil {
				return err
			}
			filePtr.WriteString(config.PouchRoot + file)
		}
	}
	return nil
}
