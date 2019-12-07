package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

//TODO: add flag to build project with hiding of console
type globalsT struct {
	//root path of qml resources
	QrcRoot string `json:"qrcRoot"`
	//root path of qt (aka {QT_INSTALL_PATH}/Qt{QT_VERSION}/{QT_VERSION}/${COMPILER_NAME})
	QtPath string `json:"qtPath"`
	//QT Tools root path of qt (aka {QT_INSTALL_PATH}/Qt{QT_VERSION}/Tools/${COMPILER_NAME})
	QtToolsPath string `json:"qtToolsPath"`
	//Golang compiller root path
	GoInstallPath string `json:"goInstallPath"`
	//final build destination path
	BuildDest string `json:"buildDest"`
	//name of Go main package
	MainGoName string `json:"mainGoName"`
	//path to directory for qt autogenerated files
	TmpDirPath string `json:"tmpDirPath"`
}

//Globals is a main build configuration
var Globals globalsT

//TODO: Add to future environment autodetection
var EOL string = "\r\n"

func generateConfig() {
	config := ""

	//TODO: Implement environment autodetection
	config += "{" + EOL
	config += "  \"//qrcRoot\": \"Path to QML resources root\"," + EOL
	config += "  \"qrcRoot\": \"\"," + EOL
	config += "  \"//qtPath\": \"Path to Qt (e.g. C:/Qt/Qt5.13.1/5.13.1/mingw73_64)\"," + EOL
	config += "  \"qtPath\": \"\"," + EOL
	config += "  \"//qtToolsPath\": \"Path to Qt Tools (e.g. C:/Qt/Qt5.13.1/Tools/mingw730_64)\"," + EOL
	config += "  \"qtToolsPath\": \"\"," + EOL
	config += "  \"//goInstallPath\": \"Path to Go (e.g. C:/Go)\"," + EOL
	config += "  \"goInstallPath\": \"\"," + EOL
	config += "  \"//buildDest\": \"Path to building output\"," + EOL
	config += "  \"buildDest\": \"\"," + EOL
	config += "  \"//mainGoName\": \"Name of main package entry\"," + EOL
	config += "  \"mainGoName\": \"\"," + EOL
	config += "  \"//tmpDirPath\": \"Path to intermediate Qt directory\"," + EOL
	config += "  \"tmpDirPath\": \"\"" + EOL
	config += "}" + EOL

	ioutil.WriteFile("cgbuilder_config.json", []byte(config), 0644)

}

func readGlobals() globalsT {
	var g globalsT

	configFile, err := os.Open("./cgbuilder_config.json")

	if err != nil {
		generateConfig()
		fmt.Println("cgbuilder_config.json is generated, please fill it to build your project")
		os.Exit(0)
	}
	defer configFile.Close()

	bytes, _ := ioutil.ReadAll(configFile)

	json.Unmarshal([]byte(bytes), &g)

	return g
}

func getFlags() map[string]bool {

	args := map[string]bool{
		"res":      false,
		"clean":    false,
		"cleanall": false,
	}

	for i := 1; i < len(os.Args); i++ {
		args[os.Args[i]] = true
	}

	return args

}

func main() {
	//TODO: Write more readable code for os.Args processing
	flags := getFlags()
	Globals = readGlobals()

	pwd, _ := os.Getwd()

	if !flags["clean"] && !flags["cleanall"] {
		if !flags["res"] {
			CreateDir(Globals.TmpDirPath)

			ioutil.WriteFile(path.Join(GetAbsPath(pwd, Globals.TmpDirPath), "CuteGo.pro"), []byte(generateProfile()), 0644)

			buildQt(path.Join(GetAbsPath(pwd, Globals.TmpDirPath), "CuteGo.pro"))

			buildGo()
			deploy()
		}

		rmIfPresent(path.Join(GetAbsPath(pwd, Globals.QrcRoot), "resources.qrc"))
		ioutil.WriteFile(path.Join(GetAbsPath(pwd, Globals.QrcRoot), "resources.qrc"), []byte(generateQrc()), 0644)

		buildResources()

	} else {
		cleanWorkspace()
		if flags["cleanall"] {
			cleanBuildDest()
		}
	}

}
