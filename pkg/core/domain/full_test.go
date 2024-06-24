package domain

import (
	"reflect"
	"testing"
)

func TestNewFullResume(t *testing.T) {
	type args struct {
		name        *string
		experiences []*Experience
		education   []*Education
	}
	type test struct {
		name *string
		args *args
		want *FullResume
	}

	var (
		name      = "Test Full NewFullResume"
		arguments = args{
			name:     &name,
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
		}
		want = FullResume{
			Name:        arguments.name,
			Experiences: arguments.experiences,
			Education:   arguments.education,
		}
	)
	testCase := test{
		name: &name,
		args: &arguments,
		want: &want,
	}
	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewFullResume(testCase.args.name,  testCase.args.experiences, testCase.args.education); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewFullResume() = %v, want %v", got, testCase.want)
		}
	})
	name = "Test Empty NewFullResume"
	arguments = args{}
	want = FullResume{}
	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewFullResume(testCase.args.name, testCase.args.experiences, testCase.args.education); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewFullResume() = %v, want %v", got, testCase.want)
		}
	})
}
