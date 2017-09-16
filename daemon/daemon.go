package main

import (
	"fmt"
	"log"

	"github.com/JakeCooper/Pouch/common"
	"github.com/fsnotify/fsnotify"
)

var userHome = ""

var pouchRoot = ""

func tumbleEvents(fs common.CloudStorage, event fsnotify.Event) {
	switch event.Op {
	case fsnotify.Create:
		fs.Create(event.Name)
		fmt.Println("CREATED")
	case fsnotify.Chmod:
		fmt.Println("CHMOD")
	case fsnotify.Remove:
		fs.Delete(event.Name)
		fmt.Println("REMOVED")
	case fsnotify.Rename:
		fs.Update(event.Name)
		fmt.Println("RENAME")
	case fsnotify.Write:
		fs.Update(event.Name)
		fmt.Println("WRITE")
	default:
		fmt.Println("NONACTION DEFAULT")
	}
}

// RunDaemon is the main function for watching the file system
func RunDaemon(config *common.Configuration) {
	// usrHome, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }

	// create an instance of a cloud store
	cloudStore, err := common.InitCloudFS(config)
	if err != nil {
		panic(err)
	}

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
				tumbleEvents(cloudStore, event)
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
