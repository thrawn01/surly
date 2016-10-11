package surly

import (
	"log"
	"strings"

	"fmt"

	"github.com/ahmetalpbalkan/go-dexec"
	"github.com/fatih/structs"
	"github.com/fsouza/go-dockerclient"
	"github.com/pkg/errors"
)

type DockerConfig struct {
	Output  string
	Runtime string
	Image   string
}

type DockerBuilder struct {
	config DockerConfig
	client *docker.Client
}

// Returns a new docker builder using the passed config
func NewDockerBuilder(config DockerConfig) (*DockerBuilder, error) {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return nil, errors.Wrap(err, "surly.NewDockerBuilder()")
	}

	return &DockerBuilder{
		config: config,
		client: client,
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
	exec := dexec.Docker{Client: self.client}

	m, _ := dexec.ByCreatingContainer(docker.CreateContainerOptions{
		Config: &docker.Config{Image: "busybox"}})

	cmd := exec.Command(m, "echo", `I am running inside a container!`)
	b, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
	return nil
}

var _ Builder = &DockerBuilder{}
