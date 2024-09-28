package internal

import (
	"fmt"
	"strconv"

	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

const resumeUsage = "Usage: cli-guide [-rootDir rootDir] [-q | -v] resume sessionId"

func ResumeArgs(args []string) {
	if len(args) == 0 {
		io.PrintFatalError2(fmt.Errorf("sessionId is expected"), resumeUsage)
	}
	if len(args) > 1 {
		io.PrintFatalError2(fmt.Errorf("too many arguments"), resumeUsage)
	}
	sessionId, err := strconv.ParseInt(args[0], 10, 32)
	if err != nil {
		io.PrintFatalError2(fmt.Errorf("sessionId must be a number"), resumeUsage)
	}
	resume(int(sessionId), false)
}

func resume(id int, manual bool) {
	io.PrintVerbose(fmt.Sprintf("Reading session %d\n", id))
	session, err := io.ReadSession(id)
	if err != nil {
		io.PrintFatalError(err)
	}
	io.PrintVerbose(fmt.Sprintf("Reading guide file %s\n", session.Guide.File))
	guides, err := io.ReadGuideFile(session.Guide.File)
	if err != nil {
		io.PrintFatalError(err)
	}
	io.PrintVerbose(fmt.Sprintf("Searching for guide by name %s\n", session.Guide.Name))
	guide := guides.FindGuideByName(session.Guide.Name)
	if guide == nil {
		io.PrintFatalError(fmt.Errorf("could not find guide\n%s", session.Guide.Name))
	}
	io.PrintVerbose("Guide was found\nResuming execution\n")
	err = executeGuideWriteSessionAndState(guide, session, id, manual)
	if err != nil {
		io.PrintFatalError(err)
	}
	io.PrintVerbose("Execution has finished\nExiting\n")
}

func executeGuideWriteSessionAndState(guide *model.Guide, session *model.Session, sessionId int, manual bool) error {
	ended, err := executeGuide(guide, session, sessionId, manual)
	if err != nil {
		return err
	}
	return writeSessionAndState(session, sessionId, ended)
}

func writeSessionAndState(session *model.Session, sessionId int, ended bool) error {
	io.PrintVerbose(fmt.Sprintf("Updating session %d\n", sessionId))
	err := io.WriteSession(sessionId, session)
	if err != nil {
		return err
	}
	io.PrintVerbose(fmt.Sprintf("Session %d was updated\n", sessionId))
	newSessionId := sessionId
	if ended {
		newSessionId = 0
		io.PrintVerbose(fmt.Sprintf("Showing decision for deleting current finished session %d\n", sessionId))
		deleteSession, err := io.ScanBoolDefault("Do you want to delete this session", true)
		if err != nil {
			return err
		}
		if deleteSession {
			io.PrintVerbose(fmt.Sprintf("Deleting session %d\n", sessionId))
			err = io.DeleteSession(sessionId)
			if err != nil {
				return err
			}
			io.Print(fmt.Sprintf("Session %d was deleted\n", sessionId))
		} else {
			io.PrintVerbose(fmt.Sprintf("Session %d will not be deleted\n", sessionId))
		}
	}
	io.PrintVerbose("Updating state\n")
	err = io.WriteState(&model.State{
		LastSession: newSessionId,
	})
	if err != nil {
		return err
	}
	io.PrintVerbose("State was updated\nExiting\n")
	return nil
}

func executeGuide(guide *model.Guide, session *model.Session, sessionId int, manual bool) (bool, error) {
	if session.Guide.Step+1 >= len(guide.Steps) {
		ended, err := guideEnded(guide, session, sessionId, manual)
		if err != nil {
			return false, err
		}
		return ended, nil
	}
	session.Guide.Step++
	guideStep := guide.Steps[session.Guide.Step]
	io.PrintVerbose(fmt.Sprintf("Executing next step %d\n", session.Guide.Step))
	_, err := io.PrintGuideAndGuideStepTitle(guide, guideStep, session)
	if err != nil {
		return false, err
	}
	for _, input := range guideStep.Inputs {
		io.PrintVerbose(fmt.Sprintf("Showing input for %s\n", input.Name))
		err := io.ScanGuideStepInput(session, input)
		if err != nil {
			return false, err
		}
	}
	_, err = io.PrintGuideStepText(guideStep, session)
	if err != nil {
		return false, err
	}
	return false, nil
}

func guideEnded(guide *model.Guide, guideState *model.Session, sessionId int, manual bool) (bool, error) {
	_, err := io.PrintGuideEndText(guide, sessionId)
	if err != nil {
		return false, err
	}
	io.PrintVerbose(fmt.Sprintf("Reading guide file %s\n", guideState.Guide.File))
	guides, err := io.ReadGuideFile(guideState.Guide.File)
	if err != nil {
		return false, err
	}
	index := guides.IndexByName(guideState.Guide.Name)
	last := index == len(guides)-1

	io.PrintVerbose("Showing decision for continuing\n")
	continueAnother, err := io.ScanBoolDefault("Do you want to continue with another guide", !last || manual)
	if err != nil {
		return false, err
	}
	if !continueAnother {
		io.PrintVerbose("Session will not continue\n")
		return true, nil
	}
	io.PrintVerbose("Showing selection for guides as the session will continue\n")
	var newGuide *model.Guide
	if last {
		newGuide, err = io.ScanSelectWithZero("Select from the list", "Exit", guides, func(g *model.Guide) string {
			return g.Name
		})
	} else {
		newGuide, err = io.ScanSelectWithZeroDefault("Select from the list", "Exit", index+1, guides, func(g *model.Guide) string {
			return g.Name
		})
	}
	if err != nil {
		return false, err
	}
	if newGuide == nil {
		io.PrintVerbose("Session will not continue\n")
		return true, nil
	}
	guideState.Guide.Name = newGuide.Name
	guideState.Guide.Step = -1
	io.PrintVerbose(fmt.Sprintf("Session will continue with %s\n", guideState.Guide.Name))
	return executeGuide(newGuide, guideState, sessionId, false)
}
