package main

import (
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
		filepath.Join(winPackagePath(), `native\murlok-uwp\bin\x64\Release\AppX`),
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
