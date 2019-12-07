package main

import (
	"os"
	"path/filepath"
	"runtime"
)

var cutegoQtSrcRoot string
var profile string

const tab = "    "

func addSources(path string, f os.FileInfo, err error) error {
	if isRoot {
		isRoot = false
		return nil
	}

	switch mode := f.Mode(); {
	case mode.IsRegular():
		if filepath.Ext(path) == ".cpp" {
			profile += tab + path + " \\" + EOL
		}
	}
	return nil
}

func addHeaders(path string, f os.FileInfo, err error) error {
	if isRoot {
		isRoot = false
		return nil
	}

	switch mode := f.Mode(); {
	case mode.IsRegular():
		if filepath.Ext(path) == ".h" {
			profile += tab + path + " \\" + EOL
		}
	}
	return nil
}

func generateProfile() string {
	_, currentFilePath, _, _ := runtime.Caller(0)

	cutegoQtSrcRoot = filepath.Join(filepath.Dir(currentFilePath), "../qt/CuteGo")
	profile = ""

	profile += "QT+=quick gui" + EOL
	profile += "TEMPLATE = lib" + EOL
	profile += "DEFINES += CUTEGO_LIBRARY" + EOL
	profile += "CONFIG += c++11" + EOL
	profile += "DEFINES += QT_DEPRECATED_WARNINGS	" + EOL

	profile += "SOURCES += \\" + EOL
	filepath.Walk(cutegoQtSrcRoot, addSources)

	profile = profile[:len(profile)-4]
	profile += EOL

	profile += "HEADERS += \\" + EOL
	filepath.Walk(cutegoQtSrcRoot, addHeaders)

	profile = profile[:len(profile)-4]
	profile += EOL

	return profile

}
