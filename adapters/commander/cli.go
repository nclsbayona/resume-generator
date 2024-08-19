package commander

// Options: prefix (Prefix to add to the CLI commands)

import (
	"fmt"

	"github.com/nclsbayona/resume-generator/domain"
	"github.com/nclsbayona/resume-generator/ports"
)

type CLI struct {
	prefix    *string
	injector  *ports.Injector
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

	var options []*string = make([]*string, 0)
	var command *string = new(string)
	var opt byte

	fmt.Println("Does the injector need to set options? (y/n)")
	_, err := fmt.Scanln(command)
	for (err != nil) || (*command != "y" && *command != "n") {
		fmt.Println("Invalid option. Please enter y or n")
		_, err = fmt.Scanln(command)
	}
	opt = (*command)[0]
	if opt == 'y' {
		for err != nil || opt != 'n' {
			fmt.Println(*c.prefix + " Enter option: ")
			_, err = fmt.Scanln(command)
			for err != nil {
				fmt.Println("Invalid input. Please enter command again")
				_, err = fmt.Scanln(command)
			}
			opt_string := *command
			options = append(options, &opt_string)
			fmt.Println("Do you want to add another option? (y/n)")
			_, err = fmt.Scanln(command)
			for (err != nil) || (*command != "y" && *command != "n") {
				fmt.Println("Invalid option. Please enter y or n")
				_, err = fmt.Scanln(command)
			}
			opt = (*command)[0]
		}
		(*c.injector).SetOptions(options...)
	}

	options = nil
	options = make([]*string, 0)

	fmt.Println("Does the generator need to set options? (y/n)")
	_, err = fmt.Scanln(command)
	for (err != nil) || (*command != "y" && *command != "n") {
		fmt.Println("Invalid option. Please enter y or n")
		_, err = fmt.Scanln(command)
	}
	opt = (*command)[0]
	if opt == 'y' {
		for err != nil || opt != 'n' {
			fmt.Println(*c.prefix + " Enter option: ")
			_, err = fmt.Scanln(command)
			for err != nil {
				fmt.Println("Invalid input. Please enter command again")
				_, err = fmt.Scanln(command)
			}
			opt_string := *command
			options = append(options, &opt_string)

			fmt.Println("Do you want to add another option? (y/n)")
			_, err = fmt.Scanln(command)
			for (err != nil) || (*command != "y" && *command != "n") {
				fmt.Println("Invalid option. Please enter y or n")
				_, err = fmt.Scanln(command)
			}
			opt = (*command)[0]
		}
		(*c.generator).SetOptions(options...)

	}

	full_resume := (*c.injector).GetFull(logger)
	(*c.generator).SetResume(full_resume)
	full_resume_string := (*c.generator).GetFull(logger)

	return full_resume_string
}

func (c *CLI) SetOptions(options ...*string) {
	c.prefix = options[0]
}

func NewCLI() *ports.Commander {
	var cli ports.Commander = new(CLI)
	return &cli
}
