package mainclient

import (
	"github.com/stewartad/singolang/utils"
	"fmt"
)

// Client is a struct to hold information about the current client
type Client struct {
	simage string // this will be assigned by the load() function
}

// GetClient creates and returns a new client
func GetClient() Client {
	return Client{simage: ""}
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
