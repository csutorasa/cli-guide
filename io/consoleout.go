package io

import (
	"fmt"

	"github.com/csutorasa/cli-guide/model"
)

func PrintGuideAndGuideStepTitle(guide *model.Guide, guideStep *model.GuideStep, session *model.Session) (int, error) {
	written := 0
	guideName, err := WithTemplate(guide.Name, session.Variables)
	if err != nil {
		return written, err
	}
	guideStepTitle, err := WithTemplate(guideStep.Title, session.Variables)
	if err != nil {
		return written, err
	}
	i, err := fmt.Printf("##########\n")
	written += i
	if err != nil {
		return written, err
	}
	i, err = fmt.Printf("# %s - %s\n", guideName, guideStepTitle)
	written += i
	if err != nil {
		return written, err
	}
	i, err = fmt.Printf("##########\n")
	written += i
	if err != nil {
		return written, err
	}
	return written, nil
}

func PrintGuideStepText(guideStep *model.GuideStep, session *model.Session) (int, error) {
	guideStepText, err := WithTemplate(guideStep.Text, session.Variables)
	if err != nil {
		return 0, err
	}
	if guideStepText != "" {
		if guideStepText[len(guideStepText)-1] == '\n' {
			return fmt.Printf("%s\n", guideStepText)
		} else {
			return fmt.Printf("%s", guideStepText)
		}
	}
	return 0, nil
}

func PrintGuideEndText(guide *model.Guide) (int, error) {
	return fmt.Printf("%s has finished\n", guide.Name)
}
