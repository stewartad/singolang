package mainclient

import (
	"github.com/stewartad/singolang/utils"
	"strings"
)

// Client is a struct to hold information about the current client
type Client struct {
	simage string // the type of this will likely change to type Image
}

// GetClient creates and returns a new client
func GetClient() Client {
	return Client{simage: ""}
}

func (c *Client) Version() string{
	return utils.GetSingularityVersion()
}

func (c *Client) String() string {
	var b strings.Builder
	b.WriteString("[singularity-golang]")
	if c.simage != "" {
		b.WriteString("[")
		b.WriteString(c.simage)
		b.WriteString("]")
	}
	return b.String()
}
