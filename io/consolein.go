package io

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/csutorasa/cli-guide/model"
)

var StdinScanner = bufio.NewScanner(os.Stdin)

func ScanGuideStepInput(session *model.Session, guideStepInput *model.GuideStepInput) error {
	var validator *regexp.Regexp
	if guideStepInput.Validator != "" {
		var err error
		validator, err = regexp.Compile(guideStepInput.Validator)
		if err != nil {
			return fmt.Errorf("validator contains invalid regexp: %w\n%s", err, guideStepInput.Validator)
		}
	}
	for {
		guideStepInputText, err := WithTemplate(guideStepInput.Text, session.Variables)
		if err != nil {
			return err
		}
		_, err = PrintQuiet(guideStepInputText)
		if err != nil {
			return err
		}
		if guideStepInput.Example != "" {
			_, err = PrintQuiet(fmt.Sprintf(" (example: %s)", guideStepInput.Example))
			if err != nil {
				return err
			}
		}
		defaultValue, hasDefaultvalue := session.Variables[guideStepInput.Name]
		if hasDefaultvalue {
			_, err = PrintQuiet(fmt.Sprintf(" [%s]", defaultValue))
			if err != nil {
				return err
			}
		}
		_, err = PrintQuiet(": ")
		if err != nil {
			return err
		}
		if !StdinScanner.Scan() {
			return fmt.Errorf("could not read user input: %w", StdinScanner.Err())
		}
		str := StdinScanner.Text()
		if str == "" && hasDefaultvalue {
			session.Variables[guideStepInput.Name] = defaultValue
			return nil
		}
		if validator == nil || validator.MatchString(str) {
			session.Variables[guideStepInput.Name] = str
			return nil
		}
		PrintQuiet(fmt.Sprintf("%s must match '%s' regexp\n", guideStepInput.Name, guideStepInput.Validator))
	}
}

func ScanBoolDefault(s string, b bool) (bool, error) {
	_, err := PrintQuiet(fmt.Sprintf("%s (", s))
	if err != nil {
		return false, err
	}
	if b {
		_, err = PrintQuiet("Y")
	} else {
		_, err = PrintQuiet("y")
	}
	if err != nil {
		return false, err
	}
	_, err = PrintQuiet("/")
	if err != nil {
		return false, err
	}
	if b {
		_, err = PrintQuiet("n")
	} else {
		_, err = PrintQuiet("N")
	}
	if err != nil {
		return false, err
	}
	_, err = PrintQuiet(")? ")
	if err != nil {
		return false, err
	}
	if !StdinScanner.Scan() {
		return false, fmt.Errorf("could not read user input: %w", StdinScanner.Err())
	}
	str := StdinScanner.Text()
	if str == "Y" || str == "y" || strings.ToLower(str) == "yes" || str == "1" {
		return true, nil
	}
	if str == "N" || str == "n" || strings.ToLower(str) == "no" || str == "0" {
		return false, nil
	}
	return b, nil
}

func ScanSelectWithZero[T any](s string, zeroText string, items []T, f func(T) string) (T, error) {
	var t T
	for {
		for i, item := range items {
			_, err := PrintQuiet(fmt.Sprintf("[%d] %s\n", i+1, f(item)))
			if err != nil {
				return t, err
			}
		}
		_, err := PrintQuiet(fmt.Sprintf("[0] %s\n", zeroText))
		if err != nil {
			return t, err
		}
		_, err = PrintQuiet(fmt.Sprintf("%s: ", s))
		if err != nil {
			return t, err
		}
		if !StdinScanner.Scan() {
			return t, fmt.Errorf("could not read user input: %w", StdinScanner.Err())
		}
		str := StdinScanner.Text()
		selected, err := strconv.ParseInt(str, 10, 32)
		if err != nil || selected < 0 || int(selected) > len(items) {
			PrintQuiet(fmt.Sprintf("Value should be a number between 1 and %d\n", len(items)))
			continue
		}
		if selected == 0 {
			return t, nil
		}
		return items[int(selected)-1], nil
	}
}

func ScanSelectWithZeroDefault[T any](s string, zeroText string, def int, items []T, f func(T) string) (T, error) {
	var t T
	for {
		for i, item := range items {
			_, err := PrintQuiet(fmt.Sprintf("[%d] %s\n", i+1, f(item)))
			if err != nil {
				return t, err
			}
		}
		_, err := PrintQuiet(fmt.Sprintf("[0] %s\n", zeroText))
		if err != nil {
			return t, err
		}
		_, err = PrintQuiet(fmt.Sprintf("%s [%d]: ", s, def+1))
		if err != nil {
			return t, err
		}
		if !StdinScanner.Scan() {
			return t, fmt.Errorf("could not read user input: %w", StdinScanner.Err())
		}
		str := StdinScanner.Text()
		if str == "" {
			return items[def], nil
		}
		selected, err := strconv.ParseInt(str, 10, 32)
		if err != nil || selected < 0 || int(selected) > len(items) {
			PrintQuiet(fmt.Sprintf("Value should be a number between 1 and %d\n", len(items)))
			continue
		}
		if selected == 0 {
			return t, nil
		}
		return items[int(selected)-1], nil
	}
}

func ScanSelectMapWithZero[V any](s string, zeroText string, items map[int]V, f func(int, V) string) (int, V, error) {
	var v V
	for {
		for k, item := range items {
			_, err := PrintQuiet(fmt.Sprintf("[%d] %s\n", k, f(k, item)))
			if err != nil {
				return 0, v, err
			}
		}
		_, err := PrintQuiet(fmt.Sprintf("[0] %s\n", zeroText))
		if err != nil {
			return 0, v, err
		}
		_, err = PrintQuiet(fmt.Sprintf("%s: ", s))
		if err != nil {
			return 0, v, err
		}
		if !StdinScanner.Scan() {
			return 0, v, fmt.Errorf("could not read user input: %w", StdinScanner.Err())
		}
		str := StdinScanner.Text()
		selected, err := strconv.ParseInt(str, 10, 32)
		if err != nil || selected < 0 || int(selected) > len(items) {
			PrintQuiet("Value should be a number\n")
			continue
		}
		if selected == 0 {
			return 0, v, nil
		}
		item, exists := items[int(selected)]
		if exists {
			return int(selected), item, nil
		}
		PrintQuiet("Selected number is not valid\n")
	}
}
