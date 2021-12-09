package main

import (
	"github.com/workcha/crossCompile/common"
	"os"
)

func main() {
	if common.FileExists(".\\project\\" + os.Args[1]) {
		common.CompileAll(os.Args[1])
	}
}
