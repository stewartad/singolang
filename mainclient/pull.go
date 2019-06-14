package mainclient

import (
	"github.com/stewartad/singolang/utils"
	"path/filepath"
	"strings"
	"regexp"
	"log"
	"fmt"
	"os"
)

// Pull pulls image from singularityhub or dockerhub and builds it.
// It stores the image in pullfolder, naming it name.ext
func (c *Client) Pull(image string, name string, ext string, pullfolder string) string {
	cmd := InitCommand("pull")
	match, err := regexp.MatchString("^(shub|docker)://", image)
	if err != nil {
		log.Fatalf("Something went wrong")
	}
	if !match {
		log.Fatalln("Pull only valid for singularity hub and docker hub")
	}

	if name == "" {
		name = GetFilename(image, ext, false)
	}

	// cmd = append(cmd, "--name")
	// cmd = append(cmd, name)

	cmd = append(cmd, image)

	fmt.Printf("%s\n", strings.Join(cmd, " "))

	utils.RunCommand(cmd, false, false)

	finalImage := filepath.Join(pullfolder, filepath.Base(name))
	name = finalImage
	if os.Stat(finalImage); err == nil {
		fmt.Println(finalImage)
	}

	return finalImage
}