package engine

import (
	"github.com/cbroglie/mustache"
)

type Achievement string

type Experience struct {
	Title        string        `yaml:"title"`
	Company      string        `yaml:"company"`
	Location     string        `yaml:"location"`
	StartDate    string        `yaml:"start_date"`
	EndDate      string        `yaml:"end_date"`
	Description  string        `yaml:"description"`
	Achievements []Achievement `yaml:"achievements"`
}

func (e Experience) Render(experienceTemplate string, achievementsTemplate string) *string {
	extraInfo := make(map[string]interface{})
	extraInfo["title"] = e.Title
	extraInfo["company"] = e.Company
	extraInfo["location"] = e.Location
	extraInfo["start_date"] = e.StartDate
	extraInfo["end_date"] = e.EndDate
	extraInfo["description"] = e.Description
	if achievementsTemplate != "" {
		template, err := mustache.ParseFile(achievementsTemplate)
		if err != nil {
			panic(err)
		}
		achievements, err := template.Render(e.Achievements)
		if err != nil {
			panic(err)
		}
		extraInfo["achievements"] = achievements
	}
	template, err := mustache.ParseFile(experienceTemplate)
	if err != nil {
		panic(err)
	}
	rendered, err := template.Render(extraInfo)
	if err != nil {
		panic(err)
	}
	return &rendered
}
