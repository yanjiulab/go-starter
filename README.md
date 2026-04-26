# Go 学习导图（标准示例版）

本文配套 `examples/std`。目标是：**每个文件一个知识点**，但内容尽量覆盖日常开发常用写法。

---

## 学习顺序（建议）

1. 语法与错误：`00_basic.go` -> `01_error.go`
2. 核心数据结构：`02_slice.go` -> `03_map.go` -> `04_string.go` -> `05_struct.go`
3. 抽象与函数：`06_func_type.go` -> `07_interface.go`
4. 并发与控制：`08_goroutine.go` -> `09_channel.go` -> `10_sync.go` -> `11_context.go`
5. 模式与进阶：`12_worker_pool.go` -> `13_pipeline.go` -> `14_closure.go` -> `15_generics.go` -> `16_pprof.go`
6. 命令行：`17_command.md`

---

## 示例清单与重点

### 基础

- `examples/std/00_basic.go`
  - 常量、短变量声明、`if/for/switch/defer`
  - 函数与多返回值
  - 错误分支处理（`divide`）
  - 注释风格：文件中已增加英文注释，方便跨语言团队协作

- `examples/std/01_error.go`
  - sentinel error（`ErrNegative`）
  - 自定义错误类型（`ValidationError`）
  - `errors.Is` / `errors.As`
  - `panic/recover` 最小可运行示例

### 数据结构

- `examples/std/02_slice.go`
  - `nil` slice vs empty slice
  - `append` 扩容特性
  - 切片共享底层数组
  - `copy` 与深拷贝技巧

- `examples/std/03_map.go`
  - map 字面量初始化与更新
  - `value, ok := m[k]`
  - `delete` / `clear`
  - 遍历顺序无保证（Go 设计如此）

- `examples/std/04_string.go`
  - UTF-8 字节长度 vs rune 个数
  - `range` 按 rune 遍历
  - `strings` 常见操作：`Contains`、`ReplaceAll`、`Split`、`ToUpper`

- `examples/std/05_struct.go`
  - struct 组合与嵌入字段
  - 值接收者与指针接收者
  - 面向对象风格最小实践

### 抽象能力

- `examples/std/06_func_type.go`
  - 函数类型定义
  - 高阶函数（函数作为参数）
  - 函数工厂（函数作为返回值）

- `examples/std/07_interface.go`
  - 接口定义与隐式实现
  - 多态调用（`[]Shape`）
  - 类型断言获取具体类型

### 并发

- `examples/std/08_goroutine.go`
  - 并发启动多个任务
  - channel 回收完成信号

- `examples/std/09_channel.go`
  - 单向 channel 参数
  - range 读取到关闭
  - `select + time.After` 超时控制

- `examples/std/10_sync.go`
  - `WaitGroup` 等待收敛
  - `Mutex` 保护共享变量
  - `Once` 一次性初始化

- `examples/std/11_context.go`
  - `WithValue` 传递请求上下文
  - `WithTimeout` 自动取消
  - `WithCancel` 主动取消

- `examples/std/12_worker_pool.go`
  - 固定 worker 数量处理任务
  - 结果结构化输出
  - 正确关闭 channel 与 goroutine 生命周期

- `examples/std/13_pipeline.go`
  - pipeline stage 拆分
  - source -> filter -> transform 链式组合

### 进阶

- `examples/std/14_closure.go`
  - 状态闭包（计数器）
  - 工厂闭包（参数固化）

- `examples/std/15_generics.go`
  - 类型约束（`Number`）
  - 泛型聚合（`Sum`）
  - 泛型映射（`Map`）

- `examples/std/16_pprof.go`
  - `net/http/pprof` 入口
  - CPU/内存负载模拟
  - 运行时 goroutine 观测

---

## 常见坑速记

- Slice 截取后修改会影响原数组（共享底层存储）
- Map 不是并发安全容器（并发写需要加锁或使用 `sync.Map`）
- 关闭 channel 只能由发送方做，而且只能关闭一次
- `context.WithValue` 只放请求元信息，不要放大对象
- `defer` 参数是注册时求值，执行在函数返回前

---

## 常用运行命令

```bash
go run ./examples/std/00_basic.go
go run ./examples/std/01_error.go
go run ./examples/std/10_sync.go
go run ./examples/std/16_pprof.go
```

pprof:

```bash
go tool pprof "http://localhost:6060/debug/pprof/profile?seconds=5"
```
