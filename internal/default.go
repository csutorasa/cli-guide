package internal

import (
	"fmt"

	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

func RestoreOrSelect() {
	io.PrintVerbose("Reading state file\n")
	state, err := io.ReadState()
	if err != nil {
		io.PrintFatalError(err)
	}
	if state.LastSession > 0 {
		io.PrintVerbose(fmt.Sprintf("Last session was found\nResuming session %d\n", state.LastSession))
		resume(state.LastSession, false)
		return
	}
	io.PrintVerbose("Last session was not found\nSearching for sessions instead\n")
	sessions, err := io.ListSessions()
	if err != nil {
		io.PrintFatalError(err)
	}
	if len(sessions) == 0 {
		io.Print("No existing sessions were found\n")
		return
	}
	io.PrintVerbose(fmt.Sprintf("%d existing sessions were found\nShwoing selection for resuming\n", len(sessions)))
	id, session, err := io.ScanSelectMapWithZero("Select session to restore", "Cancel", sessions, func(k int, v *model.Session) string {
		return fmt.Sprintf("%s (step %d)", v.Guide.Name, v.Guide.Step)
	})
	if err != nil {
		io.PrintFatalError(err)
	}
	if session == nil {
		io.PrintVerbose("No sessions were selected\nExiting\n")
		return
	}
	io.PrintVerbose(fmt.Sprintf("Resuming selected session %d\n", id))
	resume(id, true)
}
