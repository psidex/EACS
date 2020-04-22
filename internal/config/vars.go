package config

import "github.com/getlantern/systray"

const userConfigFileDir = ".\\config-files"
const masterConfigFilePath = "..\\config.txt"
const includeText = "Include: EACS\\config-files\\%s"

// EApoConfig holds the name of the file that contains the config, and the associated menu item in the tray
type EApoConfig struct {
	FileName string
	MenuItem *systray.MenuItem
}
