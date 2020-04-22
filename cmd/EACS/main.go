package main

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/config"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/util"
	"strings"
	"sync"
)

func onReady() {
	systray.SetIcon(icon.DataInactive)
	systray.SetTitle("EACS")
	systray.SetTooltip("Equalizer APO Config Switcher")

	userConfigFileDir := ".\\config-files"
	masterConfigFilePath := "..\\config.txt"

	configWriteMutex := sync.Mutex{}
	userConfigs := config.GetUserConfigs(userConfigFileDir)
	currentConfigFileNames := config.ReadEAPOConfigFromFile(masterConfigFilePath)
	startedWithActiveConfigs := false

	for _, configStruct := range userConfigs {
		configName := strings.Replace(configStruct.FileName, ".txt", "", 1)
		btn := systray.AddMenuItem(configName, "Activate / Deactivate this config")
		configStruct.MenuItem = btn

		// If this config is already in the config master file
		if util.Find(currentConfigFileNames, configStruct.FileName) {
			btn.Check()
			startedWithActiveConfigs = true
		}

		go func() {
			for {
				<-btn.ClickedCh
				if !btn.Checked() {
					btn.Check()
				} else {
					btn.Uncheck()
				}
				configWriteMutex.Lock()
				config.WriteEAPOConfigToFile(masterConfigFilePath, userConfigs)
				configWriteMutex.Unlock()
			}
		}()
	}

	if startedWithActiveConfigs == true {
		systray.SetIcon(icon.DataActive)
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
