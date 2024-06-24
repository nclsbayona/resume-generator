package domain

import (
	"reflect"
	"testing"
)

func TestNewContinousEducation(t *testing.T) {
	type args struct {
		title       *string
		institution *string
		location    *string
		summary     *string
		date        *string
		expire      *string
	}

	type test struct {
		name *string
		args *args
		want *ContinousEducation
	}

	title := "TestNewContinousEducation"
	institution := "institution"
	location := "location"
	summary := "summary"
	date := "date"
	expire := "expire"

	arguments := &args{
		title:       &title,
		institution: &institution,
		location:    &location,
		summary:     &summary,
		date:        &date,
		expire:      &expire,
	}

	want := &ContinousEducation{
		Title:       &title,
		Institution: &institution,
		Location:    &location,
		Summary:     &summary,
		Date:        &date,
		Expire:      &expire,
	}

	testCase := &test{
		name: &title,
		args: arguments,
		want: want,
	}

	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewContinousEducation(testCase.args.title, testCase.args.institution, testCase.args.location, testCase.args.summary, testCase.args.date, testCase.args.expire); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewContinousEducation() = %v, want %v", got, testCase.want)
		}
	})

	title = "Test Nil ContinousEducation"
	arguments = &args{}
	want = &ContinousEducation{}

	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewContinousEducation(testCase.args.title, testCase.args.institution, testCase.args.location, testCase.args.summary, testCase.args.date, testCase.args.expire); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewContinousEducation() = %v, want %v", got, testCase.want)
		}
	})

}
