# CMD - Cobra CLI Framework

Cobra is a powerful library for building modern CLI applications in Go.

## Installation

```bash
# Install Cobra library
go get -u github.com/spf13/cobra@latest

# Install Cobra CLI tool
go install github.com/spf13/cobra-cli@latest

# Ensure GOPATH/bin is in PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

## Project Initialization

```bash
mkdir go-starter
cd go-starter
go mod init github.com/yanjiulab/go-starter
cobra-cli init --author "Yanjiulab" --license MIT --viper

cobra-cli add serve
cobra-cli add cli

cobra-cli add create -p 'configCmd'
```

## Common Commands

### Run CLI Application

```bash
go run main.go serve
go run main.go cli
```

### Add New Command

```bash
cobra-cli add mycommand
cobra-cli add subcommand -p 'mycommandCmd'
```

### Build Binary

```bash
go build -o go-starter
./go-starter serve --help
```
