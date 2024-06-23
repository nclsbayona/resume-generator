package domain

type sAchievement *string

type Experience struct {
	Title        *string
	Company      *string
	Location     *string
	StartDate    *string
	EndDate      *string
	Description  *string
	Achievements []*sAchievement
}

func NewExperience(title *string, company *string, location *string, startDate *string, endDate *string, description *string, achievements []*string) *Experience {
	// Convert achievements slice from []*string to []*sAchievement
	var sAchievements []*sAchievement
	if achievements == nil {
		sAchievements = nil
	} else {
		for _, a := range achievements {
			s := sAchievement(a)
			sAchievements = append(sAchievements, &s)
		}
	}

	return &Experience{
		Title:        title,
		Company:      company,
		Location:     location,
		StartDate:    startDate,
		EndDate:      endDate,
		Description:  description,
		Achievements: sAchievements,
	}
}
