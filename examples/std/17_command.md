# Go 命令速查（单知识点）

```bash
# Module setup
go mod init example.com/demo
go mod tidy

# Run one sample file
go run ./examples/std/00_basic.go

# Build / test / quality
go build ./...
go test ./...
go fmt ./...
go vet ./...

# Discover docs and environment
go doc fmt.Println
go env
```
