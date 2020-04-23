package config

import (
	"io/ioutil"
	"strings"
)

// writeLinesToMasterConfig takes a slice of strings and writes them as lines to the master config.txt file.
func writeLinesToMasterConfig(lines []string) (err error) {
	completeData := []byte(strings.Join(lines, "\n"))
	if err := ioutil.WriteFile(masterConfigFilePath, completeData, 0644); err != nil {
		return err
	}
	return nil
}

// readLinesFromMasterConfig returns all the line in the master config.txt file.
func readLinesFromMasterConfig() (lines []string, err error) {
	includes, err := ioutil.ReadFile(masterConfigFilePath)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(includes), "\n"), nil
}
