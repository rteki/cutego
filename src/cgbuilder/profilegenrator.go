package main

import "path"

func generateProfile() string {
	profile := ""
	profile += "QT+=quick gui" + EOL
	profile += "TEMPLATE = lib" + EOL
	profile += "DEFINES += CUTEGO_LIBRARY" + EOL
	profile += "CONFIG += c++11" + EOL
	profile += "DEFINES += QT_DEPRECATED_WARNINGS	" + EOL
	profile += "SOURCES += \\" + EOL
	profile += "	qt/CuteGo/CuteGo.cpp \\" + EOL
	profile += "	qt/CuteGo/public.cpp" + EOL
	profile += "HEADERS += \\" + EOL
	profile += "	qt/CuteGo/CuteGo.h \\" + EOL
	profile += "	qt/CuteGo/public.h" + EOL
	profile += "unix {" + EOL
	profile += "	target.path = /usr/lib" + EOL
	profile += "}" + EOL
	profile += "!isEmpty(target.path): INSTALLS += target" + EOL

	profile += "RESOURCES += \\" + EOL
	profile += "    " + path.Join(Globals.QrcRoot, "resources.qrc") + EOL

	return profile

}
