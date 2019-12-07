package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func moveDll() {
	err := os.Rename(path.Join(Globals.TmpDirPath, "release/CuteGo.dll"), path.Join(Globals.BuildDest, "CuteGo.dll"))
	if err != nil {
		fmt.Println(err)
	}
}

func moveExe() {
	err := os.Rename(Globals.MainGoName+".exe", path.Join(Globals.BuildDest, Globals.MainGoName+".exe"))
	if err != nil {
		fmt.Println(err)
	}
}

var fullResourcesPath string

func winQtDeploy() {
	cmd := exec.Command(
		path.Join(Globals.QtPath, "bin/windeployqt.exe"),
		"--qmldir",
		fullResourcesPath,
		"CuteGo.dll",
	)

	out, err := cmd.Output()

	fmt.Println(cmd.Args)

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}
}

func deploy() {
	CreateDir(Globals.BuildDest)

	moveDll()
	moveExe()

	localFullResourcesPath, err := filepath.Abs(Globals.QrcRoot)
	fullResourcesPath = localFullResourcesPath

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}

	ExecInDir(Globals.BuildDest, winQtDeploy)
}
