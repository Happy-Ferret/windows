package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/murlokswarm/log"
)

type configuration struct {
	ID          string `json:"id"           help:"Identity name"`
	Name        string `json:"name"         help:"Application name"`
	DisplayName string `json:"display-name" help:"Application name displayed"`
	Publisher   string `json:"publisher"    help:"Publisher name."`
	PublisherID string `json:"publisher-id" help:"Publisher id."`
	Version     string `json:"version"      help:"Version of the app."`
	Icon        string `json:"icon"         help:"The app icon as .png file. Provide a big one!"`
}

func defaultConfig() configuration {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}
	name := filepath.Base(wd)
	user := os.Getenv("USERNAME")
	id := fmt.Sprintf("%s.%s", user, name)
	publisherID := fmt.Sprintf("CN=%s", user)

	return configuration{
		ID:          id,
		Name:        name,
		DisplayName: name,
		Publisher:   name,
		PublisherID: publisherID,
		Version:     "1.0.0.0",
		Icon:        "icon.png",
	}
}

func (c configuration) ExecName() string {
	return c.Name + ".exe"
}

func (c configuration) AppXName() string {
	return c.Name + ".appx"
}

func winPackagePath() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "murlokswarm", "windows")
}

func commandString() string {
	b := bytes.Buffer{}
	fmt.Fprintf(&b, "build\t Builds the AppX.")
	return b.String()
}
