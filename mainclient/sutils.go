package mainclient

import (
	"path/filepath"
	"regexp"
	"log"
	"strings"
)

func load() {

}

func setenv() {

}

func GetFilename(image string, ext string, pwd bool) string{
	if ext == "" {
		ext = "sif"
	}
	if pwd {
		image = filepath.Base(image)
	}
	match, err := regexp.Compile("^.*//")
	if err != nil {
		// TODO: Better error handling
		log.Fatalf("%s\n", "bad")
	}
	image = match.ReplaceAllString(image, "")

	if !strings.HasSuffix(image, ext) {
		image = strings.Join([]string{image, ext}, ".")
	}

	return image
}

func getURI() {
	
}