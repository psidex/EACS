package config

import (
	"fmt"
	"github.com/getlantern/systray"
	"io/ioutil"
	"os"
	"tawesoft.co.uk/go/dialog"
)

const configFileDir = "./config-files/"
const configFileMaster = "../config.txt"
const includeText = "Include: EACS\\config-files\\%s"

type EApoConfig struct {
	FileName string
	MenuItem *systray.MenuItem
}

func CreateConfigSlice() []*EApoConfig {
	var configSlice []*EApoConfig

	files, err := ioutil.ReadDir(configFileDir)
	if err != nil {
		errMsg := fmt.Sprintf("Cannot read config file directory: %s", configFileDir)
		dialog.Alert("EACS Config Error", errMsg)
		os.Exit(1)
	}

	for _, file := range files {
		if !file.IsDir() {
			configStruct := EApoConfig{FileName: file.Name()}
			configSlice = append(configSlice, &configStruct)
		}
	}

	return configSlice
}

func WriteConfigToMaster(configSlice []*EApoConfig) {
	var completeData []byte
	newline := []byte("\n")

	for _, config := range configSlice {
		if config.MenuItem.Checked() {
			includeStatement := fmt.Sprintf(includeText, config.FileName)
			completeData = append(completeData, []byte(includeStatement)...)
			completeData = append(completeData, newline...)
		}
	}

	err := ioutil.WriteFile(configFileMaster, completeData, 0644)
	if err != nil {
		errMsg := fmt.Sprintf("Cannot write to master config file: %s", configFileMaster)
		dialog.Alert("EACS Config Error", errMsg)
		os.Exit(1)
	}
}
