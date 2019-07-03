package client

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

func GetFilename(image string, ext string) string{
	if ext == "" {
		ext = "sif"
	}

	image = filepath.Base(image)
	match, err := regexp.Compile("^.*//")
	if err != nil {
		// TODO: Better error handling
		log.Fatalf("%s\n", "bad")
	}
	image = match.ReplaceAllString(image, "")

	if !strings.HasSuffix(image, ext) {
		if strings.HasPrefix(ext, ".") {
			image = strings.Join([]string{image, ext}, "")
		} else {
			image = strings.Join([]string{image, ext}, ".")
		}
	}

	return image
}

func getURI() {
	
}