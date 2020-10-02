package buttons

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/appdata"
	"github.com/psidex/EACS/internal/icon"
	"github.com/psidex/EACS/internal/userconfig"
	"github.com/psidex/EACS/internal/util"
)

// ConfigButtonsReceiverLoop starts an infinite loop that receives from the 2 channels and handles the event of that
// config being pressed.
func ConfigButtonsReceiverLoop(fileNameChan chan string, menuItemChan chan *systray.MenuItem, dc *appdata.DataController, allButtons []*systray.MenuItem) {
	for {
		clickedFileName := <-fileNameChan
		clickedBtn := <-menuItemChan

		if clickedBtn.Checked() {
			clickedBtn.Uncheck()
			dc.RemoveActiveConfigFileName(clickedFileName)
		} else {
			if !dc.AllowMultiple() {
				// If we aren't allowing multiple selections, unselect everything first.
				for _, btn := range allButtons {
					btn.Uncheck()
				}
				dc.RemoveAllActiveConfigFileNames()
			}

			clickedBtn.Check()
			dc.AddActiveConfigFileName(clickedFileName)
		}

		activeConfigFileNames := dc.ActiveConfigFileNames()
		err := userconfig.WriteIncludesToMainConfig(activeConfigFileNames)
		if err != nil {
			util.FatalError(err)
		}

		if len(activeConfigFileNames) > 0 {
			systray.SetIcon(icon.DataActive)
		} else {
			systray.SetIcon(icon.DataInactive)
		}

		// This save should happen at the end of the loop after everything else.
		// If this file write happens before the call to WriteIncludesToMainConfig, Equalizer APO's file watcher
		// sometimes doesn't pick up on the write to the config.txt file. I assume this is because 2 different file
		// writes happen in such quick succession that it only registers the 1st and misses the second.
		err = dc.Save()
		if err != nil {
			util.FatalError(err)
		}
	}
}
