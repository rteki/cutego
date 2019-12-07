package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func generateMakefile() {
	cmd := exec.Command(
		path.Join(Globals.QtPath, "bin/qmake.exe"),
		"-makefile",
		"-o",
		"Makefile",
		"-spec",
		"win32-g++",
		"\"CONFIG+=qtquickcompiler\"",
		profilePath,
	)

	out, err := cmd.CombinedOutput()

	fmt.Println(cmd.Args)

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}
}

func buildDll() {
	cmd := exec.Command(
		path.Join(Globals.QtToolsPath, "bin/mingw32-make.exe"),
		"-j8",
	)

	out, err := cmd.CombinedOutput()

	fmt.Println(cmd.Args)

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}
}

func cleanArtifacts() {
	cmd := exec.Command(
		path.Join(Globals.QtToolsPath, "bin/mingw32-make.exe"),
		"clean",
	)

	out, err := cmd.CombinedOutput()

	fmt.Println(cmd.Args)

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}
}

var profilePath string

func buildQt(profile string) {
	profilePath = profile

	ExecInDir(Globals.TmpDirPath, generateMakefile)
	ExecInDir(Globals.TmpDirPath, buildDll)
	ExecInDir(Globals.TmpDirPath, cleanArtifacts)
}

func buildResources() {
	cmd := exec.Command(
		path.Join(Globals.QtPath, "bin/rcc.exe"),
		"-binary",
		path.Join(Globals.QrcRoot, "resources.qrc"),
		"-o",
		path.Join(Globals.BuildDest, "resources.rcc"),
	)

	out, err := cmd.Output()

	fmt.Println(cmd.Args)

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}
}
