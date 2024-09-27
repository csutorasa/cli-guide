package internal

import (
	"fmt"
	"os"
)

const helpUsage = "Usage: cli-guide help command"

func HelpArgs(args []string) {
	if len(args) == 0 {
		fmt.Println(helpUsage)
		fmt.Println("Usage: cli-guide [-rootDir rootDir] [command] [params]")
		fmt.Println("Commands:")
		fmt.Println("  create  - creates a new session from a guide")
		fmt.Println("  delete  - deletes a session")
		fmt.Println("  help    - displays this message")
		fmt.Println("  list    - lists all sessions")
		fmt.Println("  restore - resumes the last used session")
		fmt.Println("  resume  - resumes a session")
		fmt.Println("If no command is provided restore will be attempted.")
		fmt.Println("If it fails a selection is displayed for resuming.")
		fmt.Println("Flags:")
		fmt.Println("  rootDir - directory for storing the state")
		return
	}
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "too many arguments\n%s\n", helpUsage)
		os.Exit(1)
	}
	switch args[0] {
	case "create":
		fmt.Println(createUsage)
	case "delete":
		fmt.Println(deleteUsage)
	case "help":
		fmt.Println(helpUsage)
	case "list":
		fmt.Println(listUsage)
	case "restore":
		fmt.Println(restoreUsage)
	case "resume":
		fmt.Println(resumeUsage)
	default:
		fmt.Fprintf(os.Stderr, "unexpected command %s\n", args[0])
		os.Exit(1)
	}
}
