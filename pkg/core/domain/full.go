package domain

type FullResume struct {
	Name        *string
	Experiences []*Experience
	Education   []*Education
}

func NewFullResume(name *string, experiences []*Experience, education []*Education) *FullResume {
	return &FullResume{
		Name:        name,
		Experiences: experiences,
		Education:   education,
	}
}
