package main

import (
	"fmt"
	"os"
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
