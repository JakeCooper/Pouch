package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/JakeCooper/Pouch/common"
	"github.com/fsnotify/fsnotify"
)

func tumbleEvents(fs common.CloudStorage, event fsnotify.Event) {
	switch event.Op {
	case fsnotify.Create:
		fmt.Println("CREATE")
		fs.Create(event.Name)
	case fsnotify.Chmod:
	case fsnotify.Remove:
		fmt.Println("REMOVE")
		fs.Delete(event.Name)
	case fsnotify.Rename:
		fmt.Println("RENAME")
		fs.Delete(event.Name)
	case fsnotify.Write:
		fmt.Println("WRITE")
		fs.Update(event.Name)
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

	bucket, err := common.GetS3Bucket(config.S3Root)
	if err != nil {
		panic(err)
	}

	cloudStore := common.NewS3CloudStorage(config, bucket)

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
				event.Name = strings.TrimLeft(event.Name, config.PouchRoot)
				event.Name = strings.TrimRight(event.Name, ".pouch")
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

func main() {
	rand.Seed(time.Now().UnixNano())
	settings := common.LoadSettings()
	common.CreatePouch(&settings)
	RunDaemon(&settings)
}
