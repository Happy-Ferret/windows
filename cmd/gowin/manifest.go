package main

import (
	"os"
	"path/filepath"
	"text/template"
)

var (
	releaseManifestTmpl = `<?xml version="1.0" encoding="utf-8"?>
<Package xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10" xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest" xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10" xmlns:rescap="http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities" xmlns:desktop="http://schemas.microsoft.com/appx/manifest/desktop/windows10" IgnorableNamespaces="uap mp rescap desktop build" xmlns:build="http://schemas.microsoft.com/developer/appx/2015/build">
	<Identity Name="{{.ID}}" Publisher="{{.PublisherID}}" Version="{{.Version}}" ProcessorArchitecture="x64" />
	<PhoneIdentity PhoneProductId="{{.ID}}" PhonePublisherId="00000000-0000-0000-0000-000000000000" />
	<Properties>
		<DisplayName>
			{{.DisplayName}}
		</DisplayName>
		<PublisherDisplayName>
			{{.Publisher}}
		</PublisherDisplayName>
		<Logo>
			Assets\StoreLogo.png
		</Logo>
	</Properties>
	<Dependencies>
		<TargetDeviceFamily Name="Windows.Universal" MinVersion="10.0.14393.0" MaxVersionTested="10.0.14393.0" />
		<PackageDependency Name="Microsoft.VCLibs.140.00" MinVersion="14.0.22929.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
		<PackageDependency Name="Microsoft.NET.Native.Framework.1.3" MinVersion="1.3.24201.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
		<PackageDependency Name="Microsoft.NET.Native.Runtime.1.4" MinVersion="1.4.24201.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
	</Dependencies>
	<Resources>
		<Resource Language="EN-US" />
	</Resources>
	<Applications>
		<Application Id="App" Executable="murlok-uwp.exe" EntryPoint="murlok_uwp.App">
			<VisualElements DisplayName="{{.DisplayName}}" Square150x150Logo="Assets\Square150x150Logo.png" Square44x44Logo="Assets\Square44x44Logo.png" Description="murlok-uwp" BackgroundColor="transparent">
				<DefaultTile Wide310x150Logo="Assets\Wide310x150Logo.png" />
				<SplashScreen Image="Assets\SplashScreen.png" />
			</VisualElements>
			<Extensions>
				<Extension Category="windows.appService">
					<AppService Name="CommunicationService" />
				</Extension>
				<Extension Category="windows.fullTrustProcess" Executable="{{.Name}}.exe" />
			</Extensions>
		</Application>
	</Applications>
	<Capabilities>
		<Capability Name="internetClient" />
		<Capability Name="runFullTrust" />
	</Capabilities>
</Package>
  `

	debugManisfestTmpl = `<?xml version="1.0" encoding="utf-8"?>
<Package xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10" xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest" xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10" xmlns:rescap="http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities" xmlns:desktop="http://schemas.microsoft.com/appx/manifest/desktop/windows10" IgnorableNamespaces="uap mp rescap desktop build" xmlns:build="http://schemas.microsoft.com/developer/appx/2015/build">
	<Identity Name="{{.ID}}" Publisher="{{.PublisherID}}" Version="{{.Version}}" ProcessorArchitecture="x64" />
	<PhoneIdentity PhoneProductId="{{.ID}}" PhonePublisherId="00000000-0000-0000-0000-000000000000" />
	<Properties>
		<DisplayName>
			{{.DisplayName}}
		</DisplayName>
		<PublisherDisplayName>
			{{.Publisher}}
		</PublisherDisplayName>
		<Logo>
			Assets\StoreLogo.png
		</Logo>
	</Properties>
	<Dependencies>
		<TargetDeviceFamily Name="Windows.Universal" MinVersion="10.0.14393.0" MaxVersionTested="10.0.14393.0" />
		<PackageDependency Name="Microsoft.NET.CoreRuntime.1.0" MinVersion="1.0.23819.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
		<PackageDependency Name="Microsoft.VCLibs.140.00.Debug" MinVersion="14.0.24210.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
	</Dependencies>
	<Resources>
		<Resource Language="EN-US" />
	</Resources>
	<Applications>
		<Application Id="App" Executable="murlok-uwp.exe" EntryPoint="murlok_uwp.App">
			<VisualElements DisplayName="{{.DisplayName}}" Square150x150Logo="Assets\Square150x150Logo.png" Square44x44Logo="Assets\Square44x44Logo.png" Description="murlok-uwp" BackgroundColor="transparent">
				<DefaultTile Wide310x150Logo="Assets\Wide310x150Logo.png" />
				<SplashScreen Image="Assets\SplashScreen.png" />
			</VisualElements>
			<Extensions>
				<Extension Category="windows.appService">
					<AppService Name="CommunicationService" />
				</Extension>
				<Extension Category="windows.fullTrustProcess" Executable="{{.Name}}.exe" />
			</Extensions>
		</Application>
	</Applications>
	<Capabilities>
		<Capability Name="internetClient" />
		<Capability Name="runFullTrust" />
	</Capabilities>
</Package>
  `
)

func generateDebugManifest() error {
	name := filepath.Join(`.gowin\bin\x64\Debug\AppX`, "AppxManifest.xml")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("manisfest").Parse(debugManisfestTmpl))
	return tmpl.Execute(f, cfg)
}

func generateReleaseManifest() error {
	name := filepath.Join(`.gowin\bin\x64\Release\AppX`, "AppxManifest.xml")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("manisfest").Parse(releaseManifestTmpl))
	return tmpl.Execute(f, cfg)
}
