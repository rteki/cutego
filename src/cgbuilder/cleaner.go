package main

import (
	"fmt"
	"os"
)

func rmIfPresent(path string) {

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		if err = os.Remove(path); err != nil {
			fmt.Println(err)
			os.Exit(666)
		}
	}

}

func cleanWorkspace() {

}
