package cutego

/*
#include "public.h"
extern void qtCallCgo(const char*, const char*, const char*);
*/
import "C"

import (
	"encoding/json"
	"fmt"
	"os"
	"unsafe"
)

var eventManagers map[string]*EventManager

func Init() bool {
	eventManagers = make(map[string]*EventManager)
	C.init((C.GoCallbackFunc)(unsafe.Pointer(C.qtCallCgo)))
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

//export callFromQt
func callFromQt(cEventManagerName *C.char, cEventName *C.char, cValue *C.char) {
	eventManagerName := C.GoString(cEventManagerName)
	eventName := C.GoString(cEventName)
	value := C.GoString(cValue)

	var parsedValue interface{}

	err := json.Unmarshal([]byte(value), &parsedValue)

	if err != nil {
		panic(err)
	}

	em := eventManagers[eventManagerName]

	if em != nil {
		subs := em.Handlers[eventName]

		if subs != nil && len(subs) > 0 {
			for _, sub := range subs {
				sub(parsedValue)
			}
		} else {
			fmt.Println("CuteGo Go Error: Cant find subscribers for " + eventName)
		}

	} else {
		fmt.Println("CuteGo Go Error: Cant find event manager with name \"" + eventManagerName + "\"")
	}

}

func callQt(eventManagerName string, eventName string, value interface{}) {
	emName := C.CString(eventManagerName)
	eName := C.CString(eventName)
	var strValue string

	if marshalled, err := json.Marshal(value); err == nil {
		strValue = string(marshalled)
		sValue := C.CString(strValue)

		C.callQt(emName, eName, sValue)

	} else {
		fmt.Println("Can't marshall value!")
		fmt.Println(err)
		os.Exit(666)
	}

}
