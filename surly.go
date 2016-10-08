package surly

import "github.com/pkg/errors"

type BuilderConfig struct {
	Runtime    string
	OutputFile string
	Image      string
}

type BuilderFactory func(map[string]interface{}) (Builder, error)

// Implement this interface to create a new builder
type Builder interface {
	Run([]string) error
}

// Builds a new Builder object from the options provided
func Factory(builder string, config map[string]interface{}) (Builder, error) {
	factory, ok := SupportedBuilders[builder]
	if !ok {
		errors.Errorf("Unsupported builder '%s' valid builders are(%s)", builder, GetBuilders())
	}
	return factory(config)
}

// A list of supported builders
func GetBuilders() []string {
	keys := make([]string, 0, len(SupportedBuilders))
	for key := range SupportedBuilders {
		keys = append(keys, key)
	}
	return keys
}

// A map of supported builder factories
var SupportedBuilders map[string]BuilderFactory = map[string]BuilderFactory{
	"docker": DockerFactory,
}
