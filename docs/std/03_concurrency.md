# Concurrency

## Goroutine

## Channel

## Select 多路复用

专门配合 Channel 使用，能同时监听多个通道的读写操作，哪个通道就绪就执行哪个，实现高效的并发调度。

## `sync` package

Go 标准库提供的并发同步工具，包含互斥锁、等待组、原子操作等，用于处理需要共享内存的并发场景。

### 互斥锁 `sync.Mutex`

### 读写锁 `sync.RWMutex`

### 等待组 `sync.WaitGroup`

### 一次性执行 `sync.Once`

### 条件变量 `sync.Cond`

### 内存映射 `sync.Map`

### 原子操作 atomic

## context

用来控制 Goroutine 的生命周期，可实现取消、超时、传值，统一管理多层嵌套的并发任务

## 模式与设计

工程化的并发实践方案
