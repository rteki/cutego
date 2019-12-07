package main

import (
	"os"
	"path/filepath"
	"strings"
)

var isRoot bool = true
var closeRes bool = false
var qrc string = ""

func tagOpenQres(prefix string) string {
	return "    <qresource prefix=\"" + prefix + "\">" + EOL
}

func tagCloseQres() string {
	return "    </qresource>" + EOL
}

func tagFile(name string) string {
	return "        <file>" + strings.Replace(name, "\\", "/", -1) + "</file>" + EOL
}

func walkFn(path string, f os.FileInfo, err error) error {
	if isRoot {
		isRoot = false
		return nil
	}

	switch mode := f.Mode(); {
	case mode.IsRegular():
		qrc += tagFile(path)
	}
	return nil
}

func generateQrc() string {

	qrc += "<RCC>" + EOL

	qrc += tagOpenQres("/")

	filepath.Walk(".", walkFn)

	qrc += tagCloseQres()

	qrc += "</RCC>" + EOL

	return qrc

}
