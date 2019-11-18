package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"tmpwriter"
)

type globalsT struct {
	//root path of qml resources
	QrcRoot string `json:"qrcRoot"`
	//root path of qt (aka {QT_INSTALL_PATH}/Qt{QT_VERSION}/{QT_VERSION}/${COMPILER_NAME})
	QtPath string `json:"qtPath"`
	//intermediate temp path for build
	TmpPathOpt string `json:"tmpPathOpt"`
	//final build destination path
	BuildDest string `json:"buildDest"`
	//name of Go main package
	MainGoName string `json:"mainGoName"`
}

//Globals is a main build configuration
var Globals globalsT
var TW *tmpwriter.TmpWriter
var EOL string = "\r\n"

func readGlobals() globalsT {
	var g globalsT

	configFile, err := os.Open("./cgbuilder_config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)

	json.Unmarshal([]byte(bytes), &g)

	return g
}

func wfu() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	reader.ReadString('\n')
}

func main() {
	Globals = readGlobals()
	TW = new(tmpwriter.TmpWriter)
	defer TW.Close()

	TW.ExecInDir(Globals.QrcRoot, "resources.qrc", generateQrc)
	TW.ExecInDir(".", "CuteGo.pro", generateProfile)

}
