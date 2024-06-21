package domain

type Education struct {
	Title       *string
	Institution *string
	Location    *string
	StartDate   *string
	EndDate     *string
	Description *string
}

func NewEducation(title *string, institution *string, location *string, startDate *string, endDate *string, description *string) *Education {
	return &Education{
		Title:       title,
		Institution: institution,
		Location:    location,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
	}
}
