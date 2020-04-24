package actions

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/config"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/util"
)

// ButtonClicked should be called when a menu item is clicked by the user.
// It takes the config controller and the file name of the config the pressed button represents.
// It will toggle the config in the controller, write the active configs, and set the tray icon.
func ButtonClicked(cc *config.Controller, fileNameOfConfig string) {
	cc.ToggleConfigActive(fileNameOfConfig)

	allConfigsDisabled, err := cc.WriteActiveConfigs()
	if err != nil {
		util.FatalError(err.Error())
	}

	if allConfigsDisabled {
		systray.SetIcon(icon.DataInactive)
	} else {
		systray.SetIcon(icon.DataActive)
	}
}
