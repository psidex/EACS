package main

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/config"
	"github.com/psidex/EACS/icon"
	"strings"
)

// Find takes a slice and looks for an element in it. If found it will return true, else false.
// https://golangcode.com/check-if-element-exists-in-slice/
func find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Equalizer APO Config Switcher")
	systray.SetTooltip("Equalizer APO Config Switcher")

	configSlice := config.CreateConfigSlice()
	currentConfigFileNames := config.ReadConfigFromMaster()

	for _, configStruct := range configSlice {
		configName := strings.Replace(configStruct.FileName, ".txt", "", 1)
		btn := systray.AddMenuItem(configName, "Activate / Deactivate this config")
		configStruct.MenuItem = btn

		// If this config is already in the config master file
		if find(currentConfigFileNames, configStruct.FileName) {
			btn.Check()
		}

		go func() {
			for {
				<-btn.ClickedCh
				if !btn.Checked() {
					btn.Check()
				} else {
					btn.Uncheck()
				}
				config.WriteConfigToMaster(configSlice)
			}
		}()
	}

	systray.AddSeparator()

	QuitBtn := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			<-QuitBtn.ClickedCh
			systray.Quit()
		}
	}()

}

func onExit() {
	// No cleanup needed
}

func main() {
	systray.Run(onReady, onExit)
}
