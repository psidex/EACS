package buttons

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/appdata"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/userconfig"
	"github.com/psidex/EACS/internal/util"
)

// UncheckAllHandler takes the "Uncheck All" button and handles it's click event.
func UncheckAllHandler(uncheckAllBtn *systray.MenuItem, dc *appdata.DataController, allButtons []*systray.MenuItem) {
	for {
		<-uncheckAllBtn.ClickedCh
		uncheckAllBtn.Uncheck()

		dc.RemoveAllActiveConfigFileNames()
		for _, btn := range allButtons {
			btn.Uncheck()
		}

		err := userconfig.WriteIncludesToMainConfig([]string{})
		if err != nil {
			util.FatalError(err)
		}

		systray.SetIcon(icon.DataInactive)

		err = dc.Save()
		if err != nil {
			util.FatalError(err)
		}
	}
}
