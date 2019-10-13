package main

import (
	"github.com/getlantern/systray"
	"github.com/psidex/E-APO-Config-Switcher/config"
	"github.com/psidex/E-APO-Config-Switcher/icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("E-APO Config Switcher")
	systray.SetTooltip("E-APO Config Switcher")

	configSlice := config.CreateConfigSlice()

	for _, configStruct := range configSlice {
		btn := systray.AddMenuItem(configStruct.Name, "Activate / Deactivate this config")
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
