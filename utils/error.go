package utils

import (
	"fmt"
	"os"
)

// CheckErrorForExit check the error.
// If error is not nil, print the error message and exit the application.
// If error is nil, do nothing.
func CheckErrorForExit(err error, code ...int) {
	if err != nil {
		exitCode := 1
		if len(code) > 0 {
			exitCode = code[0]
		}
		fmt.Printf("snips: %s (%d)\n", err.Error(), exitCode)
		os.Exit(exitCode)
	}
}
