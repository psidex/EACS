package tray

import (
	"sort"
	"strings"

	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/actions"
	"github.com/psidex/EACS/internal/config"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/util"
)

// OnReady is the function to be called when the application is ready to run.
func OnReady() {
	// systray has an issue on Windows that causes errors to happen if you call methods before setting the icon.
	// To fix this the first thing that happens in OnReady is settings the icon.
	// https://github.com/getlantern/systray/issues/158
	systray.SetIcon(icon.DataInactive)
	systray.SetTooltip("Equalizer APO Config Switcher")

	// Create the config controller and load the user configs.
	configController := config.NewController()
	if err := configController.LoadUserConfigs(); err != nil {
		util.FatalError(err.Error())
	}
	configs := configController.Configs()

	// Get all user configs and sort them by file name.
	var sortedFileNames []string
	for fileName := range configs {
		sortedFileNames = append(sortedFileNames, fileName)
	}
	sort.Strings(sortedFileNames)

	// Set up the buttons for the user created configs.
	buttonClickedChan := make(chan string)
	anyConfigsLoaded := false

	for _, fileName := range sortedFileNames {
		configName := strings.Replace(fileName, ".txt", "", 1)
		btn := systray.AddMenuItem(configName, "Activate / Deactivate this config")

		configStruct := configs[fileName]
		if configStruct.Active() {
			btn.Check()
			anyConfigsLoaded = true
		}

		go func(fileName string) {
			for {
				<-btn.ClickedCh
				if !btn.Checked() {
					btn.Check()
				} else {
					btn.Uncheck()
				}
				buttonClickedChan <- fileName
			}
		}(fileName)
	}

	// The listener for button presses.
	// access to `configs` in the above loop remains safe (as configController is not thread safe).
	go func() {
		for {
			actions.ButtonClicked(configController, <-buttonClickedChan)
		}
	}()

	// If we loaded configs, set to the active icon.
	if anyConfigsLoaded {
		systray.SetIcon(icon.DataActive)
	}

	// Add the last menu bits.
	systray.AddSeparator()
	QuitBtn := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			<-QuitBtn.ClickedCh
			systray.Quit()
		}
	}()
}
