package windows

import (
	"errors"
	"runtime"
	"syscall"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/log"
)

var (
	driver   *Driver
	dll      *syscall.DLL
	dllName  = `murlok.dll`
	launched = false
)

func init() {
	var err error
	
	runtime.LockOSThread()
	driver = NewDriver()
	app.RegisterDriver(driver)

	if dll, err = syscall.LoadDLL(dllName); err != nil {
		log.Error(err)
	}
}

// Driver is the implementation of the Windows driver.
type Driver struct {
	closeChan chan bool
}

// NewDriver creates a new Windows driver.
func NewDriver() *Driver {
	return &Driver{
		closeChan: make(chan bool),
	}
}

func (d *Driver) Run() {
	proc, err := dll.FindProc("Driver_Run")
	if err != nil {
		log.Warn(err)
		return
	}
	proc.Call()

	for closed := range d.closeChan {
		if closed {
			return
		}
	}
}

func (d *Driver) NewContext(ctx interface{}) app.Contexter {
	return nil
}

func (d *Driver) MenuBar() app.Contexter {
	return nil
}

func (d *Driver) Dock() app.Docker {
	return nil
}

func (d *Driver) Storage() app.Storer {
	return nil
}

// JavascriptBridge returns the javascript statement to allow javascript to
// call go component methods.
func (d *Driver) JavascriptBridge() string {
	return "alert(msg);"
}

func (d *Driver) Share() app.Sharer {
	return nil
}

func ensureLaunched() {
	if !launched {
		log.Panic(errors.New(`creating and interacting with contexts requires the app to be launched. set app.OnLaunch handler and launch the app by calling app.Run()`))
	}
}
