package io

import (
	"fmt"
	"os"

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
	i, err := Print(fmt.Sprintf("%s - %s\n", guideName, guideStepTitle))
	written += i
	if err != nil {
		return written, err
	}
	i, err = Print("##########\n")
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
			return PrintQuiet(guideStepText)
		} else {
			return PrintQuiet(guideStepText + "\n")
		}
	}
	return 0, nil
}

func PrintGuideEndText(guide *model.Guide, sessionId int) (int, error) {
	return Print(fmt.Sprintf("%s (session %d) has finished\n", guide.Name, sessionId))
}

func PrintVerbose(s string) (int, error) {
	if logLevel <= LogLevelVerbose {
		return fmt.Print(s)
	}
	return 0, nil
}

func Print(s string) (int, error) {
	if logLevel <= LogLevelNormal {
		return fmt.Print(s)
	}
	return 0, nil
}

func PrintQuiet(s string) (int, error) {
	if logLevel <= LogLevelQuiet {
		return fmt.Print(s)
	}
	return 0, nil
}

func PrintFatalError(err error) {
	fmt.Fprintf(os.Stderr, "FATAL ERROR: %s", err.Error())
	os.Exit(1)
}

func PrintFatalError2(err error, extra string) {
	fmt.Fprintf(os.Stderr, "FATAL ERROR: %s\n%s", err.Error(), extra)
	os.Exit(1)
}
