package main

import (
	"fmt"
	"strings"
)

var RegFileContent string = `Windows Registry Editor Version 5.00

; Infamous capabilities:

[HKEY_LOCAL_MACHINE\SOFTWARE\FakeBrowser\Capabilities]
"ApplicationDescription"="FakeBrowser"
"ApplicationIcon"="###PATH###FakeBrowser.exe,0"
"ApplicationName"="FakeBrowser"

[HKEY_LOCAL_MACHINE\SOFTWARE\FakeBrowser\Capabilities\FileAssociations]
".htm"="FakeBrowserURL"
".html"="FakeBrowserURL"
".shtml"="FakeBrowserURL"
".xht"="FakeBrowserURL"
".xhtml"="FakeBrowserURL"

[HKEY_LOCAL_MACHINE\SOFTWARE\FakeBrowser\Capabilities\URLAssociations]
"ftp"="FakeBrowserURL"
"http"="FakeBrowserURL"
"https"="FakeBrowserURL"

; Register to Default Programs

[HKEY_LOCAL_MACHINE\SOFTWARE\RegisteredApplications]
"FakeBrowser"="Software\\FakeBrowser\\Capabilities"

; FakeBrowserURL HANDLER:

[HKEY_LOCAL_MACHINE\Software\Classes\FakeBrowserURL]
@="FakeBrowser Document"
"FriendlyTypeName"="FakeBrowser Document"

[HKEY_LOCAL_MACHINE\Software\Classes\FakeBrowserURL\shell]

[HKEY_LOCAL_MACHINE\Software\Classes\FakeBrowserURL\shell\open]

[HKEY_LOCAL_MACHINE\Software\Classes\FakeBrowserURL\shell\open\command]
@="\"###PATH###FakeBrowser.exe\" \"%1\""`

func OutputRegFile() {

	path := strings.Replace(ExecDir, "/", "\\\\", -1)
	content := strings.Replace(RegFileContent, "###PATH###", path, -1)
	fmt.Println(content)
}
