package internal

import (
	"fmt"
	"os"

	"github.com/csutorasa/cli-guide/io"
)

const listUsage = "Usage: cli-guide list"

func ListArgs(args []string) {
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "too many arguments\n%s\n", listUsage)
		os.Exit(1)
	}
	list()
}

func list() {
	sessions, err := io.ListSessions()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if len(sessions) == 0 {
		fmt.Printf("There are no existing sessions\n")
		return
	}
	fmt.Printf("Existing sessions:\n")
	for id, session := range sessions {
		fmt.Printf("%d - %s (step %d)\n", id, session.Guide.Name, session.Guide.Step)
	}
}
