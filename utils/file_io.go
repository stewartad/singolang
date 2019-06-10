package utils

import (
	"os"
	"log"
)

// Mkdirp simulates mkdir -p UNIX command
func Mkdirp(path string) {
	err := os.MkdirAll(path, 0777)
	if err != nil {
		log.Fatalf("Error creating path %s, exiting.", path)
	}
}