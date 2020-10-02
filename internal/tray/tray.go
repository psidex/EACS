package tray

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/appdata"
	"github.com/psidex/EACS/internal/buttons"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/userconfig"
	"github.com/psidex/EACS/internal/util"
	"sort"
	"strings"
)

const saveDataPath = ".\\data.gob"

// OnReady is the function to be called when the application is ready to run.
func OnReady() {
	// systray has an issue on Windows that causes errors to happen if you call methods before setting the icon.
	// To fix this the first thing that happens in OnReady is settings the icon.
	// https://github.com/getlantern/systray/issues/158
	systray.SetIcon(icon.DataInactive)
	systray.SetTooltip("Equalizer APO Config Switcher")

	dc, err := appdata.NewDataController(saveDataPath)
	if err != nil {
		util.FatalError(err)
	}

	allConfigFileNames, err := userconfig.GetAllUserConfigFileNames()
	if err != nil {
		util.FatalError(err)
	}
	sort.Strings(allConfigFileNames)

	initialActiveConfigFileNames := dc.ActiveConfigFileNames()

	var allButtons []*systray.MenuItem

	// Channels for passing data to the ConfigButtonsReceiverLoop.
	fileNameChan := make(chan string)
	menuItemChan := make(chan *systray.MenuItem)

	for _, fileName := range allConfigFileNames {
		configName := strings.Replace(fileName, ".txt", "", 1)
		btn := systray.AddMenuItem(configName, "Activate / Deactivate this config")
		allButtons = append(allButtons, btn)

		if util.Find(initialActiveConfigFileNames, fileName) {
			btn.Check()
		}

		go func(fileName string) {
			for {
				<-btn.ClickedCh
				fileNameChan <- fileName
				menuItemChan <- btn
			}
		}(fileName)
	}

	// Add the last menu bits.
	systray.AddSeparator()
	allowMultipleBtn := systray.AddMenuItem("Allow Multiple", "")
	quitBtn := systray.AddMenuItem("Quit", "Quit the whole app")

	if dc.AllowMultiple() {
		allowMultipleBtn.Check()
	}

	// Having a single handler for button presses means we don't have to worry about concurrent file access.
	go buttons.ConfigButtonsReceiverLoop(fileNameChan, menuItemChan, dc, allButtons)

	go buttons.AllowMultipleHandler(allowMultipleBtn, dc)
	go buttons.QuitButtonHandler(quitBtn)

	// If we have active configs, set to the active icon.
	if len(initialActiveConfigFileNames) > 0 {
		systray.SetIcon(icon.DataActive)
	}
}
