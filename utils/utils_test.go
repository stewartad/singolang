package utils

import (
	"testing"
	"strings"
)

var TestPath = "docker://vsoch/hello-world"
var ExpectedProtocol = "docker"
var ExpectedPath = "vsoch/hello-world"

// TestSplitURI tests utils.SplitURI
func TestSplitURI(t *testing.T) {
	protocol, path := splitURI(TestPath)
	if strings.Compare(protocol, ExpectedProtocol) != 0 {
		t.Errorf("Incorrect path split, expected protocol %s, got %s", ExpectedProtocol, protocol)
	}
	if strings.Compare(path, ExpectedPath) != 0 {
		t.Errorf("Incorrect path split, expected path %s, got %s", ExpectedPath, path)
	}
}

// TestRemoveURI tests utils.RemoveURI
func TestRemoveURI(t *testing.T) {
	path := removeURI(TestPath)
	if strings.Compare(path, ExpectedPath) != 0 {
		t.Errorf("Incorrect path, expected path %s, got %s", ExpectedPath, path)
	}
}

func TestRunCommand(t *testing.T) {
	expectedCmd := "sudo touch abc"
	actualCmd := runCommand("touch abc", true, true)
	if strings.Compare(expectedCmd, actualCmd) != 0 {
		t.Errorf("Commands %s and %s do not match", expectedCmd, actualCmd)
	}
}