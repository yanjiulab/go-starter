# Master Go in 10 Minutes for C++ Developers

2500字左右

本次已对 `examples` 下示例进行归类与去重，并重组为 `examples/std`：

- 单文件单知识点
- 避免“一个文件塞多个主题”
- 降低重复演示内容

---

## 示例映射（`examples/std`）

### 基础与数据结构

- `examples/std/00_basic.go`：基础语法（变量、流程控制、函数、defer）
- `examples/std/01_error.go`：错误处理（error、errors.Is、panic/recover）
- `examples/std/02_slice.go`：切片（append、共享底层数组、copy）
- `examples/std/03_map.go`：映射（读写、判断存在、删除、遍历）
- `examples/std/04_string.go`：字符串与 UTF-8（bytes/runes）
- `examples/std/05_struct.go`：结构体与嵌入字段
- `examples/std/06_func_type.go`：函数类型与高阶函数
- `examples/std/07_interface.go`：接口与多态

### 并发

- `examples/std/08_goroutine.go`：goroutine 启动与同步退出
- `examples/std/09_channel.go`：channel 基础收发与关闭
- `examples/std/10_sync.go`：WaitGroup + Mutex 基础协作
- `examples/std/11_context.go`：context 超时取消
- `examples/std/12_worker_pool.go`：Worker Pool 模式
- `examples/std/13_pipeline.go`：Pipeline 模式

### 进阶

- `examples/std/14_closure.go`：闭包状态保持
- `examples/std/15_generics.go`：泛型约束与泛型函数
- `examples/std/16_pprof.go`：pprof 诊断入口
- `examples/std/17_command.md`：Go 命令速查

---

## 运行方式

在项目根目录执行：

```bash
go run ./examples/std/00_basic.go
go run ./examples/std/01_error.go
go run ./examples/std/02_slice.go
go run ./examples/std/03_map.go
go run ./examples/std/04_string.go
go run ./examples/std/05_struct.go
go run ./examples/std/06_func_type.go
go run ./examples/std/07_interface.go
go run ./examples/std/08_goroutine.go
go run ./examples/std/09_channel.go
go run ./examples/std/10_sync.go
go run ./examples/std/11_context.go
go run ./examples/std/12_worker_pool.go
go run ./examples/std/13_pipeline.go
go run ./examples/std/14_closure.go
go run ./examples/std/15_generics.go
go run ./examples/std/16_pprof.go
```

pprof 示例运行后访问：

- `http://localhost:6060/debug/pprof/`
