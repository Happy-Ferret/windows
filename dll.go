package windows

import "C"
import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/murlokswarm/log"
)

var (
	dll    *syscall.DLL
	dllptr []uintptr
)

func init() {
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

	initDllCallback("Init_OnLaunch", onLaunch)
	initDllCallback("Init_OnTerminate", onTerminate)
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

func initDllCallback(name string, fn interface{}) {
	a := syscall.NewCallback(fn)
	fmt.Println(a)
	callDllFunc(name, a)
	dllptr = append(dllptr, a)
}
