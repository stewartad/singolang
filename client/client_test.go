package client

import (
	"testing"
	"os"
)

var testimage = "docker://godlovedc/lolcow"
var testfilename = "lolcow.sif"

func TestExecute(t *testing.T) {
	testCases := []struct {
		desc	string
		
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

func TestInstance(t *testing.T) {
	testCases := []struct {
		desc	string
		
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
		desc		string
		image		string
		ext			string
		expected	string
	}{
		{
			desc: "lolcow",
			image: testimage,
			ext: "",
			expected: "lolcow.sif",
		},
		{
			desc: "lolcow",
			image: testimage,
			ext: "sif",
			expected: "lolcow.sif",
		},
		{
			desc: "lolcow",
			image: testimage,
			ext: "tar",
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
		desc	string
		image	string
		file	string
		pulldir	string
		expected	string
	}{
		{
			desc:		"pull to tmp",
			image:		testimage,
			file:		testfilename,
			pulldir:	"/tmp/test",
			expected:	testfilename,
			
		},
		{
			desc:		"pull to curent dir w other args",
			image:		testimage,
			file:		"",
			pulldir:	"",
			expected:	testfilename,
		},
		{
			desc:		"pull to /tmp w ext",
			image:		testimage,
			file:		"",
			pulldir:	"/tmp/test",
		},
		{
			desc:		"pull w empty opts",
			image:		"",
			file:		"",
			pulldir:	"/tmp/test",
			expected:	"error",
		},
	}
	client, teardown := NewClient()
	defer teardown(client)
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			opts := PullOptions {
				name:		tC.file,
				pullfolder:	tC.pulldir,
				force: 		true,
			}
			image, err := client.Pull(tC.image, &opts)
			if err, ok := err.(*pullError); ok && tC.expected=="error" {
				
			} else
			if err != nil {
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