package main

import (
	"os"

	"github.com/JakeCooper/Pouch/daemon"
	"github.com/JakeCooper/Pouch/launcher"
)

func main() {
	if len(os.Args) == 1 {
		daemon.Start()
	}
	if len(os.Args) == 2 {
		launcher.DownloadAndLaunch(os.Args[1])
	} else {
		panic("TOO MANY ARGS")
	}
}
