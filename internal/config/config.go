package config

import (
	"fmt"
	"github.com/psidex/EACS/internal/util"
	"io/ioutil"
)

const (
	userConfigFileDir    = ".\\config-files"                 // The directory containing config files written by the user.
	masterConfigFilePath = "..\\config.txt"                  // The "master" config file that is read by Equalizer APO.
	includeText          = "Include: EACS\\config-files\\%s" // The default text that includes a file for Equalizer APO.
)

// EApoConfig represents a config file created by the user.
type EApoConfig struct {
	includeText string // The text to be written to the master config.txt file.
	active      bool   // If the config is currently active or not.
}

// Controller provides an abstraction layer for interacting the config files.
// Any actual interaction with the fs must happen within this struct.
type Controller struct {
	configs map[string]*EApoConfig // [filename]struct.
}

// Active is a getter for the EApoConfigs `active` field.
func (c *EApoConfig) Active() bool {
	return c.active
}

// toggleActive toggles the active field.
func (c *EApoConfig) toggleActive() {
	c.active = !c.active
}

// Configs is a getter for the Controllers `configs` field.
func (c *Controller) Configs() map[string]*EApoConfig {
	return c.configs
}

// ToggleConfigActive toggles the `active` field for a given config.
func (c *Controller) ToggleConfigActive(fileName string) {
	// It is impossible to pass a fileName that isn't in the map so no need for err.
	c.configs[fileName].toggleActive()
}

// LoadUserConfigs populates the c.configs map with EApoConfig structs.
func (c *Controller) LoadUserConfigs() error {
	masterConfigLines, err := readLinesFromMasterConfig()
	if err != nil {
		return err
	}

	userConfigFiles, err := ioutil.ReadDir(userConfigFileDir)
	if err != nil {
		return err
	}

	var activeConfigFileNames []string

	// Parse all configs being used in master file.
	for _, line := range masterConfigLines {
		if fileName, err := ParseIncludeText(line); err == nil {
			// No error parsing a filename so append it.
			activeConfigFileNames = append(activeConfigFileNames, fileName)
		}
	}

	// Populate c.configs and set to active if found in master file.
	for _, file := range userConfigFiles {
		if file.IsDir() {
			continue
		}
		c.configs[file.Name()] = &EApoConfig{
			includeText: fmt.Sprintf(includeText, file.Name()),
			active:      util.Find(activeConfigFileNames, file.Name()),
		}
	}

	return nil
}

// WriteActiveConfigs writes the currently active user configs to Equalizer APOs config.txt file.
// Returns variable that shows if any configs are active or not.
func (c *Controller) WriteActiveConfigs() (noActiveConfigs bool, err error) {
	var includeTexts []string
	noActiveConfigs = true

	for _, configStruct := range c.configs {
		if !configStruct.active {
			continue
		}
		includeTexts = append(includeTexts, configStruct.includeText)
		noActiveConfigs = false
	}

	return noActiveConfigs, writeLinesToMasterConfig(includeTexts)
}

// NewController creates a new Controller and initializes the config map.
func NewController() *Controller {
	return &Controller{
		make(map[string]*EApoConfig),
	}
}
