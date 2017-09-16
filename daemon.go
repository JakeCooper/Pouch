package main

import (
	"fmt"
	"log"
	"os/user"

	"github.com/fsnotify/fsnotify"
)

var userHome = ""

var pouchRoot = ""

var fileStorage = InitCloudFS(AWS{
	Creds: config.S3Root,
})

func tumbleEvents(event fsnotify.Event) {
	switch event.Op {
	case fsnotify.Create:
		fileStorage.Create()
		fmt.Println("CREATED")
	case fsnotify.Chmod:
		fmt.Println("CHMOD")
	case fsnotify.Remove:
		fileStorage.Delete()
		fmt.Println("REMOVED")
	case fsnotify.Rename:
		fileStorage.Update()
		fmt.Println("RENAME")
	case fsnotify.Write:
		fileStorage.Update()
		fmt.Println("WRITE")
	default:
		fmt.Println("NONACTION DEFAULT")
	}
}

func daemon() {
	usrHome, err := user.Current()
	if err != nil {
		fmt.Println("FUCK")
	}
	userHome = usrHome.HomeDir

	// Watch for file changes in root file
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				tumbleEvents(event)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(pouchRoot)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
