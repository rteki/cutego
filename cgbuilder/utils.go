package main

import (
	"fmt"
	"os"
	"path"
)

func ExecInDir(path string, callback func()) {
	prevdir, err := os.Getwd()

	if err != nil {
		fmt.Println(prevdir)
		fmt.Println(err)
		os.Exit(666)
	}

	os.Chdir(path)

	callback()

	os.Chdir(prevdir)
}

func CreateDir(dirpath string) {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		if err = os.Mkdir(dirpath, 0644); err != nil {
			fmt.Println("Failed to create build destination directory!")
			fmt.Println(err)
		}
	}
}

func GetAbsPath(current string, destination string) string {
	if path.IsAbs(destination) {
		return destination
	} else {
		return path.Join(current, destination)
	}
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}
