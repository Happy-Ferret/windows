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

	// if err := convertDLL(); err != nil {
	// 	return err
	// }

	if err := copyDLL(); err != nil {
		return err
	}

	if err := copyResources(); err != nil {
		return err
	}

	if err := goBuild(); err != nil {
		return err
	}

	if err := generateManifest(); err != nil {
		return err
	}

	// if err := launchSolution(); err != nil {
	// 	return err
	// }

	if err := deploy(); err != nil {
		return err
	}

	if err := launch(); err != nil {
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

	os.RemoveAll(`.gowin\obj`)
	return nil
}

func copyDLL() error {
	if err := os.MkdirAll(`.gowin\bin\x64\Debug\AppX\`, os.ModeDir|0755); err != nil {
		return err
	}
	if err := os.MkdirAll(`.gowin\bin\x64\Release\AppX\`, os.ModeDir|0755); err != nil {
		return err
	}

	if err := cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `lib\x64\`),
		`.gowin\bin\x64\Debug\AppX\`,
		"/D",
		"/E",
		"/Y",
	); err != nil {
		return err
	}
	return cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `lib\x64\`),
		`.gowin\bin\x64\Release\AppX\`,
		"/D",
		"/E",
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
		"-ldflags",
		"-s",
		"-v",
	); err != nil {
		return err
	}

	if err := cli.Exec("powershell",
		"copy",
		cfg.ExecName(),
		filepath.Join(`.gowin\bin\x64\Debug\AppX\`, cfg.ExecName()),
	); err != nil {
		return err
	}
	return cli.Exec("powershell",
		"copy",
		cfg.ExecName(),
		filepath.Join(`.gowin\bin\x64\Release\AppX\`, cfg.ExecName()),
	)
}

func launchSolution() error {
	return cli.Exec(`C:\Program Files (x86)\Microsoft Visual Studio 14.0\Common7\IDE\devenv.exe`,
		"/runexit",
		`.gowin\murlok.sln`,
		"/nologo",
	)
}

func deploy() error {
	fmt.Println("\033[00;1mBuilding Visual Studio solution...\033[00m")
	return cli.Exec(`C:\Program Files (x86)\Microsoft Visual Studio 14.0\Common7\IDE\devenv.exe`,
		`.gowin\murlok.sln`,
		"/build",
		"Debug",
	)
}

func launch() error {
	return cli.Exec("powershell",
		"start",
		fmt.Sprintf(`shell:AppsFolder\%v_yrmhdqw7xq858!App`, cfg.ID),
		"/wait",
	)
}

func convertDLL() error {
	return cli.Exec("dlltool",
		"--dllname",
		filepath.Join(winPackagePath(), `lib\x64\murlok.dll`),
		"--def",
		filepath.Join(winPackagePath(), `lib\x64\murlok.def`),
		"--output-lib",
		filepath.Join(winPackagePath(), `lib\x64\libmurlok.a`),
	)
}
