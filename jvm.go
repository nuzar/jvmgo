package main

import (
	"fmt"
	"jvmgo/classpath"
	"jvmgo/cmdline"
	"strings"
)

func startJVM(cmd *cmdline.Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath: %v class: %v args:%v\n",
		cp, cmd.Class, cmd.Args)

	classname := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(classname)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.Class)
		return
	}

	fmt.Printf("class data: %v\n", classData)
}
