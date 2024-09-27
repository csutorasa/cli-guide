package model

type Session struct {
	Guide     *SessionGuide  `yaml:"guide"`
	Variables map[string]any `yaml:"variables"`
}

type SessionGuide struct {
	File string `yaml:"file"`
	Name string `yaml:"name"`
	Step int    `yaml:"step"`
}
