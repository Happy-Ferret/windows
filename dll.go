package windows

/*
#include "dll.hpp"
*/
import "C"
import (
	"os"
	"path/filepath"
	"syscall"

	"unsafe"

	"github.com/murlokswarm/log"
)

var (
	dll    *syscall.DLL
	dllptr []uintptr
)

func initDll() {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	packageName := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "murlokswarm", "windows")

	dllName := "murlok.dll"
	if wd == packageName {
		dllName = `lib\x64\murlok.dll`
	}

	if dll, err = syscall.LoadDLL(dllName); err != nil {
		log.Panic(err)
	}

	initDllCallback("Init_OnLaunch", C.DLL_OnLaunch)
	initDllCallback("Init_OnFocus", C.DLL_OnFocus)
	initDllCallback("Init_OnBlur", C.DLL_OnBlur)
	initDllCallback("Init_OnTerminate", C.DLL_OnTerminate)
	initDllCallback("Init_OnFinalize", C.DLL_OnFinalize)
}

func releaseDll() {
	dll.Release()
}

func callDllFunc(name string, a ...uintptr) (r uintptr) {
	proc, err := dll.FindProc(name)
	if err != nil {
		log.Error(err)
		return
	}
	r, _, _ = proc.Call(a...)
	return
}

func initDllCallback(name string, fn unsafe.Pointer) {
	callDllFunc(name, uintptr(fn))
}
