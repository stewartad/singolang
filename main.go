package main

import (
	"fmt"
	// "log"
	// "os"
	// "path/filepath"
	"github.com/stewartad/singolang/client"
	"github.com/stewartad/singolang/utils"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("Singularity Version: %s\n", utils.GetSingularityVersion())

	// instantiate a new client and defer its teardown function
	cl, finish := client.NewClient()
	defer finish(&cl)
	// img := client.Pull("docker://godlovedc/lolcow", "", "", "")

	// create a new instance
	instanceError := cl.NewInstance("lolcow_latest.sif", "lolcow3")
	if instanceError != nil {
		fmt.Printf("%s\n", instanceError)
	}

	// Run some executes
	out, err := cl.Execute("lolcow3", "which fortune")
	if err != nil {
		fmt.Printf("%s\n%s\n", out, err)
	}
	// This one is designed to fail
	out, err = cl.Execute("lolcow3", "which singularity")
	if err != nil {
		fmt.Printf("%s\n%s\n", out, err)
	}
	_, err = cl.Execute("lolcow3", "which lolcat")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// List client's stored images
	cl.ListInstances()

	// List all running singularity images
	client.ListAllInstances()
}
