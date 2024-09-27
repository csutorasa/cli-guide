package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/csutorasa/cli-guide/internal"
	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

var StdinScanner = bufio.NewScanner(os.Stdin)

func main() {
	flags, positional := parseFlags()
	io.SetRootDir(flags.RootDir)
	if len(positional) == 0 {
		internal.RestoreOrSelect()
		return
	}
	switch positional[0] {
	case "create":
		internal.CreateArgs(positional[1:])
	case "delete":
		internal.DeleteArgs(positional[1:])
	case "help":
		internal.HelpArgs(positional[1:])
	case "list":
		internal.ListArgs(positional[1:])
	case "restore":
		internal.ResumeArgs(positional[1:])
	case "resume":
		internal.ResumeArgs(positional[1:])
	default:
		fmt.Fprintf(os.Stderr, "unexpected command %s\n", positional[0])
		os.Exit(1)
	}
}

func parseFlags() (*model.Flags, []string) {
	rootDir := flag.String("rootDir", ".", "directory for storing the state")
	flag.Parse()
	return &model.Flags{
		RootDir: *rootDir,
	}, flag.Args()
}
