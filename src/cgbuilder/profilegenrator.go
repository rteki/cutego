package main

import (
	"os"
	"path"
)

func generateProfile() string {
	pwd, _ := os.Getwd()

	profile := ""
	profile += "QT+=quick gui" + EOL
	profile += "TEMPLATE = lib" + EOL
	profile += "DEFINES += CUTEGO_LIBRARY" + EOL
	profile += "CONFIG += c++11" + EOL
	profile += "DEFINES += QT_DEPRECATED_WARNINGS	" + EOL
	profile += "SOURCES += \\" + EOL

	profile += "	" + path.Join(pwd, "qt/CuteGo/CuteGo.cpp") + " \\" + EOL
	profile += "	" + path.Join(pwd, "qt/CuteGo/EventEmitter.cpp") + " \\" + EOL
	profile += "	" + path.Join(pwd, "qt/CuteGo/public.cpp") + EOL
	profile += "HEADERS += \\" + EOL
	profile += "	" + path.Join(pwd, "qt/CuteGo/CuteGo.h") + "\\" + EOL
	profile += "	" + path.Join(pwd, "qt/CuteGo/EventEmitter.h") + "\\" + EOL
	profile += "	" + path.Join(pwd, "qt/CuteGo/public.h") + EOL

	return profile

}
