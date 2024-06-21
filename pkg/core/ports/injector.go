package ports

import (
	"github.com/nclsbayona/resume-generator/pkg/core/domain"
)

type Injector interface {
	GetFull(*domain.Logger) *domain.FullResume
	SetOptions(...*string) // This is specific to each adapter
}
