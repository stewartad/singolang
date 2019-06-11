package main

import (
	"fmt"
	"github.com/stewartad/singolang/utils"
)

func main() {
	fmt.Println("hello")
	utils.Mkdirp("./testdir/two")
	utils.RunCommand([]string{"touch", "./testdir/a", "./testdir/b", "./testdir/c"}, false)
	fmt.Printf("Singularity Version: %s\n", utils.GetSingularityVersion())
	fmt.Printf(utils.RunCommand([]string{"ls", "-l", "./testdir"}, false))
	utils.RunCommand([]string{"rm", "-r", "./testdir"}, false)
}