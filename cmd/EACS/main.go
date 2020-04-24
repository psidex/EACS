package main

import (
	"github.com/getlantern/systray"
	"github.com/psidex/EACS/internal/tray"
)

func main() {
	systray.Run(tray.OnReady, tray.OnExit)
}
