package client

import (
	"fmt"
	"log"
	"strings"
	//"github.com/stewartad/singolang/instance"
	"github.com/stewartad/singolang/utils"
)

// Client is a struct to hold information about the current client
type Client struct {
	simage    string // this will be assigned by the load() function
	currInstance	*instance
}

// NewClient creates and returns a new client
func NewClient() Client {
	return Client{simage: "", currInstance: nil}
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

func (c *Client) newInstance(image string, name string) *instance {
	i := getInstance(image, name)
	c.currInstance = i
	return i
}

func (c *Client) StartInstance() error {
	err := c.currInstance.start(false)
	return err
}

func (c *Client) StopInstance(name string) error {
	fmt.Printf("Stopping Instance %s...", name)
	err := c.currInstance.stop(false)
	if err != nil {
		fmt.Printf("FAILED\n")
	} else {
		fmt.Printf("\n")
	}
	return err
}

func (c *Client) PrintInstances() {
	fmt.Println("CLIENT LOADED INSTANCES")
	fmt.Println("-----------------")
	
	fmt.Println("-----------------")
}

func ListInstances() {
	cmd := utils.InitCommand("instance", "list")

	output, err := utils.RunCommand(cmd, false, false)
	_ = output
	if err != nil {
		log.Printf("Error running command: %s\n", strings.Join(cmd, " "))
	}
}