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
func RunCommand(cmd []string, sudo bool) (string, error) {
	if sudo {
		cmd = append([]string{"sudo"}, cmd...)
	}
	name := cmd[0]
	// create command instance
	process := exec.Command(name, cmd[1:]...)
	// get stdout and stderr
	out, err := process.CombinedOutput()
	if err != nil {
		return "", err
	}
	// convert stdout to string
	var output strings.Builder
	output.Write(out)
	
	// return output
	return output.String(), nil;
}

// GetSingularityVersion gets installed Singularity version
func GetSingularityVersion() string {
	version, err := RunCommand([]string{"singularity", "--version"}, false)
	if err != nil {
		log.Fatalln(err)
	}
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