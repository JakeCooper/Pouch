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

type CloudObject struct {
	Path string
}

var FileSystemChannel = make(chan CloudObject)

func logIfErr(err error) {
	if err != nil {
		log.Println("[ERROR]", err)
	}
}

func tumbleEvents(fs common.CloudStorage, event fsnotify.Event, cfg *common.Configuration) {
	switch event.Op {
	case fsnotify.Create:
		fmt.Println("CREATE")
		handleCreate(event.Name, fs)
	case fsnotify.Chmod:
	case fsnotify.Remove:
		fmt.Println("REMOVE")
		handleDelete(event.Name, fs, cfg)
	case fsnotify.Rename:
		fmt.Println("RENAME")
		// err := fs.Delete(event.Name)
		// logIfErr(err)
		handleDelete(event.Name, fs, cfg)

	case fsnotify.Write:
		fmt.Println("WRITE")
		handleUpdate(event.Name, fs)
	default:
		fmt.Println("NONACTION DEFAULT")
	}
}

func handleCreate(relPath string, fs common.CloudStorage) {
	if !common.IsPouchFile(relPath) {
		err := fs.Create(relPath)
		logIfErr(err)
	}
}

func handleDelete(fp string, fs common.CloudStorage, cfg *common.Configuration) {
	fmt.Println(strings.Split(fp, ".pouch"))
	if !common.IsPouchFile(fp) {
		// This is a pouchfile
		fmt.Println("Gonna drop a tombstone")
		err := common.DropTombstone(fp, cfg)
		logIfErr(err)
	} else {
		err := fs.Delete(fp)
		logIfErr(err)

	}
}

func handleUpdate(relPath string, fs common.CloudStorage) {
	if !common.IsPouchFile(relPath) {
		err := fs.Update(relPath)
		logIfErr(err)
	}
}

func updateFileSystem(newObject CloudObject) {
	fmt.Printf("New File Added! %+v\n", newObject)
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

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-common.GlobalWatcher.Events:
				event.Name = common.RelativePath(event.Name, config.PouchRoot)
				tumbleEvents(cloudStore, event, config)
			case newObject := <-FileSystemChannel:
				updateFileSystem(newObject)
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

// RunPoller polls CloudStorage
func RunPoller(config *common.Configuration) {
	go func() {
		for {

			time.Sleep(time.Second * 10)
		}
	}()
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
	RunPoller(&settings)
}
