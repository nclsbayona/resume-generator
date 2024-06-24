package ports

import (
	"github.com/nclsbayona/resume-generator/pkg/core/domain"
)

type Generator interface {
	GetFull(*domain.Logger) *string
	SetResume(*domain.FullResume)
	SetOptions(...*string) // This is specific to each adapter
}
