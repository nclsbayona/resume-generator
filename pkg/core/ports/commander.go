package ports

import (
	"github.com/nclsbayona/resume-generator/pkg/core/domain"
)

type Commander interface {
	RunCommand(*domain.Logger) *string
	SetOptions(...*string) // This is specific to each adapter
	SetGenerator(*Generator)
	SetInjector(*Injector)
}