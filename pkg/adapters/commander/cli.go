package commander

import (
	"fmt"
	"github.com/nclsbayona/resume-generator/pkg/core/domain"
	"github.com/nclsbayona/resume-generator/pkg/core/ports"
)

type CLI struct {
	prefix *string
	injector *ports.Injector
	generator *ports.Generator
}

func (c *CLI) SetGenerator(generator *ports.Generator) {
	c.generator = generator
}

func (c *CLI) SetInjector(injector *ports.Injector) {
	c.injector = injector
}

func (c *CLI) RunCommand(logger *domain.Logger) *string {
	if c.prefix == nil {
		c.prefix = new(string)
		*c.prefix = "CLI:"
	}

	fmt.Println(*c.prefix + " Enter template name: ")
	var command string
	_, err := fmt.Scanln(&command)
	for err != nil {
		logger.Warn("Error reading command: " + err.Error())
		logger.Info(*c.prefix + " Enter template name: ")
		_, err = fmt.Scanln(&command)
	}
	
	full_resume := (*c.injector).GetFull(logger)
	(*c.generator).SetResume(full_resume)
	full_resume_string := (*c.generator).GetFull(logger)
	
	return full_resume_string
}

func (c *CLI) SetOptions(options ...*string) {
	c.prefix = options[0]
}

func NewCLI() ports.Commander {
	return &CLI{}
}
