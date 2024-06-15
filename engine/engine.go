package engine

import (
	"gopkg.in/yaml.v3"
	"os"
)

func start() (full *FullResume) {

	full = &FullResume{}
	conf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(conf, full)
	if err != nil {
		panic(err)
	}
	return
}

func BuildResume() (rendered string) {
	if len(os.Args) < 2 {
		panic("No configuration file provided!")
	}
	full := start()
	rendered = full.Render()
	return rendered
}
