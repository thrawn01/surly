package surly

type BuilderConfig struct {
	Runtime    string
	OutputFile string
	Image      string
}

type BuilderFactory func(BuilderConfig) (Builder, error)

// Implement this interface to create a new builder
type Builder interface {
	Run([]string) error
	SetConfig(BuilderConfig)
	Required() []string
}

// Noop concrete builder
type Build struct {
}

// Noop run impl
func (self *Build) Run(args []string) error {
	return nil
}

// No options are required
func (self *Build) Required() []string {
	return []string{}
}

var _ Builder = &Build{}

// Builds a new Builder object from the options provided
func Factory(builder string) (Builder, error) {
	return &Build{}, nil
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
