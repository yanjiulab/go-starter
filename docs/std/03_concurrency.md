# Concurrency

## Goroutine

Goroutines are lightweight threads managed by the Go runtime.

### Starting Goroutines

Use the `go` keyword to start a goroutine.

```go
go func() {
    fmt.Println("Running in goroutine")
}()
```

### Goroutine Lifecycle

Goroutines run concurrently and exit when their function completes.

## Channel

Channels are typed conduits for communication between goroutines.

### Creating Channels

Channels are created with `make`.

```go
ch := make(chan int)
```

### Sending and Receiving

Use `<-` to send and receive data.

```go
ch <- 42
value := <-ch
```

### Buffered Channels

Buffered channels have capacity.

```go
ch := make(chan int, 10)
```

### Closing Channels

Close channels to signal completion.

```go
ch := make(chan int)
go func() {
    ch <- 1
    close(ch)
}()
for v := range ch {
    fmt.Println(v)
}
```

## Select Multiplexing

Select allows waiting on multiple channel operations.

### Basic Select

Select chooses the first ready case.

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
}
```

### Default Case

Default case executes if no channel is ready.

```go
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("No message")
}
```

## sync Package

The sync package provides synchronization primitives.

### sync.Mutex

Mutex provides mutual exclusion.

```go
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

### sync.RWMutex

RWMutex allows multiple readers or one writer.

```go
var rwmu sync.RWMutex
rwmu.RLock()
// read
rwmu.RUnlock()
```

### sync.WaitGroup

WaitGroup waits for goroutines to finish.

```go
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); /* work */ }()
wg.Wait()
```

### sync.Once

Once ensures a function runs only once.

```go
var once sync.Once
once.Do(func() { /* init */ })
```

### sync.Cond

Cond provides condition variables.

```go
cond := sync.NewCond(&mu)
cond.Wait() // wait for signal
cond.Signal() // wake one waiter
```

### sync.Map

Map is a concurrent map.

```go
var m sync.Map
m.Store("key", "value")
value, ok := m.Load("key")
```

### Atomic Operations

Atomic provides low-level atomic operations.

```go
var counter int64
atomic.AddInt64(&counter, 1)
```

## Context

Context manages goroutine lifecycles and cancellation.

### Creating Contexts

Use context.WithCancel for cancellation.

```go
ctx, cancel := context.WithCancel(context.Background())
```

### Using Context

Pass context to functions for cancellation.

```go
func worker(ctx context.Context) {
    select {
    case <-ctx.Done():
        return
    // work
    }
}
```

### Timeout and Deadline

Use WithTimeout or WithDeadline.

```go
ctx, _ := context.WithTimeout(context.Background(), time.Second)
```

## Patterns and Design

Common concurrency patterns in Go.

### Worker Pools

Use channels to distribute work.

```go
jobs := make(chan int, 100)
results := make(chan int, 100)

for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
}
```

### Fan-In/Fan-Out

Combine or split channels.

### Pipeline

Chain goroutines with channels for data processing.

## Goroutine Management

### Prevent Goroutine Leaks

Always ensure goroutines can exit.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
go func(ctx context.Context) {
    select {
    case <-ctx.Done():
        return
    }
}(ctx)
```

### Timeout Pattern

Use context for timeouts.

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
select {
case result := <-ch:
    fmt.Println(result)
case <-ctx.Done():
    fmt.Println("Timeout")
}
```

## sync.Pool

Pool reuses allocated objects to reduce GC pressure.

```go
var pool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 0, 1024)
    },
}
obj := pool.Get()
defer pool.Put(obj)
```

## Race Condition Detection

Use `-race` flag to detect data races.

```bash
go run -race main.go
go test -race ./...
```
