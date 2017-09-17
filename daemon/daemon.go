package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"sync"

	"github.com/JakeCooper/Pouch/common"
	"github.com/fsnotify/fsnotify"
)

func logIfErr(err error) {
	if err != nil {
		log.Println("[ERROR]", err)
	}
}

type eventHandler struct {
	mu  sync.Mutex
	fs  common.CloudStorage
	cfg *common.Configuration
}

func (e *eventHandler) create(relPath string) {
	if !common.IsPouchFile(relPath) {
		e.mu.Lock()
		defer e.mu.Unlock()
		err := e.fs.Create(relPath)
		logIfErr(err)
	}
}

func (e *eventHandler) delete(fp string) {
	fmt.Println("deleting? file:", strings.Split(fp, ".pouch"))
	e.mu.Lock()
	if !common.IsPouchFile(fp) {
		// This is a real file
		fmt.Println(fp, "is a real file, droping tombstone")
		err := common.DropTombstone(fp, e.cfg)
		logIfErr(err)
	} else {
		fmt.Println("Deleting tombstone")
		err := e.fs.Delete(fp)
		logIfErr(err)
	}
	e.mu.Unlock()
}

func tumbleEvents(eh *eventHandler, event fsnotify.Event) {
	switch event.Op {
	case fsnotify.Create:
		fmt.Println("CREATE, file is:", event.Name)
		// eh.create(event.Name)
	case fsnotify.Chmod:
		fmt.Println("CHMOD")
	case fsnotify.Remove:
		fmt.Println("REMOVE, old file is:", event.Name)
		eh.delete(event.Name)
	case fsnotify.Rename:
		fmt.Println("RENAME, old file is:", event.Name)
		eh.delete(event.Name)
	case fsnotify.Write:
		fmt.Println("WRITE, old file is:", event.Name)
		eh.create(event.Name)
	default:
		fmt.Println("EVENT:", event.String())
		fmt.Println("NONACTION DEFAULT")
	}
}

// RunDaemon is the main function for watching the file system
func RunDaemon(config *common.Configuration) {

	bucket, err := common.GetS3Bucket(config.S3Root)
	if err != nil {
		panic(err)
	}

	cloudStore := common.NewS3CloudStorage(config, bucket)

	// Watch for file changes in root file
	if err != nil {
		log.Fatal(err)
	}
	defer common.GlobalWatcher.Close()

	eh := &eventHandler{
		mu:  sync.Mutex{},
		fs:  cloudStore,
		cfg: config,
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-common.GlobalWatcher.Events:
				event.Name = common.RelativePath(event.Name, config.PouchRoot)
				tumbleEvents(eh, event)
			case err := <-common.GlobalWatcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	fmt.Println("watching " + config.PouchRoot)
	err = common.GlobalWatcher.Add(config.PouchRoot)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	rand.Seed(time.Now().UnixNano())
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	common.GlobalWatcher = watcher

	settings := common.LoadSettings()
	common.CreatePouch(&settings)
	RunDaemon(&settings)
}
