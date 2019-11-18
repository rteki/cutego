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
		Globals.MainGoName,
	)

	gopath, _ := os.Getwd()

	cmd.Env = append(
		os.Environ(),
		"GOPATH="+gopath,
	)

	out, err := cmd.Output()

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
