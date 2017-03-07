package main

import (
	"fmt"
	"time"

	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/windows"
)

func main() {
	fmt.Println("Windows example")

	app.OnLaunch = func() {
		fmt.Println("The app is launched!")
	}

	app.OnFocus = func() {
		fmt.Println("The app is focus!")
	}

	app.OnBlur = func() {
		fmt.Println("The app is blur!")
	}

	app.OnLaunch = func() {
		fmt.Println("The app is launched!")
	}

	app.OnFinalize = func() {
		fmt.Println("The app is finalizing!")
		time.Sleep(time.Second * 1)
	}

	app.Run()
}
