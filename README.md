# Neuron OS Interview

This package contains the source code for my interview project with NeuronOS. This is written in go with 
two executables for the two deliverables (cmd/application, and cmd/installer). Start On Boot functionality
implemented for Mac and Linux.

## Build instructions

To binaries of the two executables use the following commands:

```
$ go build ./cmd/application    # HTTP Server
$ go build ./cmd/installer      # Start On Boot
```

Advanced usage can be found in the [#Installation_guide](Installation guide).


## API documentation

Our API can be broken down into two sets of documentation, one for the Go developer and one for HTTP clients.

### Go Interface

All logical source code can be found in the root directory under the package name `neuronos`. Examples for usage
can be found in `cmd/application/main.go`. Here is a basic outline of the public types and functions you will use.

```
type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
	Successful bool
	Time       time.Duration
}

type SystemInfo struct {
	Hostname  string
	IPAddress string
}

func NewCommander() Commander

func (c *commander) Ping(msg string) (PingResult, error)

func (c *commander) GetSystemInfo() (SystemInfo, error)
```

### HTTP Interface

Most clients will be using the HTTP server that is started by the application binary.
Here are example usages in `curl` that can be translated to your language of choice
using your favorite llm.

```
# Example usage for Ping feature.
$ curl -XPOST http://localhost:8080/execute -d '{"type":"ping","payload":"www.example.com"}'

# Example usage for GetSystemInfo feature.
$ curl -XPOST http://localhost:8080/execute -d '{"type":"sysinfo","payload":""}'
```

## Installation guide
To install the application server locally run the `go install` command given below.
```
$ go install ./cmd/application
```

And to have the application automatically start running on boot use the `go run` command below after running
the `go install` command above.
```
$ go run ./cmd/installer
```

## Testing
Unittests are written for all non-main packages. To run all use the `go test` command given below.
```
$ go test ./...
```

To test the server integration you can use the following commands:
```
$ go run ./cmd/application
```

```
$ curl -XPOST http://localhost:8080/execute -d '{"type":"ping","payload":"www.example.com"}'
$ curl -XPOST http://localhost:8080/execute -d '{"type":"sysinfo","payload":""}'
```

## A short clip that shows the "app in action"

https://github.com/user-attachments/assets/f9fd8c88-65e2-463d-843d-8c81db1fbb6d


## Interview Change Log

- 4:30 - Started on problem setup application structure using idiomatic go and snippets from prompt.
- 5:00 - Finished first pass of the implementation using my local linux development environment and blind coding for mac, to test on laptop.
- 5:05 - Working on documentation before switching to better test environment.
- 5:15 - Updated README file with basic go commands and instructions. Pushing first version to git version control.
- 5:50 - Finished running and polishing unittests and documentation.
