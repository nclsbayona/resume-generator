package usecase

import (
	"github.com/nclsbayona/resume-generator/domain"
	"github.com/nclsbayona/resume-generator/ports"
)

type sPrintResume struct {
	resume    *domain.FullResume
	generator *ports.Generator
	injector  *ports.Injector
}

func (p *sPrintResume) RunCommand(logger *domain.Logger) (rendered *string) {
	logger.Info("Injecting resume to print ")
	p.resume = (*p.injector).GetFull(logger)
	logger.Info("Calling output to generate resume ")
	(*p.generator).SetResume(p.resume)
	logger.Info("Experiences: ")
	for _, e := range p.resume.Experiences {
		logger.Info("Title: " + *e.Title + " Company: " + *e.Company + " Location: " + *e.Location + " Start Date: " + *e.StartDate + " End Date: " + *e.EndDate + " Description: " + *e.Description)
		logger.Info("Achievements: ")
		for _, a := range e.Achievements {
			logger.Info("Achievement: " + string(**a))
		}
	}
	rendered = (*p.generator).GetFull(logger)
	logger.Info("Output generated resume successfully")
	return
}

func (p *sPrintResume) SetInjector(injector *ports.Injector) {
	p.injector = injector
}

func NewPrintResume(resume *domain.FullResume, generator *ports.Generator, injector *ports.Injector) ports.Commander {
	return &sPrintResume{
		resume:    resume,
		generator: generator,
		injector:  injector,
	}
}

func (p *sPrintResume) SetOptions(_ ...*string) {
}

func (p *sPrintResume) SetGenerator(output *ports.Generator) {
	p.generator = output
}
