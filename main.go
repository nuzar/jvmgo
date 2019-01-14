package main

import (
	"fmt"
	"jvmgo/cmdline"
)

func main() {
	cmd, _ := cmdline.ParseCmd()
	if cmd.VersionFlag {
		fmt.Println("Version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}
