package actions

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/config"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/util"
)

// ButtonClicked should be called when a systray.MenuItem is clicked by the user.
// It takes the menu item that was clicked, the file name of the config it represents, and the config controller.
// It will toggle the button, toggle the config in the controller, write the active configs, and set the tray icon.
func ButtonClicked(btn *systray.MenuItem, fileNameOfConfig string, cc *config.Controller) {
	if !btn.Checked() {
		btn.Check()
	} else {
		btn.Uncheck()
	}
	cc.ToggleConfigActive(fileNameOfConfig)

	allConfigsDisabled, err := cc.WriteActiveConfigs()
	if err != nil {
		util.FatalError(err.Error())
	}

	// I'm pretty sure SetIcon is thread safe
	if allConfigsDisabled {
		systray.SetIcon(icon.DataInactive)
	} else {
		systray.SetIcon(icon.DataActive)
	}
}
