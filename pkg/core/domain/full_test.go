package domain

import (
	"reflect"
	"testing"
)

func TestNewFullResume(t *testing.T) {
	type args struct {
		summary            *string
		name               *string
		experiences        []*Experience
		education          []*Education
		extraInfo          []*ExtraInformation
		continousEducation []*ContinousEducation
	}

	type test struct {
		name *string
		args *args
		want *FullResume
	}

	var (
		name      = "Test Full NewFullResume"
		arguments = args{
			summary: &name,
			name:    &name,
			experiences: []*Experience{
				{
					Title:        &name,
					Company:      &name,
					Location:     &name,
					StartDate:    &name,
					EndDate:      &name,
					Description:  &name,
					Achievements: []*sAchievement{stringAchievement(&name)},
				},
			},
			education: []*Education{
				{
					Title:       &name,
					Institution: &name,
					Location:    &name,
					StartDate:   &name,
					EndDate:     &name,
					Description: &name,
				},
			},
			extraInfo: []*ExtraInformation{
				{
					Label: &name,
					Value: &name,
				},
			},
			continousEducation: []*ContinousEducation{
				{
					Title:       &name,
					Institution: &name,
					Location:    &name,
					Summary:     &name,
					Date:        &name,
					Expire:      &name,
				},
			},
		}
		want = FullResume{
			Summary:            arguments.summary,
			Name:               arguments.name,
			Experiences:        arguments.experiences,
			Education:          arguments.education,
			ExtraInfo:          arguments.extraInfo,
			ContinousEducation: arguments.continousEducation,
		}
	)

	testCase := test{
		name: &name,
		args: &arguments,
		want: &want,
	}

	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewFullResume(testCase.args.summary, testCase.args.name, testCase.args.experiences, testCase.args.education, testCase.args.extraInfo, testCase.args.continousEducation); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewFullResume() = %v, want %v", got, testCase.want)
		}
	})

	name = "Test Empty NewFullResume"
	arguments = args{}
	want = FullResume{}

	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewFullResume(testCase.args.summary, testCase.args.name, testCase.args.experiences, testCase.args.education, testCase.args.extraInfo, testCase.args.continousEducation); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewFullResume() = %v, want %v", got, testCase.want)
		}
	})
}
