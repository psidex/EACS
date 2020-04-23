package main

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/actions"
	"github.com/psidex/EACS/internal/config"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/util"
	"strings"
)

func onReady() {
	systray.SetTitle("EACS")
	systray.SetTooltip("Equalizer APO Config Switcher")

	anyConfigsLoaded := false

	configController := config.NewController()
	err := configController.LoadUserConfigs()
	if err != nil {
		util.FatalError(err.Error())
	}

	// This loop sets up the buttons for the user configs.
	for fileName, configStruct := range configController.Configs {
		configName := strings.Replace(fileName, ".txt", "", 1)
		btn := systray.AddMenuItem(configName, "Activate / Deactivate this config")

		if configStruct.Active() {
			btn.Check()
			anyConfigsLoaded = true
		}

		go func(fileName string) {
			for {
				<-btn.ClickedCh
				actions.ButtonClicked(btn, fileName, configController)
			}
		}(fileName)
	}

	systray.AddSeparator()
	QuitBtn := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			<-QuitBtn.ClickedCh
			systray.Quit()
		}
	}()

	if anyConfigsLoaded {
		systray.SetIcon(icon.DataActive)
	} else {
		systray.SetIcon(icon.DataInactive)
	}
}

func onExit() {
	// No cleanup needed.
}

func main() {
	systray.Run(onReady, onExit)
}
