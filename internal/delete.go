package internal

import (
	"fmt"
	"strconv"

	"github.com/csutorasa/cli-guide/io"
)

const deleteUsage = "Usage: cli-guide [-rootDir rootDir] [-q | -v] delete sessionId"

func DeleteArgs(args []string) {
	if len(args) == 0 {
		io.PrintFatalError2(fmt.Errorf("sessionId is expected"), deleteUsage)
	}
	if len(args) > 1 {
		io.PrintFatalError2(fmt.Errorf("too many arguments"), deleteUsage)
	}
	sessionId, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		io.PrintFatalError2(fmt.Errorf("sessionId must be a number"), deleteUsage)
	}
	delete(int(sessionId))
}

func delete(id int) {
	io.PrintVerbose(fmt.Sprintf("Deleting session %d\n", id))
	err := io.DeleteSession(id)
	if err != nil {
		io.PrintFatalError(err)
	}
	io.Print(fmt.Sprintf("Session %d was deleted\n", id))
	io.PrintVerbose("Exiting\n")
}
