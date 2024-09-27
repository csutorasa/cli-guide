package internal

import (
	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

func executeGuideWriteSessionAndState(guide *model.Guide, session *model.Session, sessionId int, manual bool) error {
	ended, err := executeGuide(guide, session, manual)
	if err != nil {
		return err
	}
	return writeSessionAndState(session, sessionId, ended)
}

func writeSessionAndState(session *model.Session, sessionId int, ended bool) error {
	err := io.WriteSession(sessionId, session)
	if err != nil {
		return err
	}
	newSessionId := sessionId
	if ended {
		newSessionId = 0
		deleteSession, err := io.ScanBoolDefault("Do you want to delete this session", true)
		if err != nil {
			return err
		}
		if deleteSession {
			err = io.DeleteSession(sessionId)
			if err != nil {
				return err
			}
		}
	}
	return io.WriteState(&model.State{
		LastSession: newSessionId,
	})
}

func executeGuide(guide *model.Guide, guideState *model.Session, manual bool) (bool, error) {
	if guideState.Guide.Step+1 >= len(guide.Steps) {
		ended, err := guideEnded(guide, guideState, manual)
		if err != nil {
			return false, err
		}
		return ended, nil
	}
	guideState.Guide.Step++
	guideStep := guide.Steps[guideState.Guide.Step]
	_, err := io.PrintGuideAndGuideStepTitle(guide, guideStep, guideState)
	if err != nil {
		return false, err
	}
	for _, input := range guideStep.Inputs {
		err := io.ScanGuideStepInput(guideState, input)
		if err != nil {
			return false, err
		}
	}
	_, err = io.PrintGuideStepText(guideStep, guideState)
	if err != nil {
		return false, err
	}
	return false, nil
}

func guideEnded(guide *model.Guide, guideState *model.Session, manual bool) (bool, error) {
	_, err := io.PrintGuideEndText(guide)
	if err != nil {
		return false, err
	}
	guides, err := io.ReadGuideFile(guideState.Guide.File)
	if err != nil {
		return false, err
	}
	index := guides.IndexByName(guideState.Guide.Name)
	last := index == len(guides)-1
	continueAnother, err := io.ScanBoolDefault("Do you want to continue with another guide", !last || manual)
	if err != nil {
		return false, err
	}
	if !continueAnother {
		return true, nil
	}
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
		return true, nil
	}
	guideState.Guide.Name = newGuide.Name
	guideState.Guide.Step = -1
	return executeGuide(newGuide, guideState, false)
}
