package engine

import (
	// "github.com/nclsbayona/resume-generator/engine/interfaces"
)

type FullResume struct {
	Name	   string `yaml:"name"`;
	Experiences []SimpleExperience `yaml:"experiences"`
}

func (f FullResume) Render() string {

	rendered := f.Name + "\n\n"
	for _, experience := range f.Experiences {
		rendered += experience.Render() + "\n"
	}
	return rendered
	
}