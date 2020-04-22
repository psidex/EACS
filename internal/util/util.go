package util

import (
	"os"
	"tawesoft.co.uk/go/dialog"
)

// Find takes a slice and looks for an element in it. If found it will return true, else false.
// https://golangcode.com/check-if-element-exists-in-slice/
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// fatalError shows an alert box to the user and exits the program with code 1
func FatalError(errMsg string) {
	dialog.Alert("EACS Fatal Error\n%s", errMsg)
	os.Exit(1)
}
