package client

import (
	"fmt"
	"log"
	"strings"
	"github.com/stewartad/singolang/utils"
)

// Client is a struct to hold information about the current client
type Client struct {
	simage    string // this will be assigned by the load() function
	instances map[string]*Instance
}

// NewClient creates and returns a new client
func NewClient() (Client, func (c *Client)) {
	return Client{simage: "", instances: make(map[string]*Instance)}, func (c *Client) {
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

// StartInstance starts an instance that was previously created in the client
// TODO: Define custom errors
func (c *Client) StartInstance(i *Instance) error {
	fmt.Printf("Starting Instance %s...\n", i.name)
	err := i.start(false)
	if err != nil {
		fmt.Printf("FAILED\n")
	} else {
		fmt.Printf("SUCCESS. %s\n", i)
	}
	return err
}

// StopInstance stops an instance previously created in the client
// TODO: Define custom errors
func (c *Client) StopInstance(name string) error {
	fmt.Printf("Stopping Instance %s...", name)
	err := c.instances[name].stop(false)
	if err != nil {
		fmt.Printf("FAILED\n")
	} else {
		fmt.Printf("\n")
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

// PrintInstances prints all client-created instances to screen
func (c *Client) PrintInstances() {
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

// ListInstances lists all currently running Singularity instances.
// It is equivalent to running `singularity instance list`
func ListInstances() {
	cmd := utils.InitCommand("instance", "list")

	output, err := utils.RunCommand(cmd, false, false)
	_ = output
	if err != nil {
		log.Printf("Error running command: %s\n", strings.Join(cmd, " "))
	}
}


func (c *Client) teardown() {
	c.StopAllInstances()
	ListInstances()
}
