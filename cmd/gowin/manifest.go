package main

import (
	"os"
	"path/filepath"
	"text/template"
)

var manifestTmpl = `<?xml version="1.0" encoding="utf-8"?>
<Package xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10"
         xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest" 
         xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10" 
         xmlns:rescap="http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities" 
         xmlns:desktop="http://schemas.microsoft.com/appx/manifest/desktop/windows10" IgnorableNamespaces="uap mp rescap desktop build" 
         xmlns:build="http://schemas.microsoft.com/developer/appx/2015/build">
         
  <Identity
    Name="{{.ID}}"  
    Publisher="{{.PublisherID}}"
    Version="{{.Version}}" 
    ProcessorArchitecture="x64" />

  <Properties>
    <DisplayName>{{.DisplayName}}</DisplayName>
    <PublisherDisplayName>{{.Publisher}}</PublisherDisplayName>
    <Logo>Assets\StoreLogo.png</Logo>
  </Properties>

  <Resources>
    <Resource Language="EN-US" />
  </Resources>

  <Dependencies>
    <TargetDeviceFamily Name="Windows.Desktop" MinVersion="10.0.14393.0" MaxVersionTested="10.0.14393.0" />
    <PackageDependency Name="Microsoft.VCLibs.140.00" MinVersion="14.0.22929.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
    <PackageDependency Name="Microsoft.NET.Native.Framework.1.3" MinVersion="1.3.24201.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
    <PackageDependency Name="Microsoft.NET.Native.Runtime.1.4" MinVersion="1.4.24201.0" Publisher="CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" />
  </Dependencies>

  <Applications>
    <Application Id="App"
      Executable="murlok-uwp.exe"
      EntryPoint="murlok_uwp.App">
      
      <uap:VisualElements
        DisplayName="{{.DisplayName}}"
        Square150x150Logo="Assets\Square150x150Logo.png"
        Square44x44Logo="Assets\Square44x44Logo.png"
        Description="murlok-uwp"
        BackgroundColor="transparent">
        <uap:DefaultTile Wide310x150Logo="Assets\Wide310x150Logo.png"/>
        <uap:SplashScreen Image="Assets\SplashScreen.png" />
      </uap:VisualElements>

      <Extensions>
        <uap:Extension Category="windows.appService">
          <uap:AppService Name="CommunicationService" />
        </uap:Extension>
        <desktop:Extension Category="windows.fullTrustProcess" Executable="{{.Name}}.exe" />
      </Extensions>
      
    </Application>
  </Applications>

  <Capabilities>
    <Capability Name="internetClient" />
    <rescap:Capability Name="runFullTrust"/>
  </Capabilities>
</Package>
`

func generateManifest() error {
	name := filepath.Join("AppX", "AppxManifest.xml")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("manisfest").Parse(manifestTmpl))
	return tmpl.Execute(f, cfg)
}
