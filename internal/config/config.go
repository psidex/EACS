package config

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/icon"
	"io/ioutil"
	"os"
	"strings"
	"tawesoft.co.uk/go/dialog"
)

const includeText = "Include: EACS\\config-files\\%s"

// EApoConfig holds the name of the file that contains the config, and the associated menu item in the tray
type EApoConfig struct {
	FileName string
	MenuItem *systray.MenuItem
}

// fatalError shows an alert box to the user and exits the program with code 1
func fatalError(errMsg string) {
	dialog.Alert("EACS Fatal Error\n%s", errMsg)
	os.Exit(1)
}

// GetUserConfigs takes a path to a directory and reads all the config files located there into EApoConfig structs
func GetUserConfigs(userConfigPath string) []*EApoConfig {
	var configSlice []*EApoConfig

	files, err := ioutil.ReadDir(userConfigPath)
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

// WriteEAPOConfigToFile takes a path to the Equalizer APO config file and a slice of EApoConfig structs to write there
// This function will also set the systray icon to the active state if any configs are active
func WriteEAPOConfigToFile(pathToConfigFile string, configSlice []*EApoConfig) {
	var completeData []byte
	newline := []byte("\n")
	anyItemsChecked := false

	for _, config := range configSlice {
		if config.MenuItem.Checked() {
			anyItemsChecked = true
			includeStatement := fmt.Sprintf(includeText, config.FileName)
			completeData = append(completeData, []byte(includeStatement)...)
			completeData = append(completeData, newline...)
		}
	}

	if anyItemsChecked == true {
		systray.SetIcon(icon.DataActive)
	} else {
		systray.SetIcon(icon.DataInactive)
	}

	err := ioutil.WriteFile(pathToConfigFile, completeData, 0644)
	if err != nil {
		fatalError("Cannot write to master config file")
	}
}

// ReadEAPOConfigFromFile takes a path to the Equalizer APO config file and reads the names of the currently included config files
func ReadEAPOConfigFromFile(pathToConfigFile string) []string {
	var currentConfigFileNames []string

	includes, err := ioutil.ReadFile(pathToConfigFile)
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
