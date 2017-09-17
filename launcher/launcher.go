package launcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/JakeCooper/Pouch/common"
)

// DownloadAndLaunch downloads the file coresponding to the pouch file, deletes
// the pouch file, and opens using the defaut program
func DownloadAndLaunch(fp string) {
	cfg := common.LoadSettings()
	fmt.Println(cfg)

	fullPath := path.Join(currentPath(), fp)
	realFilePath := strings.Split(fullPath, ".pouch")[0]
	parts := strings.Split(realFilePath, cfg.PouchRoot)
	relPath := parts[len(parts)-1]

	bucket, err := common.GetS3Bucket(cfg.S3Root)
	if err != nil {
		panic(err)
	}

	store := common.NewS3CloudStorage(&cfg, bucket)

	err = store.Get(relPath)
	if err != nil {
		panic(err)
	}

	err = os.Remove(fullPath)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("xdg-open", realFilePath)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
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
