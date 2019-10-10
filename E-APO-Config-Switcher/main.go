package main

// go build -o ./E-APO-Config-Switcher.exe github.com/psidex/E-APO-Config-Switcher/E-APO-Config-Switcher

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/psidex/E-APO-Config-Switcher/E-APO-Config-Switcher/icon"
	"io/ioutil"
)

func main() {
	systray.Run(onReady, onExit)
}

func writeConfigToMaster(configSlice []*EApoConfig) {
	// Takes a slice of EApoConfigs and writes each checked one to configFileMaster
	var completeData []byte
	newline := []byte("\n")

	for _, config := range configSlice {
		if config.MenuItem.Checked() {
			// "..." -> https://stackoverflow.com/a/16248257/6396652
			completeData = append(completeData, newline...)
			completeData = append(completeData, config.Data...)
		}
	}

	err := ioutil.WriteFile(configFileMaster, completeData, 0644)
	if err != nil {
		// TODO: Do something more than just this
		fmt.Println(err)
		return
	}
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("E-APO Config Switcher")
	systray.SetTooltip("E-APO Config Switcher")

	configSlice := CreateConfigSlice()

	for _, config := range configSlice {
		btn := systray.AddMenuItem(config.Name, "tooltip")
		config.MenuItem = btn

		// config is passed to the routine because it's one variable whose value is being changed every loop, so by the
		// time the routine is run, the value may be different. Btn is a different variable each loop so does not need
		// to be passed
		go func(config *EApoConfig) {
			for {
				<-btn.ClickedCh
				if !btn.Checked() {
					btn.Check()
				} else {
					btn.Uncheck()
				}
				writeConfigToMaster(configSlice)
			}
		}(config)
	}

	QuitBtn := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		for {
			<-QuitBtn.ClickedCh
			systray.Quit()
		}
	}()
}

func onExit() {
	// clean up here
}
