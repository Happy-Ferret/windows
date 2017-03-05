package windows

import (
	"errors"
	"runtime"

	"github.com/murlokswarm/app"
	"github.com/murlokswarm/log"
)

var (
	driver   *Driver
	launched = false
)

func init() {
	runtime.LockOSThread()
	driver = NewDriver()
	app.RegisterDriver(driver)
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
	go callDllFunc("Driver_Run")
	<-d.closeChan
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

func onLaunch() uintptr {
	launched = true
	log.Info("OMG driver is really launched")

	// app.UIChan <- func() {
	// 	if app.OnLaunch != nil {
	// 		app.OnLaunch()
	// 	}
	// }
	return 0
}

func onTerminate() uintptr {
	termChan := make(chan bool)

	app.UIChan <- func() {
		if app.OnTerminate != nil {
			termChan <- app.OnTerminate()
			return
		}
		termChan <- true
	}

	if <-termChan {
		return 1
	}
	return 0
}
