package main

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/config"
	"github.com/psidex/EACS/icon"
	"strings"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Equalizer APO Config Switcher")
	systray.SetTooltip("Equalizer APO Config Switcher")

	configSlice := config.CreateConfigSlice()

	for _, configStruct := range configSlice {
		configName := strings.Replace(configStruct.FileName, ".txt", "", 1)
		btn := systray.AddMenuItem(configName, "Activate / Deactivate this config")
		configStruct.MenuItem = btn

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

	blank := systray.AddMenuItem("", "")
	blank.Disable()

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
