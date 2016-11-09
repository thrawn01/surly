package surly

import (
	"strings"

	"fmt"

	"path/filepath"

	"github.com/ahmetalpbalkan/go-dexec"
	"github.com/fatih/structs"
	"github.com/fsouza/go-dockerclient"
	"github.com/pkg/errors"
)

type DockerConfig struct {
	Output     string // output to copy the resulting executable
	Runtime    string
	Image      string
	WorkingDir string // The working directory inside the container
	GoPath     string // $GOPATH
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

	// Calculate the absolute path of the output
	config.Output, err = filepath.Abs(config.Output)
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
		// Lowercase and remove any hyphens
		fieldName := strings.Replace(strings.ToLower(name), "-", "", -1)
		value, ok := src[fieldName]
		if !ok {
			errors.Errorf("DockerBuilder: Missing required field %s", fieldName)
		}
		field := s.Field(name)
		field.Set(value)
	}
	return NewDockerBuilder(config)
}

func (self *DockerBuilder) BuildMounts() []docker.Mount {
	return []docker.Mount{
		{
			Name:        "GoPathMount",
			Source:      self.config.GoPath,
			Destination: "/go",
		},
		/*{
			Source:      path.Base(self.config.Output),
			Destination: "/go/bin",
		},*/
	}
}

// Run the specified golang command inside a container
func (self *DockerBuilder) Run(args []string) error {
	exec := dexec.Docker{Client: self.client}

	dockerConfig := docker.Config{
		Image:      self.config.Image,
		Mounts:     self.BuildMounts(),
		Env:        []string{"CGO_ENABLED=0"},
		WorkingDir: self.config.WorkingDir,
	}

	m, err := dexec.ByCreatingContainer(docker.CreateContainerOptions{Config: &dockerConfig})
	if err != nil {
		return errors.Wrap(err, "ByCreatingContainer()")
	}

	/*args := make([]string, len(srcArgs)+1)
	args[0] = "go"
	for i := 1; i < len(args); i++ {
		args[i] = srcArgs[i-1]
	}*/

	cmd := exec.Command(m, "go", args...)
	b, err := cmd.Output()
	fmt.Printf("%s", b)
	if err != nil {
		if exitErr, ok := err.(*dexec.ExitError); ok == true {
			fmt.Print(string(exitErr.Stderr))
		}
		return errors.Wrap(err, "exec.Comand()")
	}
	return nil
}

var _ Builder = &DockerBuilder{}
