package domain

import (
	"reflect"
	"testing"
)

func TestNewExtraInformation(t *testing.T) {
	type args struct {
		label *string
		value *string
	}

	type test struct {
		name *string
		args *args
		want *ExtraInformation
	}

	name := "TestNewExtraInformation"
	label := "label"
	value := "value"

	arguments := &args{
		label: &label,
		value: &value,
	}

	want := &ExtraInformation{
		Label: &label,
		Value: &value,
	}

	testCase := &test{
		name: &name,
		args: arguments,
		want: want,
	}

	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewExtraInformation(testCase.args.label, testCase.args.value); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewExtraInformation() = %v, want %v", got, testCase.want)
		}
	})

	name = "Test Nil ExtraInformation"
	arguments = &args{}
	want = &ExtraInformation{}
	t.Run(*testCase.name, func(t *testing.T) {
		if got := NewExtraInformation(testCase.args.label, testCase.args.value); !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("NewExtraInformation() = %v, want %v", got, testCase.want)
		}
	})
}
