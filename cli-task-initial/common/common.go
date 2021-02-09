package common

import (
	"strings"
)

// FormatCMD - func
func FormatCMD(cmd string) string {
	text := strings.Split(cmd, "\\n")
	format := ""
	for _, e := range text {
		// fmt.Println(e)
		format = format + e + "\n"
	}
	return format
}
