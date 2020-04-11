package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"io/ioutil"
	"os"
	"strings"
	"tawesoft.co.uk/go/dialog"
)

const configFileDir = "./config-files/"
const configFileMaster = "../config.txt"
const includeText = "Include: EACS\\config-files\\%s"

type EApoConfig struct {
	FileName string
	MenuItem *systray.MenuItem
}

func fatalError(errMsg string) {
	dialog.Alert(fmt.Sprintf("EACS Fatal Error\n%s", errMsg))
	os.Exit(1)
}

func CreateConfigSlice() []*EApoConfig {
	var configSlice []*EApoConfig

	files, err := ioutil.ReadDir(configFileDir)
	if err != nil {
		fatalError("Cannot read EACS config file directory")
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
		fatalError("Cannot write to master config file")
	}
}

func ReadConfigFromMaster() []string {
	var currentConfigFileNames []string

	includes, err := ioutil.ReadFile(configFileMaster)
	if err != nil {
		fatalError("Cannot read from master config file")
	}

	lines := strings.Split(string(includes), "\n")

	for _, line := range lines {

		parts := strings.Split(line, "\\")
		fileName := parts[len(parts)-1]

		currentConfigFileNames = append(currentConfigFileNames, fileName)

	}

	return currentConfigFileNames
}
