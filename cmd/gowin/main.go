package main

import (
	"fmt"

	"github.com/murlokswarm/config"
)

var (
	cfg = defaultConfig()
)

func main() {
	var err error

	config.ConfigName = "windows.json"
	config.CreateIfNotExists = true
	config.Commands = commandString()

	cmds := config.Load(&cfg)
	if len(cmds) != 1 {
		fmt.Printf("\033[91mInvalid command: %v. use gowin -h for help\033[00m\n", cmds)
		return
	}

	switch cmds[0] {
	case "build":
		err = build()
	}

	if err != nil {
		fmt.Println(err)
	}
}
