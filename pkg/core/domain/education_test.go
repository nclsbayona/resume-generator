package domain

import (
	"reflect"
	"testing"
)

func TestNewEducation(t *testing.T) {

	type args struct {
		title       *string
		institution *string
		location    *string
		startDate   *string
		endDate     *string
		description *string
	}

	type test struct {
		name      string
		arguments args
		want      *Education
	}

	var (
		name        = "Test Full NewEducation"
		title       = "Test Title"
		institution = "Test Institution"
		location    = "Test Location"
		startDate   = "Test Start Date"
		endDate     = "Test End Date"
		description = "Test Description"
		arguments   = args{
			title:       &title,
			institution: &institution,
			location:    &location,
			startDate:   &startDate,
			endDate:     &endDate,
			description: &description,
		}
		want = &Education{
			Title:       arguments.title,
			Institution: arguments.institution,
			Location:    arguments.location,
			StartDate:   arguments.startDate,
			EndDate:     arguments.endDate,
			Description: arguments.description,
		}
	)

	testCase := test{name: name, arguments: arguments, want: want}
	t.Run(testCase.name, func(t *testing.T) {
		if got := NewEducation(testCase.arguments.title, testCase.arguments.institution, testCase.arguments.location, testCase.arguments.startDate, testCase.arguments.endDate, testCase.arguments.description); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewEducation() = %v, want %v", got, testCase.want)
		}
	})

	name = "Test Empty NewEducation"
	title = ""
	institution = ""
	location = ""
	startDate = ""
	endDate = ""
	description = ""
	t.Run(testCase.name, func(t *testing.T) {
		if got := NewEducation(testCase.arguments.title, testCase.arguments.institution, testCase.arguments.location, testCase.arguments.startDate, testCase.arguments.endDate, testCase.arguments.description); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewEducation() = %v, want %v", got, testCase.want)
		}
	})

	name = "Test Nil NewEducation"
	arguments = args{}
	want = &Education{}
	t.Run(testCase.name, func(t *testing.T) {
		if got := NewEducation(testCase.arguments.title, testCase.arguments.institution, testCase.arguments.location, testCase.arguments.startDate, testCase.arguments.endDate, testCase.arguments.description); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewEducation() = %v, want %v", got, testCase.want)
		}
	})

}
