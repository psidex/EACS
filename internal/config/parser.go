package config

import (
	"errors"
	"strings"
)

// ParseIncludeText takes a line from a config file and determines if it's an includeText line, if it is, return the
// name of the file that is included.
func ParseIncludeText(line string) (parsedFileName string, err error) {
	// If the line does not (start with "Include:" || end with ".txt")
	if !strings.HasPrefix(line, "Include:") || !strings.HasSuffix(line, ".txt") {
		return "", errors.New("cannot parse IncludeText statement")
	}

	parts := strings.Split(line, "\\")
	fileName := parts[len(parts)-1]
	return fileName, nil
}
