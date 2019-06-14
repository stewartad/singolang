package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/stewartad/singolang/utils"
	"github.com/stewartad/singolang/mainclient"
	"github.com/stewartad/singolang/instance"
)

func main() {
	fmt.Println("hello")
	fmt.Printf("Singularity Version: %s\n", utils.GetSingularityVersion())

	if _, err := os.Stat("lolcow_latest.sif"); err == nil {
		utils.RunCommand([]string{"rm", "lolcow_latest.sif"}, false, false)
	}
	
	client := mainclient.GetClient()
	img := client.Pull("docker://godlovedc/lolcow", "", "", "")
	utils.RunCommand([]string{"ls", "-l", filepath.Dir(img)}, false, false)

	i := instance.GetInstance("lolcow_latest.sif")
	i.Start("", "lolcow1", false)

	utils.RunCommand([]string{"singularity", "instance", "list"}, false, false)

	i.Stop("", false)

	utils.RunCommand([]string{"singularity", "instance", "list"}, false, false)
}