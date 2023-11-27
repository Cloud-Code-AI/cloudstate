package utils

import (
	"fmt"
	"os"
)

// Error handling incase of fatal error
func ExitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
