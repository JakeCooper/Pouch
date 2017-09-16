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
		fileStorage.Create(event.Name)
		fmt.Println("CREATED")
	case fsnotify.Chmod:
		fmt.Println("CHMOD")
	case fsnotify.Remove:
		fileStorage.Delete(event.Name)
		fmt.Println("REMOVED")
	case fsnotify.Rename:
		fileStorage.Update(event.Name)
		fmt.Println("RENAME")
	case fsnotify.Write:
		fileStorage.Update(event.Name)
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

	fmt.Println("watching " + config.PouchRoot)
	err = watcher.Add(config.PouchRoot)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
