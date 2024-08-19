package commander

// Options: properties_file_name (File to read the properties from)

import (
	"os"

	"github.com/nclsbayona/resume-generator/domain"
	"github.com/nclsbayona/resume-generator/ports"
	"gopkg.in/yaml.v3"
)

type sProperties struct {
	GeneratorProperties []*string `yaml:"generator_properties"`
	InjectorProperties  []*string `yaml:"injector_properties"`
}

type Config struct {
	propertiesFileName *string
	properties         *sProperties
	injector           *ports.Injector
	generator          *ports.Generator
}

func (c *Config) RunCommand(logger *domain.Logger) *string {
	if c.properties == nil {
		logger.Error("No properties found")
	}
	logger.Info("Running command")
	logger.Info("Setting options")
	logger.Info("Setting injector options")
	logger.Info(*c.properties.InjectorProperties[0])
	(*c.injector).SetOptions(c.properties.InjectorProperties...)
	logger.Info("Setting generator options")
	(*c.generator).SetOptions(c.properties.GeneratorProperties...)
	full_resume := (*c.injector).GetFull(logger)
	(*c.generator).SetResume(full_resume)
	full_resume_string := (*c.generator).GetFull(logger)

	return full_resume_string
}

func (c *Config) SetGenerator(generator *ports.Generator) {
	c.generator = generator
}

func (c *Config) SetInjector(injector *ports.Injector) {
	c.injector = injector
}

func NewConfig() *ports.Commander {
	var config ports.Commander = new(Config)
	return &config
}
func (c *Config) parseProperties() {
	c.properties = new(sProperties)
	var out []byte
	out, err := os.ReadFile(*c.propertiesFileName)
	if err == nil {
		_ = yaml.Unmarshal(out, c.properties)
	} else {
		panic("Error reading file: " + *c.propertiesFileName)
	}
	os.Stdout.WriteString("Properties read successfully\n" + string(out) + "\n")
}

func (c *Config) SetOptions(options ...*string) {

	c.propertiesFileName = options[0]
	c.parseProperties()

}
