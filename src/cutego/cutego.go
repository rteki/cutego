package cutego

// #cgo CFLAGS: -I../../qt/CuteGo/
// #cgo LDFLAGS: -L../../release -lCuteGo
// #include "public.h"
import "C"

func Init() bool {
	C.init()
	return true
}

func Start() bool {
	C.start()
	return true
}
