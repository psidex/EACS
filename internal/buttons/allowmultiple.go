package buttons

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/appdata"
	"github.com/psidex/EACS/internal/util"
)

// AllowMultipleHandler takes the "Allow Multiple" button and handles it's click event.
func AllowMultipleHandler(allowMultipleBtn *systray.MenuItem, dc *appdata.DataController) {
	for {
		<-allowMultipleBtn.ClickedCh

		if allowMultipleBtn.Checked() {
			allowMultipleBtn.Uncheck()
			dc.SetAllowMultiple(false)
		} else {
			allowMultipleBtn.Check()
			dc.SetAllowMultiple(true)
		}

		err := dc.Save()
		if err != nil {
			util.FatalError(err)
		}
	}
}
