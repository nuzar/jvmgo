package main

import (
	"jvmgo/cmdline"
	"os"
)

func main() {
	cmd, err := cmdline.ParseCmd(os.Args)
	if err != nil {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}
