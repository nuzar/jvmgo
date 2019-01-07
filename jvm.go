package main

import (
	"fmt"
	"jvmgo/cmdline"
)

func startJVM(cmd *cmdline.Cmd) {
	fmt.Println("start " + cmd.Class)
}
