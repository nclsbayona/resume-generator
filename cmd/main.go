package main

import (
	"flag"
	"github.com/nclsbayona/resume-generator/pkg/adapters/commander"
	"github.com/nclsbayona/resume-generator/pkg/adapters/generator"
	"github.com/nclsbayona/resume-generator/pkg/adapters/injector"
	"github.com/nclsbayona/resume-generator/pkg/core/domain"
	"github.com/nclsbayona/resume-generator/pkg/core/ports"
	"gopkg.in/yaml.v3"
	"os"
)

type sProperties struct {
	PropertiesFileName *string
	Injector           *string `yaml:"injector"`
	Generator          *string `yaml:"generator"`
	Commander          *string `yaml:"commander"`
	Template           *string `yaml:"template"`
	InputFile          *string `yaml:"input_file_name"`
	OutputFile         *string `yaml:"output_file_name"`
}

func commanderDI(properties *sProperties, injector_implementation ports.Injector, generator_implementation ports.Generator) (commander_implementation ports.Commander) {
	switch *properties.Commander {
	case "html":
		commander_implementation = commander.NewCLI()
		commander_implementation.SetInjector(&injector_implementation)
		commander_implementation.SetGenerator(&generator_implementation)
	default:
		commander_implementation = commander.NewCLI()
		commander_implementation.SetInjector(&injector_implementation)
		commander_implementation.SetGenerator(&generator_implementation)
	}
	return
}

func generatorDI(properties *sProperties) (generator_implementation ports.Generator) {
	switch *properties.Generator {
	case "html":
		generator_implementation = generator.NewWriteHTML()
		generator_implementation.SetOptions(properties.Template)
	default:
		generator_implementation = generator.NewWriteHTML()
		generator_implementation.SetOptions(properties.Template)
	}
	return
}

func injectorDI(properties *sProperties) (injector_implementation ports.Injector) {
	switch *properties.Injector {
	case "yaml":
		injector_implementation = injector.NewYaml()
		injector_implementation.SetOptions(properties.InputFile, properties.Template)
	default:
		injector_implementation = injector.NewYaml()
		injector_implementation.SetOptions(properties.InputFile, properties.Template)
	}
	return
}

func parseProperties(properties_file *string, logger *domain.Logger) (properties *sProperties) {
	properties = new(sProperties)
	properties.PropertiesFileName = properties_file
	var out []byte
	out, err := os.ReadFile(*properties_file)
	if err != nil {
		logger.Error("Error reading file: " + *properties_file)
	}
	err = yaml.Unmarshal(out, properties)
	if err != nil {
		logger.Error("Error unmarshalling yaml: " + err.Error())
		return nil
	}
	logger.Info("Properties file " + *properties_file + " read successfully")
	return
}
func main() {

	var logger *domain.Logger = domain.NewLogger(os.Stdout)
	var config_properties string

	flag.StringVar(&config_properties, "c", "config.yaml", "Configuration file. Default is ../config.yaml")
	flag.Parse()

	properties := parseProperties(&config_properties, logger)

	logger.Info("Output file: " + *properties.OutputFile)
	logger.Info("Starting resume generation using template \"" + *properties.Template + "\"")

	file, err := os.Create(*properties.OutputFile)
	if err != nil {
		logger.Error("Error creating file " + err.Error())
	}
	defer file.Close()
	logger.Info("Output file " + *properties.OutputFile + " created successfully")

	// Here dependency injection is supposed to happen
	var generator_implementation ports.Generator = generatorDI(properties)
	var injector_implementation ports.Injector = injectorDI(properties)
	var commander_implementation ports.Commander = commanderDI(properties, injector_implementation, generator_implementation)
	//

	if _, err = file.WriteString(*commander_implementation.RunCommand(logger)); err != nil {
		logger.Error("Resume generation failed!")
	}
	logger.Info("Resume generation complete")
}
