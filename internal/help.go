package internal

import (
	"fmt"

	"github.com/csutorasa/cli-guide/io"
)

const Usage = "Usage: cli-guide [-rootDir rootDir] [-q | -v] [command] [params]"
const helpUsage = "Usage: cli-guide help command"

func HelpArgs(args []string) {
	if len(args) == 0 {
		io.PrintQuiet(helpUsage + "\n")
		io.PrintQuiet(Usage + "\n")
		io.PrintQuiet("Commands:\n")
		io.PrintQuiet("  create  - creates a new session from a guide\n")
		io.PrintQuiet("  delete  - deletes a session\n")
		io.PrintQuiet("  help    - displays this message\n")
		io.PrintQuiet("  list    - lists all sessions\n")
		io.PrintQuiet("  restore - resumes the last used session\n")
		io.PrintQuiet("  resume  - resumes a session\n")
		io.PrintQuiet("If no command is provided restore will be attempted.\n")
		io.PrintQuiet("If it fails a selection is displayed for resuming.\n")
		io.PrintQuiet("Flags:\n")
		io.PrintQuiet("  q - quiet output\n")
		io.PrintQuiet("  v - verbose output\n")
		io.PrintQuiet("  rootDir - directory for storing the state, defaults to CLI_GUIDE_ROOT_DIR environment variable, then to HOME/.cli-guide, then to current directory\n")
		return
	}
	if len(args) > 1 {
		io.PrintFatalError2(fmt.Errorf("too many arguments"), helpUsage)
	}
	switch args[0] {
	case "create":
		io.PrintQuiet(createUsage + "\n")
	case "delete":
		io.PrintQuiet(deleteUsage + "\n")
	case "help":
		io.PrintQuiet(helpUsage + "\n")
	case "list":
		io.PrintQuiet(listUsage + "\n")
	case "restore":
		io.PrintQuiet(restoreUsage + "\n")
	case "resume":
		io.PrintQuiet(resumeUsage + "\n")
	default:
		io.PrintFatalError2(fmt.Errorf("unexpected command %s", args[0]), helpUsage)
	}
}
