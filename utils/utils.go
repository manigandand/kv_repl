package utils

import (
	"errors"
	"strings"
)

const (
	// Tab holds unicode value of tab
	Tab = "\t"
	// Space const holds the space string
	Space = " "
	// NewLineDelim holds unicode value of tab
	NewLineDelim = "\n"
	// EndLineDelim holds unicode value of tab
	EndLineDelim = '\n'
)

var (
	// ErrInvalidTabSpace is returned when the input string contains tab space
	ErrInvalidTabSpace = errors.New("Don't use Tab spaces in the commands")
)

// SplitCmdArguments attempts to split the input string by command and arguments
// Assuming that the string is seperated by space and the first instance is command
// and the rest followed by command is arguments.
func SplitCmdArguments(str string) (res []string, err error) {
	if strings.Contains(str, Tab) {
		err = ErrInvalidTabSpace
		return
	}
	// CMD KEY VALUE...
	return strings.SplitN(str, Space, 3), nil
}
