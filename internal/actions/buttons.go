package actions

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/config"
	"github.com/psidex/EACS/internal/icon"
	"sync"
)

func ButtonClicked(btn *systray.MenuItem, writeMutex *sync.Mutex, configs []*config.EApoConfig) {
	if !btn.Checked() {
		btn.Check()
	} else {
		btn.Uncheck()
	}

	activeConfigs := config.WriteEAPOConfigToFile(writeMutex, configs)

	// I'm pretty sure SetIcon is thread safe
	if activeConfigs {
		systray.SetIcon(icon.DataActive)
	} else {
		systray.SetIcon(icon.DataInactive)
	}
}
