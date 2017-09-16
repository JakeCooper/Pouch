package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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

var config = loadSettings()

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
		// Generate them a new S3 bucket
	}
	return configuration
}

func pullFromCloud() {
	argsWithoutProg := os.Args[1:]
	file := argsWithoutProg[0]
	fmt.Println(file)
	// Get that file from the S3 bucket. Location probably dumped into the tombstone so it's ez to read.

}

func main() {
	// setup()
	loadSettings()
	pullFromCloud()
	daemon()
}
