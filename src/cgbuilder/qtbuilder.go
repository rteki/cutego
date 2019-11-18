package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func generateMakefile(profile string) {
	cmd := exec.Command(
		path.Join(Globals.QtPath, "bin/qmake.exe"),
		"-makefile",
		"-o",
		"Makefile",
		"-spec",
		"win32-g++",
		"\"CONFIG+=qtquickcompiler\"",
		profile,
	)

	out, err := cmd.Output()

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

	out, err := cmd.Output()

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

	out, err := cmd.Output()

	fmt.Println(cmd.Args)

	fmt.Println(string(out))

	if err != nil {
		fmt.Println(err)
		os.Exit(666)
	}
}

func buildQt(profile string) {
	generateMakefile(profile)
	buildDll()
	cleanArtifacts()
}
