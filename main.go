package main

import (
	"fmt"
	"github.com/stewartad/singolang/utils"
)

func main() {
	fmt.Println("hello")
	utils.Mkdirp("./testdir/two/three")
}