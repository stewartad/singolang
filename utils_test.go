package singolang

import (
	"strings"
	"testing"
)

var TestPath = "docker://godlovedc/lolcow"
var ExpectedProtocol = "docker"
var ExpectedPath = "godlovedc/lolcow"

// TestSplitURI tests utils.SplitURI
func TestSplitURI(t *testing.T) {
	protocol, path := SplitURI(TestPath)
	if strings.Compare(protocol, ExpectedProtocol) != 0 {
		t.Errorf("Incorrect path split, expected protocol %s, got %s", ExpectedProtocol, protocol)
	}
	if strings.Compare(path, ExpectedPath) != 0 {
		t.Errorf("Incorrect path split, expected path %s, got %s", ExpectedPath, path)
	}
}

// TestRemoveURI tests utils.RemoveURI
func TestRemoveURI(t *testing.T) {
	path := RemoveURI(TestPath)
	if strings.Compare(path, ExpectedPath) != 0 {
		t.Errorf("Incorrect path, expected path %s, got %s", ExpectedPath, path)
	}
}
