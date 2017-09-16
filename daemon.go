package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

var fileStorage = initCloudFS(AWS{
	Creds: "JakeCooper",
})

func tumbleEvents(event fsnotify.Event) {
	switch event.Op {
	case fsnotify.Create:
		fileStorage.create()
		fmt.Println("CREATED")
	case fsnotify.Chmod:
		fmt.Println("CHMOD")
	case fsnotify.Remove:
		fileStorage.delete()
		fmt.Println("REMOVED")
	case fsnotify.Rename:
		fileStorage.update()
		fmt.Println("RENAME")
	case fsnotify.Write:
		fileStorage.update()
		fmt.Println("WRITE")
	default:
		fmt.Println("NONACTION DEFAULT")
	}
}

func daemon() {
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

	err = watcher.Add("/tmp/test")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
