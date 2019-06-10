package utils

import (
	//"unicode/utf8"
	//"os"
	"strings"
)

func runCommand(cmd string, sudo bool, capture bool) string {
	if sudo {
		cmd = strings.Join([]string{"sudo", cmd}, " ")
	}
	return cmd;
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