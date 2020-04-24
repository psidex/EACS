package config

import (
	"fmt"
	"github.com/psidex/EACS/internal/util"
	"io/ioutil"
	"strings"
	"sync"
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

// Controller is the main controller for reading / writing configs.
type Controller struct {
	mutex   sync.Mutex
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
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// It is impossible to pass a fileName that isn't in the map so no need for err.
	c.configs[fileName].toggleActive()
}

// LoadUserConfigs populates the mc.configs map with EApoConfig structs.
func (c *Controller) LoadUserConfigs() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	masterConfigLines, err := readLinesFromMasterConfig()
	if err != nil {
		return err
	}

	userConfigFiles, err := ioutil.ReadDir(userConfigFileDir)
	if err != nil {
		return err
	}

	var activeConfigFileNames []string

	for _, line := range masterConfigLines {
		parts := strings.Split(line, "\\")
		fileName := parts[len(parts)-1]
		activeConfigFileNames = append(activeConfigFileNames, fileName)
	}

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
func (c *Controller) WriteActiveConfigs() (allConfigsDisabled bool, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var includeTexts []string
	allDisabled := true

	for _, configStruct := range c.configs {
		if !configStruct.active {
			continue
		}
		includeTexts = append(includeTexts, configStruct.includeText)
		allDisabled = false
	}

	return allDisabled, writeLinesToMasterConfig(includeTexts)
}

// NewController creates a new Controller and initializes the config map.
func NewController() *Controller {
	var m Controller
	m.configs = make(map[string]*EApoConfig)
	return &m
}
