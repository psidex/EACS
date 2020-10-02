package userconfig

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	userConfigFileDir  = ".\\config-files"
	mainConfigFilePath = "..\\config.txt"
	includeText        = "Include: EACS\\config-files\\%s"
)

// WriteIncludesToMainConfig writes the given user config file names as Include statements to the main config.
func WriteIncludesToMainConfig(configFileNames []string) error {
	var includeTexts []string

	for _, fileName := range configFileNames {
		includeTexts = append(includeTexts, fmt.Sprintf(includeText, fileName))
	}

	completeData := []byte(strings.Join(includeTexts, "\n"))

	if err := ioutil.WriteFile(mainConfigFilePath, completeData, 0666); err != nil {
		return err
	}
	return nil
}

func GetAllUserConfigFileNames() ([]string, error) {
	var userConfigFileNames []string

	userConfigFiles, err := ioutil.ReadDir(userConfigFileDir)
	if err != nil {
		return userConfigFileNames, err
	}

	for _, file := range userConfigFiles {
		if file.IsDir() {
			continue
		}
		userConfigFileNames = append(userConfigFileNames, file.Name())
	}

	return userConfigFileNames, nil
}
