package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/kardianos/osext"
)

var ExecDir string
var ConfigPath string
var App AppConfig
var GenerateReg bool

func init() {
	ExecDir, _ = osext.ExecutableFolder()
	ExecDir += "/"
	ConfigPathTmp := ExecDir + "config.json"
	flag.StringVar(&ConfigPath, "c", ConfigPathTmp, "Config file location, default is exec dir + config.json")
	flag.BoolVar(&GenerateReg, "gen", false, "If true generate new reg file")
	flag.Parse()
	flag.Parse()
	App = NewConfig()
	App.ImportSettings(ConfigPath)
}

type AppConfig struct {
	DirectoryToSave string
	UseFiles        bool
	UseHttp         bool
	SecretKey       string
	HttpUrl         string
}

func (this *AppConfig) ImportSettings(path string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Can't open config file! \n" + err.Error())
	}
	err = json.Unmarshal(raw, this)
	if err != nil {
		panic("Can't read config file! \n" + err.Error())
	}
}

func NewConfig() AppConfig {
	return AppConfig{}
}
