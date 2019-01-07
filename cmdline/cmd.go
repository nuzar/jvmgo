package cmdline

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
	XjreOption  string
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}

func ParseCmd(args []string) (*Cmd, error) {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	flag.Parse()

	if len(args) > 1 {
		cmd.Class = args[1]
		cmd.Args = args[2:]
	}

	return cmd, nil
}
