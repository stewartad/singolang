package client

import (
	"github.com/stewartad/singolang/utils"
	"path/filepath"
	"strings"
	"regexp"
	"fmt"
	"os"
)

type PullOptions struct {
	name		string
	ext			string
	pullfolder	string
	force		bool
}

type pullError struct {
	msg	string
}

func (e *pullError) Error() string {
	return e.msg
}

// Pull pulls image from singularityhub or dockerhub and builds it.
// It stores the image in pullfolder, naming it name.ext
func (c *Client) Pull(image string, opts *PullOptions) (string, error) {
	cmd := utils.InitCommand("pull")

	if opts.force {
		cmd = append(cmd, "-F")
	}

	match, err := regexp.MatchString("^(shub|docker)://", image)
	if err != nil {
		return "", err
	}
	if !match {
		return "", &pullError{msg: "Pull only valid for singularity hub and docker hub"}
	}

	name := opts.name

	if opts.name == "" {
		name = GetFilename(image, "")
	}
	if opts.pullfolder != "" {
		name = filepath.Join(opts.pullfolder, name)
	}

	cmd = append(cmd, name)
	cmd = append(cmd, image)

	fmt.Printf("%s\n", strings.Join(cmd, " "))

	utils.RunCommand(cmd, utils.DefaultRunCommandOptions())

	finalImage := filepath.Join(opts.pullfolder, filepath.Base(name))

	_, err = os.Stat(finalImage)
	if err != nil {
		return "", err
	}
	fmt.Println(finalImage)
	return finalImage, nil
}