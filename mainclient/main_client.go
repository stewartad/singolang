package mainclient

import (
	"fmt"
	"github.com/stewartad/singolang/instance"
	"github.com/stewartad/singolang/utils"
)

// Client is a struct to hold information about the current client
type Client struct {
	simage    string // this will be assigned by the load() function
	instances map[string]*instance.Instance
}

// NewClient creates and returns a new client
func NewClient() Client {
	return Client{simage: "", instances: make(map[string]*instance.Instance)}
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

func (c *Client) NewInstance(image string, name string) *instance.Instance {
	i := instance.GetInstance(image, name)
	c.instances[name] = i
	i.Start(false)
	return i
}

func (c *Client) StartInstance(name string) error {
	err := c.instances[name].Start(false)
	return err
}

func (c *Client) StopInstance(name string) error {
	fmt.Printf("Stopping Instance %s...", name)
	err := c.instances[name].Stop(false)
	if err != nil {
		fmt.Printf("FAILED\n")
	} else {
		fmt.Printf("\n")
	}
	return err
}

func (c *Client) StopClientInstances() error {
	var err error
	for k := range c.instances {
		err = c.StopInstance(k)
	}
	return err
}

func (c *Client) PrintInstances() {
	fmt.Println("RUNNING INSTANCES")
	fmt.Println("-----------------")
	for k, v := range c.instances {
		fmt.Printf("%s\t%s\n", k, v)
	}
	fmt.Println("-----------------")
}