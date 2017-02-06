package windows

import (
	"errors"
	"log"
	"runtime"
	"syscall"

	"github.com/murlokswarm/app"
)

var (
	driver   *Driver
	dll      *syscall.DLL
	launched = false
)

func init() {
	runtime.LockOSThread()
	driver = NewDriver()
	app.RegisterDriver(driver)

	var err error
	if dll, err = syscall.LoadDLL(`native\x64\Release\murlok.dll`); err != nil {
		log.Panic(err)
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
		log.Panic(err)
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
