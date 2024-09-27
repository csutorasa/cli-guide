package io

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/csutorasa/cli-guide/model"
	"gopkg.in/yaml.v3"
)

var stateFilePath = filepath.Join(rootDir, "state.yml")

func ReadState() (*model.State, error) {
	f, err := os.Open(stateFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open state file: %w\n%s", err, stateFilePath)
	}
	defer f.Close()
	d := yaml.NewDecoder(f)
	state := &model.State{}
	err = d.Decode(&state)
	if err != nil {
		return nil, fmt.Errorf("could not decode yml state file: %w\n%s", err, stateFilePath)
	}
	return state, nil
}

func WriteState(state *model.State) error {
	f, err := os.OpenFile(stateFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("could not open state file: %w\n%s", err, stateFilePath)
	}
	defer f.Close()
	e := yaml.NewEncoder(f)
	err = e.Encode(state)
	if err != nil {
		return fmt.Errorf("could not decode yml state file: %w\n%s", err, stateFilePath)
	}
	return nil
}
