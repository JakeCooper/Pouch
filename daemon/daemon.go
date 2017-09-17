package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"sync"

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
	RunPoller(config)
	<-done
}

func glob(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func stringNotInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return false
		}
	}
	return true
}

// RunPoller polls CloudStorage
func RunPoller(config *common.Configuration) {
	go func() {
		for {
			resp, err := http.Get("https://smpzbbu1uk.execute-api.us-west-2.amazonaws.com/prod/pouch_getmetadata")
			if err != nil {
				log.Fatal(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			var res []common.Metadata
			err = json.Unmarshal(body, &res)
			if err != nil {
				log.Fatal(err)
			}

			// Add new files
			remoteFiles := map[string]bool{}
			removeFileArray := []string{}
			for _, file := range res {
				_, errNoPouch := os.Stat(config.PouchRoot + file.FilePath)
				_, errPouch := os.Stat(config.PouchRoot + file.FilePath + ".pouch")

				if os.IsNotExist(errPouch) && os.IsNotExist(errNoPouch) {
					// Make a pouch tombstone/folder for it
					if file.ObjectType == "folder" {
						os.MkdirAll(config.PouchRoot+file.FilePath, os.ModePerm)
					} else {
						common.DropTombstone(file.FilePath, config)
					}
				}
				localFilePath := config.PouchRoot + file.FilePath
				remoteFiles[strings.ToLower(localFilePath)] = true
			}

			filepath.Walk(config.PouchRoot, func(path string, f os.FileInfo, err error) error {
				filePath := path
				if f.IsDir() {
					filePath += "/"
				} else {
					filePath = strings.Split(filePath, ".pouch")[0]
				}
				if remoteFiles[strings.ToLower(filePath)] == false && path != config.PouchRoot {
					removeFileArray = append(removeFileArray, path)
				}
				return nil
			})

			for _, filePath := range removeFileArray {
				os.Remove(filePath)
			}

			time.Sleep(time.Second * 2)
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
}
