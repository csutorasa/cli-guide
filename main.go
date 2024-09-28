package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/csutorasa/cli-guide/internal"
	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

const RootDirEnv = "CLI_GUIDE_ROOT_DIR"

func main() {
	flags, positional := parseFlags()
	setLogLevel(flags)
	io.PrintVerbose(fmt.Sprintf("Flags: %s\n", flags))
	rootDir, err := setRootDir(flags)
	if err != nil {
		io.PrintFatalError(fmt.Errorf("invalid root directory: %s\n%s", err.Error(), flags.RootDir))
		os.Exit(1)
	}
	io.PrintVerbose(fmt.Sprintf("Root dir: %s\n", rootDir))
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
		internal.RestoreArgs(positional[1:])
	case "resume":
		internal.ResumeArgs(positional[1:])
	default:
		io.PrintFatalError2(fmt.Errorf("unexpected command %s", positional[0]), internal.Usage)
	}
}

func setRootDir(flags *model.Flags) (string, error) {
	rootDir := flags.RootDir
	if rootDir == "" {
		env, ok := os.LookupEnv(RootDirEnv)
		if ok {
			rootDir = env
		} else {
			home, err := os.UserHomeDir()
			if err == nil {
				rootDir = filepath.Join(home, ".cli-guide")
			} else {
				rootDir = "."
			}
		}
	}
	rootDir, _ = filepath.Abs(rootDir)
	return rootDir, io.SetRootDir(rootDir)
}

func setLogLevel(flags *model.Flags) {
	if flags.Quiet {
		io.SetLogLevel(io.LogLevelQuiet)
		return
	}
	if flags.Verbose {
		io.SetLogLevel(io.LogLevelVerbose)
		return
	}
	io.SetLogLevel(io.LogLevelNormal)
}

func parseFlags() (*model.Flags, []string) {
	rootDir := ""
	verbose := false
	quiet := false
	flag.StringVar(&rootDir, "rootDir", "", "directory for storing the state")
	flag.BoolVar(&verbose, "v", false, "verbose logging")
	flag.BoolVar(&quiet, "q", false, "quiet logging")
	flag.Parse()
	return &model.Flags{
		RootDir: rootDir,
		Quiet:   quiet,
		Verbose: verbose,
	}, flag.Args()
}
