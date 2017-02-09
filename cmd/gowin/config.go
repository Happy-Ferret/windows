package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/murlokswarm/log"
)

type configuration struct {
	Name      string `json:"name"      help:"Package name."`
	Publisher string `json:"publisher" help:"Publisher name."`
	Version   string `json:"version"   help:"Version of the app."`
	Icon      string `json:"icon"      help:"The app icon as .png file. Provide a big one!"`
}

func defaultConfig() configuration {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	name := filepath.Base(wd)
	publisher := fmt.Sprintf("CN=%s", name)

	return configuration{
		Name:      name,
		Publisher: publisher,
		Version:   "1.0.0.1",
		Icon:      "icon.png",
	}
}

func (c configuration) ExecName() string {
	return c.Name + ".exe"
}

func winPackagePath() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "murlokswarm", "windows")
}

func commandString() string {
	b := bytes.Buffer{}
	fmt.Fprintf(&b, "build\t Builds the AppX.")
	return b.String()
}
