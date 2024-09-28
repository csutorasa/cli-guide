package internal

import (
	"fmt"

	"github.com/csutorasa/cli-guide/io"
)

const restoreUsage = "Usage: cli-guide [-rootDir rootDir] [-q | -v] restore"

func RestoreArgs(args []string) {
	if len(args) != 0 {
		io.PrintFatalError2(fmt.Errorf("too many arguments"), restoreUsage)
	}
	restore()
}

func restore() {
	io.PrintVerbose("Reading state\n")
	state, err := io.ReadState()
	if err != nil {
		io.PrintFatalError(err)
	}
	if state.LastSession < 1 {
		io.PrintFatalError(fmt.Errorf("could not restore: there is no saved state"))
	}
	io.PrintVerbose(fmt.Sprintf("Last session was found\nResuming session %d\n", state.LastSession))
	resume(state.LastSession, false)
}
