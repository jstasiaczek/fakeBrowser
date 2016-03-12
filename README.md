# fakeBrowser

This application is part of bigger set with urlConstantOpener.

Basically it is used on virtual machine to simulate browser and save url to file in some predefined location. 
It can also send url to predefined url.

I was tested on Windows 7 machine with shared folder.

## Configuration

```
{
	"DirectoryToSave" 	: "/where/to/save/UrlSync",
	"UseFiles" 			: true,
	"UseHttp" 			: false,
	"HttpUrl"			: "http://localhost:4040",
	"SecretKey"			: "some-secret-key"
}
```

You can use Files or Http request, but not both. SecretKey and HttpUrl is used only for Http request.
* UseFiles **[true|false]** _means save url to file_
* DirectoryToSave **[existing directory path]** _where files should be saved, path should point to directory that host machine can access._
* UseHttp **[true|false]** _use http request to send url_
* HttpUrl **[url of server]** _fakeBrowser will add /openUrl to url_
* SecretKey **[some random key]** _it should be the same for fakeBrowser and urlConstantOpener_

Command line parameters:
* -c config file location
* -gen generate reg file working for windows 7
* plain text url

## How to use

On Windows you need to move file to some directory. Copy **config.json.dist** to **config.json** and fill it with correct values.
Now run **fakeBrowser.exe -gen > fakeBrowser.reg** to generate reg file, and execute this file. Sometimes you need to restart you Windows machine.
Go to Windows setting and select new browser as default. Now every clicked url in email or IM will be opened by fakeBrowser.
