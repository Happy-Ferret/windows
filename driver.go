package windows

import (
	"errors"
	"log"
	"runtime"

	"github.com/murlokswarm/app"
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
}

// NewDriver creates a new Windows driver.
func NewDriver() *Driver {
	return &Driver{}
}

func (d *Driver) Run() {
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
