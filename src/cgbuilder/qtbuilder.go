package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func buildQt(profile string) {
	out, err := exec.Command(
		path.Join(Globals.QtPath, "qmake.exe"),
		"CuteGo.pro",
		"-spec win32-g++",
		"\"CONFIG+=qtquickcompiler\""
	).Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}

	fmt.Println(out)

}
