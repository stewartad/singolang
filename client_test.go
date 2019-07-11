package singolang

import (
	"os"
	"strings"
	"testing"
)

var testPull = "docker://godlovedc/lolcow"
var testfilename = "lolcow.sif"
var testImage = "lolcow_latest.sif"

func TestExecute(t *testing.T) {
	testCases := []struct {
		desc    string
		cmd     []string
		expSout string
		expSerr string
		expCode int
		expErr  bool
	}{
		{
			desc:    "basic echo",
			cmd:     []string{"echo", "hello"},
			expSout: "hello",
			expSerr: "",
			expCode: 0,
			expErr:  false,
		},
		{
			desc:    "exit error",
			cmd:     []string{"grep"},
			expSout: "",
			expSerr: "Usage:",
			expCode: -1,
			expErr:  true,
		},
		{
			desc:    "non-zero return",
			cmd:     []string{"false"},
			expSout: "",
			expSerr: "",
			expCode: -1,
			expErr:  true,
		},
		{
			desc:    "zero return",
			cmd:     []string{"true"},
			expSout: "",
			expSerr: "",
			expCode: 0,
			expErr:  false,
		},
	}

	c, teardown := NewClient()
	err := c.NewInstance(testImage, "lolcow_test")
	if err != nil {
		t.Errorf("Error creating insctance")
	}
	defer teardown(c)

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			stdout, stderr, code, err := c.Execute("lolcow_test", tC.cmd, DefaultExecOptions())
			if !strings.Contains(strings.TrimSpace(stdout), tC.expSout) {
				t.Errorf("Unexpected Stdout: %s", stdout)
			} else if !strings.Contains(strings.TrimSpace(stderr), tC.expSerr) {
				t.Errorf("Unexpected Stderr: %s", stderr)
			} else if code != tC.expCode {
				t.Errorf("Unexpected Return Code: %d", code)
			} else if err == nil && tC.expErr {
				t.Errorf("Expected Error but didn't get one")
			} else if err != nil && !tC.expErr {
				t.Errorf("Unexpected Error")
			}

		})
	}

}

func TestInstance(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestGetFilename(t *testing.T) {
	testCases := []struct {
		desc     string
		image    string
		ext      string
		expected string
	}{
		{
			desc:     "lolcow",
			image:    testPull,
			ext:      "",
			expected: "lolcow.sif",
		},
		{
			desc:     "lolcow",
			image:    testPull,
			ext:      "sif",
			expected: "lolcow.sif",
		},
		{
			desc:     "lolcow",
			image:    testPull,
			ext:      "tar",
			expected: "lolcow.tar",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			name := GetFilename(tC.image, tC.ext)
			if name != tC.expected {
				t.Errorf("Unexpected Filename")
			}
		})
	}
}

func TestPull(t *testing.T) {
	testCases := []struct {
		desc     string
		image    string
		file     string
		pulldir  string
		expected string
	}{
		{
			desc:     "pull to tmp",
			image:    testPull,
			file:     testfilename,
			pulldir:  "/tmp/test",
			expected: testfilename,
		},
		{
			desc:     "pull w empty opts",
			image:    testPull,
			file:     "",
			pulldir:  "",
			expected: "lolcow.sif",
		},
		{
			desc:     "pull w no image",
			image:    "",
			file:     "",
			pulldir:  "",
			expected: "error",
		},
	}
	client, teardown := NewClient()
	defer teardown(client)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			opts := PullOptions{
				Name:       tC.file,
				Pullfolder: tC.pulldir,
				Force:      true,
			}
			image, err := client.Pull(tC.image, &opts)
			if err, ok := err.(*pullError); ok && tC.expected == "error" {

			} else if err != nil {
				t.Errorf("Pull failed: %s not found", image)
			} else {
				os.RemoveAll(image)
			}
		})
	}
}

func setupClient() func(t *testing.T) {
	return nil
}
