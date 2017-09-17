package common

import (
	"fmt"
	"math/rand"
	"path"
	"strings"

	"io/ioutil"
	"os"
)

func WriteToLog(text string) {
	fmt.Println(text)
	filename := path.Join("/home/" + os.Getenv("USER") + "/log")
	fmt.Println(filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		os.Create(filename)
	}

	defer f.Close()

	if _, err = f.WriteString(text + "\n"); err != nil {
		panic(err)
	}

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes returns a random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RelativePath returns a relative path from an absolute path
// (does not include .pouch in return)
func RelativePath(fp, pouchRoot string) string {
	parts := strings.Split(fp, pouchRoot)
	token := parts[len(parts)-1]
	return token
}

// AbsPath returns and absolute path from a relative path
// (does not include .pouch at the end)
func AbsPath(fp, pouchRoot string) string {
	return path.Join(pouchRoot, fp)
}

// DropTombstone writes an empty file ending with .pouch to disk
func DropTombstone(relPath string, cfg *Configuration) error {
	absPath := path.Join(cfg.PouchRoot, relPath)
	fullPath := absPath + ".pouch"
	fmt.Println("dropping tombstone at", fullPath)
	err := ioutil.WriteFile(fullPath, []byte(fullPath), os.ModePerm)
	return err
}

// IsPouchFile checks if a filepath is that of a pouch file
func IsPouchFile(fp string) bool {
	parts := strings.Split(fp, ".")
	fmt.Println("IsPouchFile parts", parts)
	token := parts[len(parts)-1]
	rslt := token == "pouch"
	fmt.Println(rslt)
	return rslt
}
