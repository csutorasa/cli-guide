package model

type GuideFile []*Guide

type Guide struct {
	Name  string       `yaml:"name"`
	Steps []*GuideStep `yaml:"steps"`
}

type GuideStep struct {
	Title  string            `yaml:"title"`
	Inputs []*GuideStepInput `yaml:"inputs"`
	Text   string            `yaml:"text"`
}

type GuideStepInput struct {
	Name      string `yaml:"name"`
	Text      string `yaml:"text"`
	Example   string `yaml:"example"`
	Validator string `yaml:"validator"`
}

func (g GuideFile) IndexByName(name string) int {
	for i, guide := range g {
		if guide.Name == name {
			return i
		}
	}
	return -1
}

func (g GuideFile) FindGuideByName(name string) *Guide {
	for _, guide := range g {
		if guide.Name == name {
			return guide
		}
	}
	return nil
}
