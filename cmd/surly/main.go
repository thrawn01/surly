package main

import (
	"fmt"
	"os"

	"github.com/thrawn01/args"
	"github.com/thrawn01/surly"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

// Invoke like so
//    $ export CGO_ENABLED=0
//               |- image -| |-          standard golang arguments           -|
//    $ go-build golang:1.7  build -o amd64-my-prog -installsuffix static ./...
func main() {
	var config surly.BuilderConfig

	parser := args.NewParser()
	parser.AddArgument("image").
		Help("name of the docker image to build with")
	parser.AddOption("-output").Alias("-o").Required().
		Help("write the resulting executable or object to the named output file")
	parser.AddOption("-runtime").Default("docker").Choices(surly.GetBuilders()).
		Help("specify which image runtime to use")

	// Parse and exit with error if missing required arguments
	options := parser.ParseArgsSimple(nil)

	// Create a builder object that will run the go command
	builder, err := surly.Factory(config)
	checkErr(err)

	// Verify all the required options are provided
	checkErr(options.Required(builder.Required()))

	// Run the go command within the selected builder (rkt, docker, kvm)
	checkErr(builder.Run(parser.GetArgs()))

}
