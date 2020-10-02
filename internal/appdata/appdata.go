package appdata

import (
	"encoding/gob"
	"os"
)

// appData represents the app data to store.
// Fields are public so gob works.
type appData struct {
	AllowMultiple         bool
	ActiveConfigFileNames []string
}

// DataController controls the loading, saving, and storage of app appData.
type DataController struct {
	fileName string
	appData  appData
}

// NewDataController creates a new DataController and loads from the file.
func NewDataController(fileName string) (*DataController, error) {
	d := &DataController{fileName: fileName}
	err := d.load()
	return d, err
}

// AllowMultiple is a getter for appData.AllowMultiple.
func (d DataController) AllowMultiple() bool {
	return d.appData.AllowMultiple
}

// SetAllowMultiple is a setter for appData.AllowMultiple.
func (d *DataController) SetAllowMultiple(allowMultiple bool) {
	d.appData.AllowMultiple = allowMultiple
}

// ActiveConfigFileNames is a getter for appData.ActiveConfigFileNames.
func (d DataController) ActiveConfigFileNames() []string {
	return d.appData.ActiveConfigFileNames
}

// AddActiveConfigFileName appends a config file name to appData.ActiveConfigFileNames.
func (d *DataController) AddActiveConfigFileName(fileName string) {
	d.appData.ActiveConfigFileNames = append(d.appData.ActiveConfigFileNames, fileName)
}

// RemoveActiveConfigFileName removes a config file name from appData,ActiveConfigFileNames,
func (d *DataController) RemoveActiveConfigFileName(fileName string) {
	// https://stackoverflow.com/a/34070691/6396652
	for i, v := range d.appData.ActiveConfigFileNames {
		if v == fileName {
			d.appData.ActiveConfigFileNames = append(d.appData.ActiveConfigFileNames[:i], d.appData.ActiveConfigFileNames[i+1:]...)
		}
	}
}

// RemoveAllActiveConfigFileNames removes all active config file names.
func (d *DataController) RemoveAllActiveConfigFileNames() {
	d.appData.ActiveConfigFileNames = []string{}
}

// Save writes the current appData to the set file.
func (d DataController) Save() error {
	dataFile, err := os.Create(d.fileName)
	if err != nil {
		return err
	}

	dataEncoder := gob.NewEncoder(dataFile)

	err = dataEncoder.Encode(d.appData)
	if err != nil {
		return err
	}

	err = dataFile.Close()
	if err != nil {
		return err
	}

	return nil
}

// load reads from the set file into the appData field.
func (d *DataController) load() error {
	dataFile, err := os.Open(d.fileName)
	if err != nil {
		// Probably no data file, set defaults and save.
		d.SetAllowMultiple(true)
		d.appData.ActiveConfigFileNames = []string{}
		return d.Save()
	}

	dataDecoder := gob.NewDecoder(dataFile)
	err = dataDecoder.Decode(&d.appData)
	if err != nil {
		return err
	}

	err = dataFile.Close()
	if err != nil {
		return err
	}

	return nil
}
