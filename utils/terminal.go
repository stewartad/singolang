package utils

import (
	//"unicode/utf8"
	"log"
	"os/exec"
	"strings"
)

// cmd is a slice of strings, the first being the name of the command
// the rest are arguments
func runCommand(cmd []string, sudo bool) string {
	name := cmd[0]
	if sudo {
		name = strings.Join([]string{"sudo", cmd[0]}, " ")
	}
	// create command instance
	process := exec.Command(name, cmd[1:]...)
	// get stdout and stderr
	out, err := process.Output()
	if err != nil {
		log.Fatal(err)
	}
	// convert stdout to string
	var output strings.Builder
	output.Write(out)
	
	// return output
	return output.String();
}

func getSingularityVersion() string {
	version := runCommand([]string{"singularity", "--version"}, false)
	return strings.TrimSpace(version)
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