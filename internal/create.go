package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/csutorasa/cli-guide/io"
	"github.com/csutorasa/cli-guide/model"
)

const createUsage = "Usage: cli-guide create guideFilePath"

func CreateArgs(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "guideFilePath is expected\n%s\n", createUsage)
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "too many arguments\n%s\n", createUsage)
		os.Exit(1)
	}
	create(args[0])
}

func create(guidePath string) {
	guides, err := io.ReadGuideFile(guidePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if len(guides) == 0 {
		fmt.Fprintf(os.Stderr, "could not find any guides\n%s", guidePath)
		os.Exit(1)
	}
	guide, err := io.ScanSelectWithZeroDefault("Select guide from the file", "Cancel", 0, guides, func(g *model.Guide) string {
		return g.Name
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if guide == nil {
		return
	}
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
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = executeGuideWriteSessionAndState(guide, session, id, true)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
