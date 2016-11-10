# Synopsis
A tool to assist with cross compiling golang projects that
require a full linux operating system at compile time.

Using docker images, surly mounts your `$GOPATH` inside a docker image
preforms the compile and then places the resulting build in your local
 `$GOPATH` go environment.

## Origin
Started as an idea from @NuttySwiss
https://twitter.com/nuttyswiss/status/770049332506660866

# How do I use it?
```bash
$ surly -h
Usage:  [OPTIONS]

Options:
  -image         name of the docker image to build with (Default=golang:1.7.1-alpine)
  -runtime       specify which image runtime to use (Default=docker)
  -working-dir   working directory inside the container (Default=/)
  -go-path       path to our go development environment (Env=GOPATH)
  -h, --help     Display this help message and exit
```

Build any golang project using golang1.7.1-alpine docker image
```bash
$ surly build github.com/thrawn01/surly
```
Install any golang project CLI
```bash
$ surly build github.com/thrawn01/surly/cmd/surly
```

# Installation
```bash
$ go install github.com/thrawn01/surly/cmd/surly
```

# Why Name it Surly?
* Because I didn't want to name it 'go-image-cross-compiler'
* Because I didn't want to name it 'cross'
* Because 'surly' is a synonym of 'cross'


