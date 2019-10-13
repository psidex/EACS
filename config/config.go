package config

import (
	"github.com/getlantern/systray"
	"io/ioutil"
	"os"
	"strings"
	"tawesoft.co.uk/go/dialog"
)

const configFileDir = "./config-files/"
const configFileMaster = "../config.txt"

type EApoConfig struct {
	Name     string
	Data     []byte
	MenuItem *systray.MenuItem
}

// Takes a FileInfo interface. Assuming that file is in configFileDir, load the contents of that file into an
// EApoConfig struct and return it.
func configStructFromFile(file os.FileInfo) (EApoConfig, error) {
	fileName := file.Name()

	dat, err := ioutil.ReadFile(configFileDir + fileName)
	if err != nil {
		return EApoConfig{}, err
	}

	configName := strings.Replace(fileName, ".txt", "", 1)

	return EApoConfig{
		Name: configName,
		Data: dat,
	}, nil
}

// Creates a slice of EApoConfigs which contain the data for all config files in configFileDir.
func CreateConfigSlice() []*EApoConfig {
	var configSlice []*EApoConfig

	files, err := ioutil.ReadDir(configFileDir)
	if err != nil {
		dialog.Alert("E-APO Config Error", "Cannot read config file directory: "+configFileDir)
		// Can't do anything else.
		os.Exit(1)
	}

	for _, file := range files {
		if !file.IsDir() {
			configStruct, err := configStructFromFile(file)
			if err == nil {
				configSlice = append(configSlice, &configStruct)
			}
		}
	}

	return configSlice
}

// Write all items in configSlice to configFileMaster if their associated menu item is checked.
func WriteConfigToMaster(configSlice []*EApoConfig) {
	var completeData []byte
	newline := []byte("\n")

	for _, config := range configSlice {
		if config.MenuItem.Checked() {
			// "..." -> https://stackoverflow.com/a/16248257/6396652
			completeData = append(completeData, newline...)
			completeData = append(completeData, config.Data...)
		}
	}

	err := ioutil.WriteFile(configFileMaster, completeData, 0644)
	if err != nil {
		dialog.Alert("E-APO Config Error", "Cannot write to master config file: "+configFileMaster)
		os.Exit(1)
		return
	}
}
