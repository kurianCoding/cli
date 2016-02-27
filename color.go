// +build !windows

package cli

import (
	"fmt"
)

func gray(format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[90m"+format+"\x1b[0m", args...)
}

func red(format string, args ...interface{}) string {
	return fmt.Sprintf("\x1b[31m"+format+"\x1b[0m", args...)
}