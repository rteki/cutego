package cutego




/*
#include "public.h"
extern void qtCallCgo(const char*, const char*, const char*);
*/
import "C"

import (
	"fmt"
	"unsafe"
	"encoding/json"
)

var eventManagers map[string]*EventManager

//export qtCall
func qtCall(cEventManagerName *C.char, cEventName *C.char, cValue *C.char) {
	eventManagerName := C.GoString(cEventManagerName)
	eventName := C.GoString(cEventName)
	value := C.GoString(cValue)	

	mapValue := make(map[string]string)

	err := json.Unmarshal([]byte(value), &mapValue)

	if err != nil {
		panic(err)
	}

	em := eventManagers[eventManagerName]

	if em != nil {
		handler := em.Handlers[eventName]

		if handler != nil {
			handler(mapValue)
		} else {
			fmt.Println("Cant find handler")
		}

	} else {
		fmt.Println("Cant find event manager")
	}

}



func Init() bool {
	eventManagers = make(map[string]*EventManager)
	C.init( (C.GoCallbackFunc)(unsafe.Pointer(C.qtCallCgo)) )
	return true
}

func Start() bool {
	C.start()
	return true
}

func LoadQmlEntry(qmlPath string) bool {
	cstr := C.CString(qmlPath)
	C.loadQmlEntry(cstr)
	return true
}

func newEventManager(name string) bool {
	cstr := C.CString(name)
	C.newEventManager(cstr)
	return true
}