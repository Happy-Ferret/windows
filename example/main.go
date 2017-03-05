package main

import (
	"fmt"

	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/windows"
)

func main() {
	fmt.Println("Windows example")
	app.Run()

	// i := 0
	// for {
	// 	time.Sleep(time.Second)
	// 	log.Info(i)
	// 	if i == 3 {
	// 		return
	// 	}
	// 	i++
	// }
}
