package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/murlokswarm/cli"
)

func build() error {
	if err := copyDLL(); err != nil {
		return err
	}

	if err := copyUWP(); err != nil {
		return err
	}

	if err := copyResources(); err != nil {
		return err
	}

	if err := goBuild(); err != nil {
		return nil
	}

	if err := generateManifest(); err != nil {
		return err
	}

	if err := packAppx(); err != nil {
		return err
	}
	return nil
}

func copyDLL() error {
	return cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `lib\*.dll`),
		`AppX\lib\`,
		"/D",
		"/S",
		"/Y",
	)
}

func copyUWP() error {
	return cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `uwp`),
		`AppX\`,
		"/D",
		"/S",
		"/Y",
	)
}

func copyResources() error {
	if err := os.MkdirAll("resources", os.ModeDir); err != nil {
		return err
	}

	return cli.Exec("xcopy",
		"resources",
		`AppX\Resources\`,
		"/D",
		"/E",
		"/Y",
	)
}

func goBuild() error {
	return cli.Exec("go",
		"build",
		"-o",
		filepath.Join("AppX", cfg.ExecName()),
	)
}

func packAppx() error {
	name := fmt.Sprintf("%s.appx", cfg.Name)
	os.Remove(name)

	return cli.Exec(`C:\Program Files (x86)\Windows Kits\10\bin\x64\MakeAppx.exe`,
		"pack",
		"/d",
		`AppX\`,
		"/p",
		name,
		"/l",
	)
}
