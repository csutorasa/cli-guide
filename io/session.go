package io

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/csutorasa/cli-guide/model"
	"gopkg.in/yaml.v3"
)

var sessionFileNameRegexp = regexp.MustCompile("^(?P<id>\\d{1,4})_session.yml$")

func idToFilePath(i int) string {
	filename := fmt.Sprintf("%d_session.yml", i)
	return filepath.Join(rootDir, filename)
}

func CreateSession() (int, error) {
	for i := 1; i < 1000; i++ {
		p := idToFilePath(i)
		_, err := os.Stat(p)
		if errors.Is(err, os.ErrNotExist) {
			f, err := os.Create(p)
			if err != nil {
				return 0, fmt.Errorf("could not create session file: %w\n%s", err, p)
			}
			defer f.Close()
			return i, nil
		}
		return 0, fmt.Errorf("could not check session file existance: %w\n%s", err, p)
	}
	return 0, fmt.Errorf("1_state.yml to 1000_state.yml files all exist")
}

func ListSessionIds() ([]int, error) {
	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, fmt.Errorf("could not read session directory: %w\n%s", err, rootDir)
	}
	sessionIds := []int{}
	for _, entry := range entries {
		matches := sessionFileNameRegexp.FindStringSubmatch(entry.Name())
		if matches != nil {
			id, err := strconv.ParseInt(matches[1], 10, 32)
			if err != nil {
				return nil, fmt.Errorf("invalid regexp match: %w", err)
			}
			sessionIds = append(sessionIds, int(id))
		}
	}
	return sessionIds, nil
}

func ListSessions() (map[int]*model.Session, error) {
	ids, err := ListSessionIds()
	if err != nil {
		return nil, err
	}
	sessions := map[int]*model.Session{}
	for _, id := range ids {
		session, err := ReadSession(id)
		if err != nil {
			return nil, err
		}
		sessions[id] = session
	}
	return sessions, nil
}

func ReadSession(id int) (*model.Session, error) {
	p := idToFilePath(id)
	f, err := os.Open(p)
	if err != nil {
		return nil, fmt.Errorf("could not open session file: %w\n%s", err, p)
	}
	defer f.Close()
	d := yaml.NewDecoder(f)
	session := &model.Session{}
	err = d.Decode(&session)
	if err != nil {
		return nil, fmt.Errorf("could not decode yml session file: %w\n%s", err, p)
	}
	return session, nil
}

func WriteSession(id int, session *model.Session) error {
	p := idToFilePath(id)
	f, err := os.OpenFile(p, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("could not open session file: %w\n%s", err, p)
	}
	defer f.Close()
	e := yaml.NewEncoder(f)
	err = e.Encode(session)
	if err != nil {
		return fmt.Errorf("could not decode yml session file: %w\n%s", err, p)
	}
	return nil
}

func DeleteSession(id int) error {
	p := idToFilePath(id)
	err := os.Remove(p)
	if err != nil {
		return fmt.Errorf("could not delete session file: %w\n%s", err, p)
	}
	return nil
}
