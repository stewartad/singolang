package utils

import (
	//"unicode/utf8"
	"log"
	"os/exec"
	"strings"
)

// RunCommand runs a terminal command
/*	cmd is a slice of strings, 
	the first element of cmd must be the name of the command
	the remaining elements are arguments 
*/
func RunCommand(cmd []string, sudo bool) string {
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

// GetSingularityVersion gets installed Singularity version
func GetSingularityVersion() string {
	version := RunCommand([]string{"singularity", "--version"}, false)
	return strings.TrimSpace(version)
}

// SplitURI splits the URI into protocol and path"
func SplitURI(container string) (string, string) {
	// Splits
	parts := strings.Split(container, "://")
	return parts[0], parts[1]
}

// RemoveURI strips the protocol and returns only the path
func RemoveURI(container string) string {
	_, path := SplitURI(container)
	return path
}