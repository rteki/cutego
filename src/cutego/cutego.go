package cutego

import "fmt"

// #cgo CFLAGS: -I../../qt/CuteGo/
// #cgo LDFLAGS: -L../../qt/build/release -lCuteGo
// #include "public.h"
import "C"

func Init() bool {
	C.init();
	fmt.Println("hello")
	return true
}


func Start() bool {
	fmt.Println("here")
	C.start();
	return true;
}
