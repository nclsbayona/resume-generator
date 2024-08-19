package ports

import (
	"github.com/nclsbayona/resume-generator/domain"
)

type Commander interface {
	RunCommand(*domain.Logger) *string
	SetGenerator(*Generator)
	SetInjector(*Injector)
	SetOptions(...*string) // This is specific to each adapter
}
