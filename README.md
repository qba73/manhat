[![Go](https://github.com/qba73/manhat/actions/workflows/go.yml/badge.svg)](https://github.com/qba73/manhat/actions/workflows/go.yml)

# manhat

Module ```manhat``` provides functions and a ```CLI``` for calculating Manhattan-Distance.

# Installation and Usage

You can install the app in two ways:

- download binary for your OS and run it from your machine
- clone the repo, build binary and run it

## Download release

You can download latest release (binary) for your operating system and architecture from [here](https://github.com/qba73/manhat/releases).

## Clone repository
```
$ git clone https://github.com/qba73/manhat
```

## Build a binary for your OS

Linux
```
$ make build
```
macOS
```
$ make build_macos
```
Windows
```
$ make build_win
```

## Run the app

### Get help
```
$ ./bin/manhat -h
Usage of manhat:
  -location int
    	calculate Manhattan-Distance from given location to the center: manhat -location 12
  -version
    	show the version of the manhat app: manhat -version
```

### Get version
```
$ ./bin/manhat -version
Version: 0.1.0
GitRef: f2ae2914b2b072493176db2f5af0a24ed933c136
Build Time: 2021-07-19-20-31-25Z
```

### Calculate distance from given location
```
$ ./bin/manhat -location 12
3
```
```
$ ./bin/manhat -location 1024
31
```
```
$ ./bin/manhat -location 368078
371
```

# Development
## Make targets
```bash
$ make
dep                  Install Go dependencies
clean                Cleanup and remove artifacts
build                Build binary for Linux
build_macos          Build binary for Darwin (macOS)
build_win            Build binary for Windows
cover                Run tests with coverage report html format
test                 Run tests
vet                  Run Go vet
```

## Verify dependencies
```
$ make dep
```
or
```
$ go mod tidy
$ go mod verify
$ go mod vendor
```

## Run tests
```
$ make test
```
or
```
$ go test -race 
```

## Run tests and show coverage report
```
$ make cover
```
