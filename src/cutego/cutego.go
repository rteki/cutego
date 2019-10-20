package cutego

// #cgo CFLAGS: -I../../qt/CuteGo/
// #cgo LDFLAGS: -L../../qt/build/release -lCuteGo
// #include "public.h"
import "C"

func Init() bool {
	C.init()
	return true
}
