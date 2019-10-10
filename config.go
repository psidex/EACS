package main

import (
	"github.com/getlantern/systray"
	"io/ioutil"
	"os"
)

const configFileDir = "./config-files/"
const configFileMaster = "../config.txt"

type EApoConfig struct {
	Name     string
	Data     []byte
	MenuItem *systray.MenuItem
}

func configFromFile(file os.FileInfo) EApoConfig {
	fileName := file.Name()

	// TODO: Deal with err (currently _)
	dat, _ := ioutil.ReadFile(configFileDir + fileName)

	return EApoConfig{
		Name: fileName,
		Data: dat,
	}
}

func CreateConfigSlice() []*EApoConfig {
	var configSlice []*EApoConfig

	// TODO: Deal with err (currently _)
	files, _ := ioutil.ReadDir(configFileDir)

	for _, file := range files {
		if !file.IsDir() {
			c := configFromFile(file)
			configSlice = append(configSlice, &c)
		}
	}

	return configSlice
}
