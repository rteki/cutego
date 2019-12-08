package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

const supportedCompiler = "mingw"

var EOL = "\r\n"
var suggestedTools = ""

func DetectQtPath() string {
	posibleQtPaths := []string{
		"C:/Qt",
		"D:/Qt",
		"C:/Program Files",
		"C:/Program Files (x86)",
	}

	paths := readPaths(posibleQtPaths, "Qt")

	detected := detectQt(paths)

	if len(detected) > 0 {
		if toolsPath, err := filepath.Abs(filepath.Join(detected, "../..", "Tools")); err == nil {
			if PathExists(toolsPath) {
				suggestedTools = toolsPath
			}
		}

		detected = strings.ReplaceAll(detected, "\\", "/")
		fmt.Println("qtPath detected: " + detected)
	}

	return detected
}

func DetectQtToolsPath() string {

	if len(suggestedTools) > 0 {
		possibleTools := getPathsWhichContains(readDir(suggestedTools), supportedCompiler)

		for _, tool := range possibleTools {
			if PathExists(filepath.Join(tool, "bin/mingw32-make.exe")) {
				tool = strings.ReplaceAll(tool, "\\", "/")
				fmt.Println("qtToolsPath detected: " + tool)
				return tool
			}
		}

	}

	fmt.Println("qtToolsPath is not detected")

	return ""
}

func DetectGoRoot() string {
	posibleGoRoots := []string{
		os.Getenv("GOROOT"),
		"C:/Go",
		"D:/Go",
		"D:/Go",
		"C:/Program Files/Go",
		"C:/Program Files (x86)/Go",
	}

	for _, path := range posibleGoRoots {
		if PathExists(filepath.Join(path, "bin/go.exe")) {
			path = strings.ReplaceAll(path, "\\", "/")
			fmt.Println("goRootPath detected: " + path)
			return path
		}
	}

	fmt.Println("goRootPath is not detected")
	return ""
}

func DetectOS() {
	if runtime.GOOS == "windows" {
		EOL = "\r\n"
	} else {
		EOL = "\n"
	}
}

func readPaths(rootsToSearch []string, matchingStr string) []string {
	var posiblePaths []string

	pathEnv := strings.Split(os.Getenv("PATH"), ";")

	posiblePaths = getPathsWhichContains(pathEnv, matchingStr)

	posiblePaths = removeNotExistingPaths(posiblePaths)

	for _, root := range rootsToSearch {
		var filesPaths []string

		files, err := ioutil.ReadDir(root)

		if err == nil {

			for _, f := range files {
				filesPaths = append(filesPaths, filepath.Join(root, f.Name()))
			}

			filesPaths := getPathsWhichContains(filesPaths, matchingStr)
			posiblePaths = append(posiblePaths, filesPaths...)
		}

	}

	return posiblePaths
}

func removeNotExistingPaths(paths []string) []string {
	var existing []string

	for _, path := range paths {
		if PathExists(path) {
			existing = append(existing, path)
		}
	}

	return existing
}

func getPathsWhichContains(paths []string, substr string) []string {
	var matches []string

	for _, val := range paths {
		if strings.Contains(val, substr) {
			matches = append(matches, val)
		}
	}

	return matches
}

func readDir(root string) []string {
	var filePaths []string
	files, err := ioutil.ReadDir(root)

	if err == nil {

		for _, f := range files {
			filePaths = append(filePaths, filepath.Join(root, f.Name()))
		}

	}
	return filePaths
}

func detectQt(paths []string) string {

	qtPath := ""

	for _, path := range paths {
		if strings.Contains(path, supportedCompiler) {
			if re, err := regexp.Compile(`.+\d\.\d\d.\d.` + supportedCompiler + `.+?(\\|/|$)`); err == nil {
				matched := string(re.Find([]byte(path)))

				if len(matched) > 0 {
					if PathExists(filepath.Join(matched, "bin", "qmake.exe")) &&
						PathExists(filepath.Join(matched, "bin", "rcc.exe")) &&
						PathExists(filepath.Join(matched, "bin", "windeployqt.exe")) {
						return matched
					}
				}
			}

		} else {
			dirs := readDir(path)
			containsCompiller := getPathsWhichContains(dirs, supportedCompiler)

			if len(containsCompiller) > 0 {
				qtPath = detectQt(containsCompiller)
			} else {
				qtPath = detectQt(dirs)
			}

			if len(qtPath) > 0 {
				return qtPath
			}

		}
	}

	return ""
}
