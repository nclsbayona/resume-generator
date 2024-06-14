package engine

import (
	"fmt"
)

type SimpleExperience struct {
	Title        string   `yaml:"title"`
	Company      string   `yaml:"company"`
	Location     string   `yaml:"location"`
	StartDate    string   `yaml:"start_date"`
	EndDate      string   `yaml:"end_date"`
	Description  string   `yaml:"description"`
	Achievements []string `yaml:"achievements"`
}

func (e SimpleExperience) Render() string {
	achievements := ""
	for _, achievement := range e.Achievements {
		achievements += "- " + achievement + "\n"
	}
	return fmt.Sprintf("%s (%s, %s, %s - %s)\n%s\n%s", e.Title, e.Company, e.Location, e.StartDate, e.EndDate, e.Description, achievements)
}
