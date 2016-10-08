package surly

import (
	"strings"

	"github.com/fatih/structs"
	"github.com/pkg/errors"
)

type DockerConfig struct {
	Output  string
	Runtime string
	Image   string
}

type DockerBuilder struct {
	config DockerConfig
}

// Returns a new docker builder using the passed config
func NewDockerBuilder(config DockerConfig) (*DockerBuilder, error) {
	return &DockerBuilder{
		config: config,
	}, nil
}

// Create a new instance of the docker builder object
func DockerFactory(src map[string]interface{}) (Builder, error) {
	config := DockerConfig{}
	s := structs.New(&config)
	for _, name := range s.Names() {
		fieldName := strings.ToLower(name)
		value, ok := src[fieldName]
		if !ok {
			errors.Errorf("DockerBuilder: Missing required field %s", fieldName)
		}
		field := s.Field(name)
		field.Set(value)
	}
	return NewDockerBuilder(config)
}

// Run the specified golang command inside a container
func (self *DockerBuilder) Run(args []string) error {
	return nil
}

var _ Builder = &DockerBuilder{}
