package io

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"text/template"
)

var rootDir string

type LogLevel int

const (
	LogLevelVerbose LogLevel = iota
	LogLevelNormal
	LogLevelQuiet
)

var logLevel LogLevel

func (l LogLevel) String() string {
	switch l {
	case LogLevelVerbose:
		return "verbose"
	case LogLevelNormal:
		return "normal"
	case LogLevelQuiet:
		return "quiet"
	default:
		panic("unknown LogLevel")
	}
}

func SetLogLevel(l LogLevel) {
	logLevel = l
}

func SetRootDir(s string) error {
	stat, err := os.Stat(s)
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("directory does not exist %s", s)
	}
	if err != nil {
		return err
	}
	if !stat.IsDir() {
		return fmt.Errorf("not a directory %s", s)
	}
	rootDir = s
	return nil
}

func WithTemplate(s string, data map[string]any) (string, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w\n%s", err, s)
	}
	var b = bytes.NewBuffer(make([]byte, 16*1024))
	err = t.Execute(b, data)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w\n%s", err, s)
	}
	return b.String(), nil
}
