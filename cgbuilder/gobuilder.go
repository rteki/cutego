package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

func buildExecutable() {
	_, currentFilePath, _, _ := runtime.Caller(0)

	cutegoQtSrcRoot = filepath.Join(filepath.Dir(currentFilePath), "../qt/CuteGo")

	cmd := exec.Command(
		path.Join(Globals.GoRootPath, "bin/go.exe"),
		"build",
		"-x",
		Globals.MainGoName,
	)

	gopath, _ := os.Getwd()

	cmd.Env = append(
		os.Environ(),
		"GOPATH="+gopath,
		"CGO_LDFLAGS=-L"+path.Join(GetAbsPath(gopath, Globals.TmpDirPath), "release")+" -lCuteGo",
		"CGO_CFLAGS=-I"+cutegoQtSrcRoot,
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
