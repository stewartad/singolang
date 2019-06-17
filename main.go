package main

import (
	"fmt"
	// "log"
	// "os"
	// "path/filepath"
	// "github.com/stewartad/singolang/instance"
	"github.com/stewartad/singolang/mainclient"
	"github.com/stewartad/singolang/utils"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("Singularity Version: %s\n", utils.GetSingularityVersion())

	// if _, err := os.Stat("lolcow_latest.sif"); err == nil {
	// 	utils.RunCommand([]string{"rm", "lolcow_latest.sif"}, false, false)
	// }

	client := mainclient.NewClient()
	// img := client.Pull("docker://godlovedc/lolcow", "", "", "")

	// utils.RunCommand([]string{"ls", "-l", filepath.Dir(img)}, false, false)

	client.NewInstance("lolcow_latest.sif", "lolcow1")
	// // err := i.Start(false)
	// if err != nil {
	// 	log.Printf("%s\n%s", err, "")
	// }
	client.Execute("instance://lolcow1", "which fortune")
	client.Execute("instance://lolcow1", "which singularity")
	client.Execute("instance://lolcow1", "which lolcat")

	utils.RunCommand([]string{"singularity", "instance", "list"}, false, false)
	client.PrintInstances()
	// i.Stop(false)

	utils.RunCommand([]string{"singularity", "instance", "list"}, false, false)
}
