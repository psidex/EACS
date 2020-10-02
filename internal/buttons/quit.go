package buttons

import "github.com/getlantern/systray"

// QuitButtonHandler handles the quit button being clicked.
func QuitButtonHandler(quitBtn *systray.MenuItem) {
	<-quitBtn.ClickedCh
	systray.Quit()
}
