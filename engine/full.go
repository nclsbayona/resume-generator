package engine

import (
	"log"
	"github.com/cbroglie/mustache"
)

type FullResume struct {
	Templates   map[string]string `yaml:"templates"`
	Name        string            `yaml:"name"`
	Experiences []Experience      `yaml:"experiences"`
}

type experienceString string

func newExperienceString(val string) *experienceString {
	return (*experienceString)(&val)
}

func (f FullResume) Render() string {
	extraInfo := make(map[string]interface{})
	extraInfo["name"] = f.Name
	if f.Templates["experience"] != "" {
		if f.Templates["experience"] == "default" {
			log.Println("Using default template file for experience!")
			f.Templates["experience"] = "templates/experience.html.mustache"
		} else {
			log.Println("Using custom template file for experience!")
		}
		if f.Templates["achievements"] == "default" {
			log.Println("Using default template file for achievements!")
			f.Templates["achievements"] = "templates/achievements.html.mustache"
		} else {
			log.Println("No provided template file for achievements!")
		}
		experiences := make([]experienceString, len(f.Experiences))
		for i, e := range f.Experiences {
			experiences[i] = *newExperienceString(*e.Render(f.Templates["experience"], f.Templates["achievements"]))
		}
		extraInfo["experiences"] = experiences
	} else {
		log.Println("No experience template file provided!")
	}

	templateFile := f.Templates["full"]
	if templateFile == "" {
		log.Println("No template file for the full YAML generation provided. Using default one!")
		templateFile = "templates/full.html.mustache"
	}
	template, err := mustache.ParseFile(templateFile)
	if err != nil {
		panic(err)
	}
	rendered, err := template.Render(extraInfo)
	if err != nil {
		panic(err)\
	}
	return rendered
}
