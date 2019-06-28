package client

import (
	"fmt"
	"github.com/stewartad/singolang/utils"
	"log"
	"path/filepath"
	"strings"
	"io/ioutil"
	"compress/gzip"
	"bytes"
	"archive/tar"
)

type instanceError struct {
	name   string
	action string
}

func (e *instanceError) Error() string {
	return fmt.Sprintf("Error %sing Instance: %s. Must be stopped manually in Singularity", e.name, e.action)
}

// Client is a struct to hold information about the current client
type Client struct {
	simage    string // this will be assigned by the load() function
	instances map[string]*Instance
}

// NewClient creates and returns a new client as well as a teardown function.
// Assign this teardown function and defer it to exit cleanly
func NewClient() (*Client, func(c *Client)) {
	return &Client{simage: "", instances: make(map[string]*Instance)}, func(c *Client) {
		c.teardown()
	}
}

// Version returns the version of the system's Singularity installation
func (c *Client) Version() string {
	return utils.GetSingularityVersion()
}

func (c *Client) String() string {
	baseClient := "[singularity-golang]"
	if c.simage != "" {
		baseClient = fmt.Sprintf("%s[%s]", baseClient, c.simage)
	}
	return baseClient
}

// NewInstance creates a new instance and adds it to the client, if it is able to be started
func (c *Client) NewInstance(image string, name string) error {
	i := getInstance(image, name)
	err := i.start(false)
	if err != nil {
		return err
	}
	c.instances[name] = i
	return nil
}

// CopyTarball creates a Tar archive of a directory or file and places it in /tmp. 
// It returns the path to the archive, and a reader for the archive
func (c *Client) CopyTarball(instance string, path string) (string, *tar.Reader, error) {
	// Make directory for archive and set up filepath
	parentDir := filepath.Dir(path)
	dir := fmt.Sprintf("/tmp/%s", instance)
	utils.Mkdirp(dir)
	archivePath := fmt.Sprintf("%s/%s-archive.tar.gz", dir, filepath.Base(parentDir))

	// Create archive
	cmd := []string{"tar", "-czvf", archivePath, path}
	_, _, code, err := c.Execute(instance, cmd, DefaultExecOptions())
	if err != nil || code != 0 {
		return "", nil, err
	}

	// Create reader for archive
	b, err := ioutil.ReadFile(archivePath)
	if err != nil {
		panic(fmt.Sprintf("Could not read file %s", err))
	}
	gzr, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		panic("READ ERROR")
	}

	return archivePath, tar.NewReader(gzr), nil
}

// StopInstance stops an instance previously created in the client
// TODO: Define custom errors
func (c *Client) StopInstance(name string) error {
	// fmt.Printf("Stopping Instance %s...", name)
	err := c.instances[name].stop(false)
	if err != nil {
		// fmt.Printf("FAILED\n")
	} else {
		// fmt.Printf("\n")
	}
	return err
}

// StopAllInstances stops all instances created in the client
func (c *Client) StopAllInstances() error {
	var err error
	for k := range c.instances {
		err = c.StopInstance(k)
	}
	return err
}

// ListInstances prints all client-created instances to screen
func (c *Client) ListInstances() {
	fmt.Println("CLIENT LOADED INSTANCES")
	fmt.Println("-----------------")
	if len(c.instances) < 1 {
		fmt.Println("No Loaded Instances\n-----------------")
		return
	}
	for k, v := range c.instances {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println("-----------------")
}

// ListAllInstances lists all currently running Singularity instances.
// It is equivalent to running `singularity instance list`
func ListAllInstances() {
	cmd := utils.InitCommand("instance", "list")

	output, stderr, status, err := utils.RunCommand(cmd, false, false)
	// TODO: do something with these values
	_, _, _ = output, status, stderr
	if err != nil {
		log.Printf("Error running command: %s\n", strings.Join(cmd, " "))
	}
}

func (c *Client) teardown() {
	fmt.Println("Performing Cleanup")
	c.StopAllInstances()
	ListAllInstances()
}

