package logger

import (
	"fmt"
)

var (
	prefix = ""
)

func setPrefix(name string, version string) {
	prefix = fmt.Sprintf("%s (%s):", name, version)
}
