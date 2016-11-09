package main

import (
	"os"

	"fmt"

	"strings"

	"github.com/thrawn01/args"
	"github.com/thrawn01/surly"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func removeHyphens(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for key, value := range src {
		dst[strings.Replace(key, "-", "", -1)] = value
	}
	return dst
}

// Invoke like so
// surly build cmd/eventbus/main.go -image=golang:1.7.1-alpine -o eventbus-go
func main() {
	parser := args.NewParser()
	parser.AddOption("-image").
		Help("name of the docker image to build with").Default("golang:1.7.1-alpine")
	//parser.AddOption("-output").Alias("-o").Required().
	//	Help("write the resulting executable or object to the named output file")
	parser.AddOption("-runtime").Default("docker").Choices(surly.GetBuilders()).
		Help("specify which image runtime to use")
	parser.AddOption("-working-dir").Default("/").Help("working directory inside the container")
	parser.AddOption("-go-path").Env("GOPATH").Help("path to our go development environment")

	// Parse and exit with error if missing required arguments
	options := parser.ParseArgsSimple(nil)

	// Create a builder object that will run the go command
	builder, err := surly.Factory(options.String("runtime"), removeHyphens(options.ToMap()))
	checkErr(err)

	// Run the go command within the selected builder (rkt, docker, kvm)
	checkErr(builder.Run(parser.GetArgs()))
}
