package main

import (
	"fmt"
	"jvmgo/cmdline"
	"os"
)

func main() {
	cmd, _ := cmdline.ParseCmd(os.Args)
	if cmd.VersionFlag {
		fmt.Println("Version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		cmdline.PrintUsage()
	} else {
		startJVM(cmd)
	}
}
