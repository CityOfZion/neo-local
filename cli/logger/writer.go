package logger

import (
	"fmt"
)

type (
	// Writer is a custom logger.
	Writer struct{}
)

// NewWriter creates a new Writer.
func NewWriter(name string, version string) Writer {
	setPrefix(name, version)
	return Writer{}
}

func (w Writer) Write(bytes []byte) (int, error) {
	output := fmt.Sprintf("%s %s", prefix, string(bytes))
	return fmt.Print(output)
}
