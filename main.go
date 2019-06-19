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

	// if _, err := os.Stat("lolcow_latest.sif"); err == nil {
	// 	utils.RunCommand([]string{"rm", "lolcow_latest.sif"}, false, false)
	// }

	cl, finish := client.NewClient()
	defer finish(&cl)
	// img := client.Pull("docker://godlovedc/lolcow", "", "", "")

	// utils.RunCommand([]string{"ls", "-l", filepath.Dir(img)}, false, false)

	_ = cl.NewInstance("lolcow_latest.sif", "lolcow3")

	out, err := cl.Execute("lolcow3", "which fortune")
	if err != nil {
		fmt.Printf("%s\n%s\n", out, err)
	}
	out, err = cl.Execute("lolcow3", "which singularity")
	if err != nil {
		fmt.Printf("%s\n%s\n", out, err)
	}
	_, err = cl.Execute("lolcow3", "which lolcat")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	cl.ListInstances()
	client.ListAllInstances()
}
