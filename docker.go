package surly

import (
	"os/exec"
	"strings"

	"fmt"

	"os"

	"github.com/fatih/structs"
	"github.com/pkg/errors"
)

type DockerConfig struct {
	Image      string
	WorkingDir string // The working directory inside the container
	GoPath     string // $GOPATH
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
func (self *DockerBuilder) Run(srcArgs []string) error {

	args := []string{
		"run",
		"--rm",
		"-w", self.config.WorkingDir,
		"-v", fmt.Sprintf("%s:/go", self.config.GoPath),
		"-e", "CGO_ENABLED=0",
		self.config.Image,
		"go",
	}

	args = append(args, srcArgs...)

	fmt.Printf("args: %+v\n", args)
	cmd := exec.Command("docker", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		/*if exitErr, ok := err.(*exec.ExitError); ok == true {
			fmt.Println("err")
			fmt.Print(string(exitErr.Stderr))
		}*/
		return errors.Wrap(err, "exec.Comand()")
	}
	return nil
}

var _ Builder = &DockerBuilder{}
