package internal

import (
	"fmt"
	"os"
	"strconv"

	"github.com/csutorasa/cli-guide/io"
)

const resumeUsage = "Usage: cli-guide [-rootDir rootDir] resume sessionId"

func ResumeArgs(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "sessionId is expected\n%s\n", resumeUsage)
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "too many arguments\n%s\n", resumeUsage)
		os.Exit(1)
	}
	sessionId, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "sessionId must be a number\n%s\n", resumeUsage)
		os.Exit(1)
	}
	resume(int(sessionId), false)
}

func resume(id int, manual bool) {
	session, err := io.ReadSession(id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	guides, err := io.ReadGuideFile(session.Guide.File)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	guide := guides.FindGuideByName(session.Guide.Name)
	if guide == nil {
		fmt.Fprintf(os.Stderr, "could not find guide\n%s", session.Guide.Name)
		os.Exit(1)
	}
	err = executeGuideWriteSessionAndState(guide, session, id, manual)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
