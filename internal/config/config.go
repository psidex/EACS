package config

import (
	"fmt"
	"github.com/psidex/EACS/internal/util"
	"io/ioutil"
	"strings"
	"sync"
)

// GetUserConfigs takes a path to a directory and reads all the config files located there into EApoConfig structs
func GetUserConfigs() (userConfigs []*EApoConfig) {
	var configSlice []*EApoConfig

	files, err := ioutil.ReadDir(userConfigFileDir)
	if err != nil {
		util.FatalError("Cannot read EACS config file directory")
	}

	for _, file := range files {
		if !file.IsDir() {
			configStruct := EApoConfig{FileName: file.Name()}
			configSlice = append(configSlice, &configStruct)
		}
	}

	return configSlice
}

// WriteEAPOConfigToFile takes a mutex and a slice of EApoConfig structs to write to the master config file. Returns
// true if any configs were written, else false
func WriteEAPOConfigToFile(writeMutex *sync.Mutex, configSlice []*EApoConfig) (configsWritten bool) {
	var completeData []byte
	newline := []byte("\n")
	anyActiveConfigs := false

	for _, config := range configSlice {
		if config.MenuItem.Checked() {
			anyActiveConfigs = true
			includeStatement := fmt.Sprintf(includeText, config.FileName)
			completeData = append(completeData, []byte(includeStatement)...)
			completeData = append(completeData, newline...)
		}
	}

	writeMutex.Lock()
	if err := ioutil.WriteFile(masterConfigFilePath, completeData, 0644); err != nil {
		util.FatalError("Cannot write to master config file")
	}
	writeMutex.Unlock()

	return anyActiveConfigs
}

// ReadEAPOConfigFromFile takes a path to the Equalizer APO config file and reads the names of the currently included config files
func ReadEAPOConfigFromFile() (fileNames []string) {
	var currentConfigFileNames []string

	includes, err := ioutil.ReadFile(masterConfigFilePath)
	if err != nil {
		util.FatalError("Cannot read from master config file")
	}

	lines := strings.Split(string(includes), "\n")
	for _, line := range lines {
		parts := strings.Split(line, "\\")
		fileName := parts[len(parts)-1]

		currentConfigFileNames = append(currentConfigFileNames, fileName)
	}

	return currentConfigFileNames
}
