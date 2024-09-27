package internal

import (
	"fmt"
	"os"

	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

func RestoreOrSelect() {
	state, err := io.ReadState()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if state.LastSession > 0 {
		resume(state.LastSession, false)
		return
	}
	sessions, err := io.ListSessions()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if len(sessions) == 0 {
		fmt.Printf("There are no existing sessions\n")
		return
	}
	id, session, err := io.ScanSelectMapWithZero("Select session to restore", "Cancel", sessions, func(k int, v *model.Session) string {
		return fmt.Sprintf("%s (step %d)", v.Guide.Name, v.Guide.Step)
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if session != nil {
		resume(id, true)
	}
}
