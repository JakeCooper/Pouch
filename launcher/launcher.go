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

func heartBeat(file string) {
	common.WriteToLog("Starting heartBeat : " + file)
	for range time.Tick(time.Second * 5) {
		common.WriteToLog("getting: " + file)
	}
}

//r DownloadAndLaunch downloads the file coresponding to the pouch file, deletes
// the pouch file, and opens using the defaut program
func DownloadAndLaunch(fp string) {

	cfg := common.LoadSettings()

	common.WriteToLog("realfilePath: " + fmt.Sprintf("%v", cfg))

	common.WriteToLog("filePath: " + fp)
	go heartBeat(fp)

	fullPath := path.Join(currentPath(), fp)
	realFilePath := strings.Split(fullPath, ".pouch")[0]

	common.WriteToLog("realfilePath:" + realFilePath)
	if realFilePath != fullPath {
		common.WriteToLog("Going and getting the file")

		parts := strings.Split(realFilePath, cfg.PouchRoot)
		relPath := parts[len(parts)-1]

		bucket, err := common.GetS3Bucket(cfg.S3Root)
		if err != nil {
			common.WriteToLog("S3 Bucket was not found: " + cfg.S3Root)
			panic(err)
		}

		store := common.NewS3CloudStorage(&cfg, bucket)

		err = store.Get(relPath)
		if err != nil {
			common.WriteToLog("could not get the relPath from the store: " + relPath)
			panic(err)
		}

		err = os.Remove(fullPath)
		if err != nil {
			common.WriteToLog("could not remove: " + fullPath + " error: " + err.Error())
			//panic(err)
		}

	}

	common.WriteToLog(realFilePath)
	cmd := exec.Command("xdg-open", realFilePath)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	common.WriteToLog("Executed command")

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
