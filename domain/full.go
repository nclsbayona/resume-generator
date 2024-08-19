package domain

type FullResume struct {
	Summary            *string
	Name               *string
	Experiences        []*Experience
	Education          []*Education
	ExtraInfo          []*ExtraInformation
	ContinousEducation []*ContinousEducation
}

func NewFullResume(summary *string, name *string, experiences []*Experience, education []*Education, extraInformation []*ExtraInformation, continousEducation []*ContinousEducation) *FullResume {
	return &FullResume{
		Summary:            summary,
		Name:               name,
		Experiences:        experiences,
		Education:          education,
		ExtraInfo:          extraInformation,
		ContinousEducation: continousEducation,
	}
}
