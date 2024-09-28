package internal

import (
	"fmt"
	"path/filepath"

	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

const createUsage = "Usage: cli-guide [-rootDir rootDir] [-q | -v] create guideFilePath"

func CreateArgs(args []string) {
	if len(args) == 0 {
		io.PrintFatalError2(fmt.Errorf("guideFilePath is expected"), createUsage)
	}
	if len(args) > 1 {
		io.PrintFatalError2(fmt.Errorf("too many arguments"), createUsage)
	}
	create(args[0])
}

func create(guidePath string) {
	io.PrintVerbose(fmt.Sprintf("Reading guide file: %s\n", guidePath))
	guides, err := io.ReadGuideFile(guidePath)
	if err != nil {
		io.PrintFatalError(err)
	}
	if len(guides) == 0 {
		io.PrintFatalError(fmt.Errorf("could not find any guides\n%s", guidePath))
	}
	io.PrintVerbose(fmt.Sprintf("%d guides were found\nShowing selection for creating a session\n", len(guides)))
	guide, err := io.ScanSelectWithZeroDefault("Select guide from the file", "Cancel", 0, guides, func(g *model.Guide) string {
		return g.Name
	})
	if err != nil {
		io.PrintFatalError(err)
	}
	if guide == nil {
		io.PrintVerbose("No guide were selected\nExiting\n")
		return
	}
	io.PrintVerbose(fmt.Sprintf("%s guide was selected\nCreating session\n", guide.Name))
	file := guidePath
	file, _ = filepath.Abs(file)
	session := &model.Session{
		Guide: &model.SessionGuide{
			File: file,
			Step: -1,
			Name: guide.Name,
		},
		Variables: map[string]any{},
	}
	id, err := io.CreateSession()
	if err != nil {
		io.PrintFatalError(err)
	}
	err = io.WriteSession(id, session)
	if err != nil {
		io.PrintFatalError(err)
	}
	io.Print(fmt.Sprintf("Session %d was created\n", id))
	io.PrintVerbose("Exiting\n")
}
