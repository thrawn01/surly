package surly

type Docker struct{}

// Run the specified golang command inside a container
func (self *Docker) Run(args []string) error {
	return nil
}

// Return what map keys are required to be passed into DockerFactory()
func (self *Docker) Required() []string {
	return []string{"image", "output"}
}

// Create a new instance of the docker builder object
func DockerFactory(config BuilderConfig) (Builder, error) {
	return nil, nil
}

var _ Builder = &Docker{}
