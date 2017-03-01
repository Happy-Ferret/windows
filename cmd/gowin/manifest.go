package main

import (
	"os"
	"path/filepath"
	"text/template"
)

var (
	manisfestTmpl = `<?xml version="1.0" encoding="utf-8"?>
<Package 
	xmlns="http://schemas.microsoft.com/appx/manifest/foundation/windows10" 
	xmlns:mp="http://schemas.microsoft.com/appx/2014/phone/manifest" 
	xmlns:uap="http://schemas.microsoft.com/appx/manifest/uap/windows10"
	xmlns:rescap="http://schemas.microsoft.com/appx/manifest/foundation/windows10/restrictedcapabilities"
	xmlns:desktop="http://schemas.microsoft.com/appx/manifest/desktop/windows10" IgnorableNamespaces="uap mp rescap desktop">	
  <Identity Name="{{.ID}}" Publisher="{{.PublisherID}}" Version="{{.Version}}" />
  <mp:PhoneIdentity PhoneProductId="5f3c9717-347d-4f5c-8e57-ca5ea643ff85" PhonePublisherId="00000000-0000-0000-0000-000000000000" />
  <Properties>
    <DisplayName>{{.DisplayName}}</DisplayName>
    <PublisherDisplayName>{{.Publisher}}</PublisherDisplayName>
    <Logo>Assets\StoreLogo.png</Logo>
  </Properties>
  <Dependencies>
    <TargetDeviceFamily Name="Windows.Desktop" MinVersion="10.0.14332.0" MaxVersionTested="10.0.14332.0" />
  </Dependencies>
  <Resources>
    <Resource Language="x-generate" />
  </Resources>
  <Applications>
    <Application Id="App" Executable="$targetnametoken$.exe" EntryPoint="murlok_uwp.App">
      <uap:VisualElements DisplayName="{{.DisplayName}}" Square150x150Logo="Assets\Square150x150Logo.png" Square44x44Logo="Assets\Square44x44Logo.png" Description="murlok-uwp" BackgroundColor="transparent">
        <uap:DefaultTile Wide310x150Logo="Assets\Wide310x150Logo.png">
        </uap:DefaultTile>
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
	<rescap:Capability Name="runFullTrust" />
  </Capabilities>
</Package>
  `
)

func generateManifest() error {
	name := filepath.Join(`.gowin`, "Package.appxmanifest")
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	tmpl := template.Must(template.New("manisfest").Parse(manisfestTmpl))
	return tmpl.Execute(f, cfg)
}
