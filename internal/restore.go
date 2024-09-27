package internal

import (
	"fmt"
	"os"

	"github.com/csutorasa/cli-guide/io"
)

const restoreUsage = "Usage: cli-guide restore"

func RestoreArgs(args []string) {
	if len(args) != 0 {
		fmt.Fprintf(os.Stderr, "too many arguments\n%s\n", listUsage)
		os.Exit(1)
	}
	restore()
}

func restore() {
	state, err := io.ReadState()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if state.LastSession < 1 {
		fmt.Fprintf(os.Stderr, "could not restore: there is no saved state\n")
		os.Exit(1)
	}
	resume(state.LastSession, false)
}
