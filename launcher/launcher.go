package launcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/JakeCooper/Pouch/common"
)

func WriteToLog(text string) {
	fmt.Println(text)
	filename := path.Join(os.Getenv("HOME"), ".pouch") + "/log"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text + "\n"); err != nil {
		panic(err)
	}

}

func heartBeat(file string) {
	WriteToLog("Starting heartBeat : " + file)
	for range time.Tick(time.Second * 5) {
		WriteToLog("getting: " + file)
	}
}

//r DownloadAndLaunch downloads the file coresponding to the pouch file, deletes
// the pouch file, and opens using the defaut program
func DownloadAndLaunch(fp string) {

	cfg := common.LoadSettings()

	WriteToLog("realfilePath: " + fmt.Sprintf("%v", cfg))

	WriteToLog("filePath: " + fp)
	go heartBeat(fp)

	fullPath := path.Join(currentPath(), fp)
	realFilePath := strings.Split(fullPath, ".pouch")[0]

	WriteToLog("realfilePath:" + realFilePath)
	if realFilePath != fullPath {
		WriteToLog("Going and getting the file")

		parts := strings.Split(realFilePath, cfg.PouchRoot)
		relPath := parts[len(parts)-1]

		bucket, err := common.GetS3Bucket(cfg.S3Root)
		if err != nil {
			WriteToLog("S3 Bucket was not found: " + cfg.S3Root)
			panic(err)
		}

		store := common.NewS3CloudStorage(&cfg, bucket)

		err = store.Get(relPath)
		if err != nil {
			WriteToLog("could not get the relPath from the store: " + relPath)
			panic(err)
		}

		err = os.Remove(fullPath)
		if err != nil {
			WriteToLog("could not remove: " + fullPath + " error: " + err.Error())
			//panic(err)
		}

	}

	WriteToLog(realFilePath)
	cmd := exec.Command("xdg-open", realFilePath)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	WriteToLog("Executed command")

	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
}

func currentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}
