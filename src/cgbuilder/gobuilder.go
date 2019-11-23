package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func buildExecutable() {
	cmd := exec.Command(
		path.Join(Globals.GoInstallPath, "bin/go.exe"),
		"build",
		"-x",
		Globals.MainGoName,
	)

	gopath, _ := os.Getwd()

	cmd.Env = append(
		os.Environ(),
		"GOPATH="+gopath,
		"CGO_LDFLAGS=-L" + path.Join(GetAbsPath(gopath, Globals.TmpDirPath), "release") + " -lCuteGo", 
		"CGO_CFLAGS=-I" + path.Join(gopath, "qt/CuteGo/"),
	)

	out, err := cmd.CombinedOutput()
	
	fmt.Println(cmd.Args)

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}
}


func buildGo() {
	buildExecutable()
}
