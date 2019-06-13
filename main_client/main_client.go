package main_client

import (
	"github.com/stewartad/singolang/utils"
	"path/filepath"
	"strings"
	"regexp"
	"log"
	"fmt"
	"os"
)

// Client is a struct to hold information about the current client
type Client struct {
	simage string // the type of this will likely change to type Image
}

// GetClient creates and returns a new client
func GetClient() Client {
	return Client{simage: ""}
}

func (c Client) Version() string{
	return utils.GetSingularityVersion()
}

func (c Client) String() string {
	var b strings.Builder
	b.WriteString("[singularity-golang]")
	if c.simage != "" {
		b.WriteString("[")
		b.WriteString(c.simage)
		b.WriteString("]")
	}
	return b.String()
}

func (c Client) Pull(image string, name string, ext string, pullfolder string) string {
	cmd := initCommand("pull")
	match, err := regexp.MatchString("^(shub|docker)://", image)
	if err != nil {
		log.Fatalf("why")
	}
	if !match {
		log.Fatalln("pull only valid for singularity hub and docker hub")
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