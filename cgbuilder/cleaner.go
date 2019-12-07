package main

import (
	"fmt"
	"os"
	"path"
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
	rmIfPresent(path.Join(Globals.QrcRoot, "resources.qrc"))
	os.RemoveAll(Globals.TmpDirPath)
}

func cleanBuildDest() {
	os.RemoveAll(Globals.BuildDest)
}

