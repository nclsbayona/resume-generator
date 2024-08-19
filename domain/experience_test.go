package domain

import (
	"reflect"
	"testing"
)

func stringAchievement(s *string) *sAchievement {
	achievement := sAchievement(s)
	return &achievement
}
func TestNewExperience(t *testing.T) {

	str := "Test"
	achievement := stringAchievement(&str)

	type args struct {
		title        *string
		company      *string
		location     *string
		startDate    *string
		endDate      *string
		description  *string
		achievements []*string
	}

	type test struct {
		name      string
		arguments *args
		want      *Experience
	}

	var (
		name         = "Test Full NewExperience"
		title        = "Test Title"
		company      = "Test Company"
		location     = "Test Location"
		startDate    = "Test Start Date"
		endDate      = "Test End Date"
		description  = "Test Description"
		achievements = []*string{&str}
		arguments    = &args{
			title:        &title,
			company:      &company,
			location:     &location,
			startDate:    &startDate,
			endDate:      &endDate,
			description:  &description,
			achievements: achievements,
		}
		want = &Experience{
			Title:        arguments.title,
			Company:      arguments.company,
			Location:     arguments.location,
			StartDate:    arguments.startDate,
			EndDate:      arguments.endDate,
			Description:  arguments.description,
			Achievements: []*sAchievement{achievement},
		}
	)

	testCase := test{name: name, arguments: arguments, want: want}
	t.Run(testCase.name, func(t *testing.T) {
		if got := NewExperience(testCase.arguments.title, testCase.arguments.company, testCase.arguments.location, testCase.arguments.startDate, testCase.arguments.endDate, testCase.arguments.description, testCase.arguments.achievements); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewExperience() = %v, want %v", got, testCase.want)
		}
	})

	name = "Test Empty NewExperience"
	title = ""
	company = ""
	location = ""
	startDate = ""
	endDate = ""
	description = ""
	achievements = nil
	t.Run(testCase.name, func(t *testing.T) {
		if got := NewExperience(testCase.arguments.title, testCase.arguments.company, testCase.arguments.location, testCase.arguments.startDate, testCase.arguments.endDate, testCase.arguments.description, testCase.arguments.achievements); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewExperience() = %v, want %v", got, testCase.want)
		}
	})

	name = "Test Nil NewExperience"
	arguments = &args{}
	want = &Experience{}
	t.Run(testCase.name, func(t *testing.T) {
		if got := NewExperience(testCase.arguments.title, testCase.arguments.company, testCase.arguments.location, testCase.arguments.startDate, testCase.arguments.endDate, testCase.arguments.description, testCase.arguments.achievements); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewExperience() = %v, want %v", got, testCase.want)
		}
	})
}
