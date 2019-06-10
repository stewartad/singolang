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

func TestSingularityVersion(t *testing.T) {
	// This is going to be invalidated when the rest of the function is written
	expectedOut := "2.4.2-dist"
	actualOut := getSingularityVersion()
	if strings.Compare(expectedOut, actualOut) != 0 {
		t.Errorf("Singularity versions %s and %s do not match", expectedOut, actualOut)
	}
}