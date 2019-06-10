package utils

import (
	//"unicode/utf8"
	"log"
	"os/exec"
	"strings"
)

// cmd is a slice of strings, the first being the name of the command
// the rest are arguments
func runCommand(cmd []string, sudo bool, capture bool) string {
	name := cmd[0]
	if sudo {
		name = strings.Join([]string{"sudo", cmd[0]}, " ")
	}
	process := exec.Command(name, cmd[1:]...)
	out, err := process.Output()
	if err != nil {
		log.Fatal(err)
	}
	var output strings.Builder
	output.Write(out)
	return output.String();
}

// SplitURI splits the URI into protocol and path"
func splitURI(container string) (string, string) {
	// Splits
	parts := strings.Split(container, "://")
	return parts[0], parts[1]
}

// RemoveURI strips the protocol and returns only the path
func removeURI(container string) string {
	_, path := splitURI(container)
	return path
}