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
	return cli.Exec("xcopy",
		filepath.Join(winPackagePath(), `native\*.dll`),
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
	name := cfg.AppXName()
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

func generateSelfSignedCertificate() error {
	// err := cli.Exec("powershell",
	// 	"New-SelfSignedCertificate",
	// 	"-Type",
	// 	"Custom",
	// 	"-Subject",
	// 	cfg.PublisherID,
	// 	"-KeyUsage",
	// 	"DigitalSignature",
	// 	"-FriendlyName",
	// 	cfg.Publisher,
	// 	"-CertStoreLocation",
	// 	`Cert:\CurrentUser\My`,
	// )
	// if err != nil {
	// 	return err
	// }

	return cli.Exec("powershell",
		"$pwd = ConvertTo-SecureString -String murlok42 -Force -AsPlainText;",
		"Export-PfxCertificate",
		"-cert",
		`Cert:\CurrentUser\My\`+"F8E241C0AFDADC3F82AEE6B0FB4551E58D231E06",
		"-FilePath",
		cfg.ID+".pfx",
		"-Password",
		"$pwd",
	)
}

func sign() error {
	return cli.Exec(`C:\Program Files (x86)\Windows Kits\10\bin\x64\SignTool.exe`,
		"sign",
		"/fd",
		"SHA256",
		"/a",
		"/f",
		cfg.ID+".pfx",
		"/p",
		"murlok42",
		cfg.AppXName(),
	)
}

func install() error {
	return cli.Exec(`C:\Program Files (x86)\Windows Kits\10\bin\x86\WinAppDeployCmd.exe`,
		"install",
		"-file",
		cfg.AppXName(),
		"-ip",
		"127.0.0.1",
	)
}
