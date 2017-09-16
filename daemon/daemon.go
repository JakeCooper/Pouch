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
				event.Name = common.RelativePath(event.Name, config.PouchRoot)
				tumbleEvents(cloudStore, event, config)
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
