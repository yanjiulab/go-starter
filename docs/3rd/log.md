# Logging - Structured Logging with Zap

Structured logging is essential for production applications. This guide uses Zap and Lumberjack.

## Installation

```bash
go get -u go.uber.org/zap
go get -u gopkg.in/natefinch/lumberjack.v2
```

## Basic Usage

### Simple Logger

```go
import "go.uber.org/zap"

logger, err := zap.NewProduction()
defer logger.Sync()

logger.Info("Application started", zap.String("version", "1.0.0"))
```

### Development Logger

```go
logger, err := zap.NewDevelopment()
defer logger.Sync()

logger.Info("Debug info", zap.Int("count", 42))
```

## Structured Fields

Zap provides type-safe field logging.

```go
logger.Info("User login",
    zap.String("username", "alice"),
    zap.Int("user_id", 123),
    zap.Duration("response_time", 50*time.Millisecond),
)
```

## Log Levels

- **Debug** — Detailed diagnostic information
- **Info** — General informational messages
- **Warn** — Warning messages for potential issues
- **Error** — Error messages
- **DPanic** — Debug panic
- **Panic** — Panic
- **Fatal** — Fatal error

## File Rotation with Lumberjack

```go
import (
    "gopkg.in/natefinch/lumberjack.v2"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

w := &lumberjack.Logger{
    Filename:   "/var/log/app.log",
    MaxSize:    100, // megabytes
    MaxBackups: 3,
    MaxAge:     28, // days
}

core := zapcore.NewCore(
    zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
    zapcore.AddSync(w),
    zap.InfoLevel,
)

logger := zap.New(core)
defer logger.Sync()
```

## Global Logger

```go
// Initialize global logger
var Logger *zap.Logger

func init() {
    var err error
    Logger, err = zap.NewProduction()
    if err != nil {
        panic(err)
    }
}

// Use anywhere
Logger.Info("Event", zap.String("event_type", "user_login"))
