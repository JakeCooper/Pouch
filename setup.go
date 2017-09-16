package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "sh install.sh")
	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Setup failed")
	}
}
