package logger

import (
	"os"
	"strings"
)

const critical = -4
// more vars will go here
const info = 1

// SingularityMessage emits messages to the console
type SingularityMessage struct {

}

// Emit prints a message to the console
func Emit(out *SingularityMessage, level int, message string, prefix string, color string) {

}

func getLoggerLevel() int {
	level := os.Getenv("MESSAGELEVEL")
	if strings.EqualFold(level, "CRITICAL") {
		return critical
	}
	// more logic here
	return 1
}