package domain

type ContinousEducation struct {
	Title       *string `yaml:"title"`
	Institution *string `yaml:"institution"`
	Location    *string `yaml:"location"`
	Summary     *string `yaml:"summary"`
	Date        *string `yaml:"date"`
	Expire      *string `yaml:"expire"`
}

func NewContinousEducation(title *string, institution *string, location *string, summary *string, date *string, expire *string) *ContinousEducation {
	return &ContinousEducation{
		Title:       title,
		Institution: institution,
		Location:    location,
		Summary:     summary,
		Date:        date,
		Expire:      expire,
	}
}
