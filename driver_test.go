package windows

import (
	"testing"
	"time"
)

func TestDriverRun(t *testing.T) {
	go func() {
		time.Sleep(time.Second * 1)
		driver.closeChan <- true
	}()
	driver.Run()
}
