package windows

/*
#include "bridge.hpp"
*/
import "C"
import (
	"fmt"
	"syscall"

	"github.com/murlokswarm/log"
)

func init() {
	dll, err := syscall.LoadDLL(`native\x64\Release\murlok.dll`)
	if err != nil {
		log.Panic(err)
	}

	proc, err := dll.FindProc("InitOnMurlokPrint")
	if err != nil {
		log.Panic(err)
	}
	proc.Call(uintptr(C.CPP_OnPrint))
}

//export  goCallback
func goCallback(cstr *C.char) {
	fmt.Printf("from c++: %v\n", C.GoString(cstr))
}
