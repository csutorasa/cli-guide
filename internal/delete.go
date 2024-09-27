package internal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/csutorasa/cli-guide/io"
)

const deleteUsage = "Usage: cli-guide [-rootDir rootDir] delete sessionId"

func DeleteArgs(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "sessionId is expected\n%s\n", deleteUsage)
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "too many arguments\n%s\n", deleteUsage)
		os.Exit(1)
	}
	sessionId, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sessionId must be a number\n%s\n", deleteUsage)
		os.Exit(1)
	}
	delete(int(sessionId))
}

func delete(state int) {
	err := io.DeleteSession(state)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
