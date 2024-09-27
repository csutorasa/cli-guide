package io

import (
	"fmt"
	"os"

	"github.com/csutorasa/cli-guide/model"
	"gopkg.in/yaml.v3"
)

func ReadGuideFile(p string) (model.GuideFile, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("could not open guide file: %w\n%s", err, p)
	}
	defer f.Close()
	d := yaml.NewDecoder(f)
	guideFile := model.GuideFile{}
	err = d.Decode(&guideFile)
	if err != nil {
		return nil, fmt.Errorf("could not decode yml guide file: %w\n%s", err, p)
	}
	return guideFile, nil
}
