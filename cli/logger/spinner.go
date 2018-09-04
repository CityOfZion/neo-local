package logger

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

var (
	characterSet   = []string{".", "..", "...", "....", "....."}
	updateInterval = 750 * time.Millisecond
)

// NewSpinner creates a new Spinner.
func NewSpinner(textPrefix string) *spinner.Spinner {
	s := spinner.New(characterSet, updateInterval)
	s.Prefix = fmt.Sprintf("%s %s", prefix, textPrefix)
	s.FinalMSG = fmt.Sprintf("%s %s.....DONE\n", prefix, textPrefix)
	return s
}
