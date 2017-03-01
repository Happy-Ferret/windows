package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/murlokswarm/cli"
)

func build() error {
	if err := initSolution(); err != nil {
		return err
	}

	if err := copyDLL(); err != nil {
		return err
	}

	if err := copyResources(); err != nil {
		return err
	}

	if err := goBuild(); err != nil {
		return err
	}

	if err := generateDebugManifest(); err != nil {
		return err
	}

	if err := generateReleaseManifest(); err != nil {
		return err
	}

	if err := launchSolution(); err != nil {
		return err
	}
	return nil
}

func initSolution() error {
	_, err := os.Stat(".gowin")
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}

	fmt.Println("\033[00;1mCreating Visual Studio solution...\033[00m")

	if err = cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `native\murlok-uwp`),
		`.gowin\`,
		"/D",
		"/S",
		"/Y",
		"/Q",
	); err != nil {
		return err
	}

	os.RemoveAll(`.gowin\bin`)
	os.RemoveAll(`.gowin\obj`)
	return nil
}

func launchSolution() error {
	return cli.Exec(`C:\Program Files (x86)\Microsoft Visual Studio 14.0\Common7\IDE\devenv.exe`,
		"/runexit",
		`.gowin\murlok.sln`,
		"/nologo",
	)
}

func copyDLL() error {
	if err := cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `lib\Win32\*.dll`),
		`.gowin\bin\x86\Debug\AppX\`,
		"/D",
		"/S",
		"/Y",
	); err != nil {
		return err
	}
	if err := cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `lib\Win32\*.dll`),
		`.gowin\bin\x86\Release\AppX\`,
		"/D",
		"/S",
		"/Y",
	); err != nil {
		return err
	}

	if err := cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `lib\x64\*.dll`),
		`.gowin\bin\x64\Debug\AppX\`,
		"/D",
		"/S",
		"/Y",
	); err != nil {
		return err
	}
	return cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `lib\x64\*.dll`),
		`.gowin\bin\x64\Release\AppX\`,
		"/D",
		"/S",
		"/Y",
	)
}

func copyResources() error {
	if err := os.MkdirAll("resources", os.ModeDir); err != nil {
		return err
	}

	if err := cli.Exec("xcopy",
		"resources",
		`.gowin\bin\x64\Debug\AppX\Resources\`,
		"/D",
		"/E",
		"/Y",
	); err != nil {
		return err
	}

	return cli.Exec("xcopy",
		"resources",
		`.gowin\bin\x64\Release\AppX\Resources\`,
		"/D",
		"/E",
		"/Y",
	)
}

func goBuild() error {
	if err := cli.Exec("go",
		"build",
		"-o",
		filepath.Join(`.gowin\bin\x64\Debug\AppX`, cfg.ExecName()),
	); err != nil {
		return err
	}

	return cli.Exec("go",
		"build",
		"-o",
		filepath.Join(`.gowin\bin\x64\Release\AppX`, cfg.ExecName()),
	)
}
