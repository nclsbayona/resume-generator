package injector

// Options: input_reading (File to read the resume from) 

import (
	"github.com/nclsbayona/resume-generator/pkg/core/domain"
	"github.com/nclsbayona/resume-generator/pkg/core/ports"
	"gopkg.in/yaml.v3"
	"os"
)

type sAchievement string

type sExperience struct {
	Title        *string         `yaml:"title"`
	Company      *string         `yaml:"company"`
	Location     *string         `yaml:"location"`
	StartDate    *string         `yaml:"startDate"`
	EndDate      *string         `yaml:"endDate"`
	Description  *string         `yaml:"description"`
	Achievements []*sAchievement `yaml:"achievements"`
}

type sEducation struct {
	Title       *string `yaml:"title"`
	Institution *string `yaml:"institution"`
	Location    *string `yaml:"location"`
	StartDate   *string `yaml:"startDate"`
	EndDate     *string `yaml:"endDate"`
	Description *string `yaml:"description"`
}

type sFullResumeYaml struct {
	input_reading *string
	Name          *string        `yaml:"full_name"`
	Experiences   []*sExperience `yaml:"experiences"`
	Education     []*sEducation  `yaml:"education"`
}

func (input *sFullResumeYaml) SetOptions(options ...*string) {
	input.input_reading = options[0]
}

func NewYaml() *ports.Injector {
	var full_resume ports.Injector = &sFullResumeYaml{}
	return &full_resume
}

func (input *sFullResumeYaml) GetFull(logger *domain.Logger) *domain.FullResume {
	var out []byte
	out, err := os.ReadFile(*input.input_reading)
	if err != nil {
		logger.Error("Error reading file: " + *input.input_reading)
	}
	logger.Info("File " + *input.input_reading + " read successfully")
	logger.Info("Unmarshalling yaml\n" + string(out))
	err = yaml.Unmarshal(out, &input)
	if err != nil {
		logger.Error("Error unmarshalling yaml: " + err.Error())
		return nil
	}
	logger.Info("Starting resume generation for \"" + *input.Name + "\"")

	logger.Info("Education: ")
	for _, e := range input.Education {
		logger.Info("Title: " + *e.Title + " Institution: " + *e.Institution + " Location: " + *e.Location + " Start Date: " + *e.StartDate + " End Date: " + *e.EndDate + " Description: " + *e.Description)
	}
	experience := convertExperience(input.Experiences)

	education := convertEducation(input.Education)
	var tmp_fullresume *domain.FullResume = domain.NewFullResume(input.Name, experience, education)
	return tmp_fullresume
	//usecase.NewPrintResume(tmp_fullresume)
	//print_usecase.SetGenerator(input.output)
	//return print_usecase.GetFull(logger)
}

func convertExperience(experience []*sExperience) []*domain.Experience {
	var experiences []*domain.Experience = make([]*domain.Experience, 0)
	for _, e := range experience {
		experiences = append(experiences, domain.NewExperience(e.Title, e.Company, e.Location, e.StartDate, e.EndDate, e.Description, convertAchievement(e.Achievements)))
	}
	return experiences
}

func convertEducation(education []*sEducation) []*domain.Education {
	var educations []*domain.Education = make([]*domain.Education, 0)
	for _, e := range education {
		educations = append(educations, domain.NewEducation(e.Title, e.Institution, e.Location, e.StartDate, e.EndDate, e.Description))
	}
	return educations
}

func convertAchievement(achievement []*sAchievement) []*string {
	var achievements []*string = make([]*string, 0)
	for _, a := range achievement {
		achievement := string(*a)
		achievements = append(achievements, &achievement)
	}
	return achievements
}
