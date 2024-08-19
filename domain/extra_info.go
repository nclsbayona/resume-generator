package domain

type ExtraInformation struct {
	Label *string `yaml:"label"`
	Value *string `yaml:"value"`
}

func NewExtraInformation(label *string, value *string) *ExtraInformation {
	return &ExtraInformation{
		Label: label,
		Value: value,
	}
}
