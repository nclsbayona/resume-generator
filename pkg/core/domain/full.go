package domain

type FullResume struct {
	Template    *string
	Name        *string
	Experiences []*Experience
	Education   []*Education
}

func NewFullResume(name *string, template *string, experiences []*Experience, education []*Education) *FullResume {
	return &FullResume{
		Template:    template,
		Name:        name,
		Experiences: experiences,
		Education:   education,
	}
}
