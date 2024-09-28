package internal

import (
	"fmt"

	"github.com/csutorasa/cli-guide/io"
)

const listUsage = "Usage: cli-guide [-rootDir rootDir] [-q | -v] list"

func ListArgs(args []string) {
	if len(args) != 0 {
		io.PrintFatalError2(fmt.Errorf("too many arguments"), listUsage)
	}
	list()
}

func list() {
	io.PrintVerbose("Listing sessions\n")
	sessions, err := io.ListSessions()
	if err != nil {
		io.PrintFatalError(err)
	}
	if len(sessions) == 0 {
		io.Print("There are no existing sessions\n")
		return
	}
	io.PrintVerbose(fmt.Sprintf("%d existing sessions were found\n", len(sessions)))
	io.Print("Existing sessions:\n")
	for id, session := range sessions {
		io.PrintQuiet(fmt.Sprintf("[%d] %s (step %d)\n", id, session.Guide.Name, session.Guide.Step))
	}
	io.PrintVerbose("Exiting\n")
}
