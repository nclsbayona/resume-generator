package main

import (
	"flag"
	"github.com/nclsbayona/resume-generator/adapters/commander"
	"github.com/nclsbayona/resume-generator/adapters/generator"
	"github.com/nclsbayona/resume-generator/adapters/injector"
	"github.com/nclsbayona/resume-generator/domain"
	"github.com/nclsbayona/resume-generator/ports"
	"os"
)

func commanderDI(commander_to_use *string, injector_implementation *ports.Injector, generator_implementation *ports.Generator, options ...*string) (commander_implementation *ports.Commander) {
	switch *commander_to_use {
	case "cli":
		commander_implementation = commander.NewCLI()
		(*commander_implementation).SetInjector(injector_implementation)
		(*commander_implementation).SetGenerator(generator_implementation)
		(*commander_implementation).SetOptions(options...)
	case "config":
		commander_implementation = commander.NewConfig()
		(*commander_implementation).SetInjector(injector_implementation)
		(*commander_implementation).SetGenerator(generator_implementation)
		(*commander_implementation).SetOptions(options...)
	default:
		commander_implementation = commander.NewCLI()
		(*commander_implementation).SetInjector(injector_implementation)
		(*commander_implementation).SetGenerator(generator_implementation)
		(*commander_implementation).SetOptions(options...)
	}
	return
}

func generatorDI(generator_to_use *string) (generator_implementation *ports.Generator) {
	switch *generator_to_use {
	case "html":
		generator_implementation = generator.NewWriteHTML()
	default:
		generator_implementation = generator.NewWriteHTML()
	}
	return
}

func injectorDI(injector_to_use *string) (injector_implementation *ports.Injector) {
	switch *injector_to_use {
	case "yaml":
		injector_implementation = injector.NewYaml()
	default:
		injector_implementation = injector.NewYaml()
	}
	return
}

func main() {

	var logger *domain.Logger = domain.NewLogger(os.Stdout)
	var commander *string = new(string)
	var injector *string = new(string)
	var generator *string = new(string)
	var output_file *string = new(string)
	var config_properties []*string = make([]*string, 0)

	flag.StringVar(commander, "c", "cli", "Commander to use. Default is cli")
	flag.StringVar(injector, "i", "yaml", "Injector to use. Default is yaml")
	flag.StringVar(generator, "g", "html", "Generator to use. Default is html")
	flag.StringVar(output_file, "o", "resume.html", "Output file. Default is resume.html")

	flag.Parse()

	var commander_properties []string = flag.Args()

	for _, arg := range commander_properties {
		config_properties = append(config_properties, &arg)
	}

	// Here dependency injection is supposed to happen
	var generator_implementation *ports.Generator = generatorDI(generator)
	var injector_implementation *ports.Injector = injectorDI(injector)
	var commander_implementation *ports.Commander = commanderDI(commander, injector_implementation, generator_implementation, config_properties...)
	//

	file, err := os.Create(*output_file)
	if err != nil {
		logger.Fatal("Error creating output file: " + err.Error())
	}
	defer file.Close()
	file.WriteString(*(*commander_implementation).RunCommand(logger))
}
