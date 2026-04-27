# Project

## Module & Package

Go modules and packages provide the structure for building and managing Go projects.

### Module Initialization

Use `go mod init` to create a new module and manage dependencies.

### Package Organization

Group related code into packages and use clear package names.

### init()

The `init()` function runs before `main()` and is useful for package initialization.

```go
func init() {
    // setup logic
}
```

## Project Structure

Follow standard directory conventions to keep code organized.

### Common Layout

Use directories like `cmd`, `pkg`, `internal`, `configs`, and `examples`.

### Separation of Concerns

Keep reusable libraries separate from application-specific code.

## Coding Standards

Use consistent formatting and naming rules.

### Formatting

Use `gofmt` or `go fmt` to format code automatically.

### Naming Conventions

Use short, descriptive names and avoid underscores in package names.

## Logging

Use structured logging for observability and debugging.

### Log Levels

Use levels like `Info`, `Warn`, `Error`, and `Debug`.

### Output Format

Prefer JSON or structured logs for production systems.

## Time Handling

Use the `time` package for working with timestamps and durations.

### Time Parsing and Formatting

Use `time.Parse` and `time.Format`.

### Time Zones

Handle time zones explicitly with `time.Location`.

## File IO

Use the `os` and `io` packages for reading and writing files.

### Reading Files

Use `os.ReadFile` or `bufio.Reader`.

### Writing Files

Use `os.WriteFile` or create writers for large output.

## Encoding

Common encodings convert data between formats.

### JSON

Use `encoding/json` for JSON serialization and deserialization.

### hex

Use `encoding/hex` to encode and decode hexadecimal strings.

### base64

Use `encoding/base64` for Base64 encoding and decoding.

## Network

Go supports low-level networking through the standard library.

### net

Use `net` for generic networking operations and address resolution.

### UDP

Use `net.ListenPacket` for UDP sockets and datagram communication.

### TCP

Use `net.Listen` and `net.Dial` for TCP servers and clients.

## HTTP

Build HTTP clients and servers using `net/http`.

### HTTP Server

Use `http.HandleFunc` and `http.ListenAndServe`.

### HTTP Client

Use `http.Get`, `http.Post`, or custom `http.Client`.

## Command-Line Arguments

Use the `flag` package or third-party libraries for command-line parsing.

### flag Package

Define flags with `flag.String`, `flag.Int`, and `flag.Bool`.

### CLI Libraries

Use libraries like Cobra for complex CLI applications.

## Configuration Management

Load configuration from files, environment variables, or remote sources.

### Environment Variables

Use `os.Getenv` and `os.LookupEnv`.

### Configuration Files

Use JSON, YAML, or TOML files and parse them at startup.

## Timer

Use timers and tickers for scheduled tasks.

### time.Timer

Create a one-time timer with `time.NewTimer`.

### time.Ticker

Use `time.NewTicker` for periodic actions.

## Database

Access databases using drivers and libraries.

### Redis

Use a Redis client for caching and fast data access.

### MySQL

Use database/sql with a MySQL driver for relational storage.

### ORM

Use ORMs like GORM for higher-level database mapping.

### Cache

Use local or distributed caching for performance.

## Unit Test

Write tests using the standard `testing` package.

### testing

Create test functions with `func TestXxx(t *testing.T)`.

### Mock

Use stubs or mock libraries to isolate test dependencies.

## Memory Optimization

Reduce memory usage with careful data choices.

### Avoid Large Allocations

Prefer streaming and reuse buffers when possible.

### Use Value Types Appropriately

Choose between values and pointers to minimize allocations.

## Performance Analysis

Profile and benchmark your application.

### Benchmarking

Use `go test -bench` to benchmark functions.

### Profiling

Use `pprof` to collect CPU and memory profiles.

## Garbage Collection

Understand Go's garbage collector and minimize pressure.

### GC Tuning

Use runtime metrics and environment variables to monitor GC.

### Reducing Allocation Rate

Reuse objects and avoid unnecessary allocations.

## Cross Compilation

Build binaries for different OS and architectures.

### GOOS and GOARCH

Use environment variables like `GOOS=linux GOARCH=amd64 go build`.

### Multi-Platform Builds

Use build scripts or CI to produce artifacts for multiple targets.
