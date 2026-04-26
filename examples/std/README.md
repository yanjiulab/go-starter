# examples/std

该目录是去重后的标准示例集：**每个文件只讲一个知识点**，并尽量覆盖日常常用写法。
Each example is focused but practical.

## 基础与数据结构

- `00_basic.go`：基础语法（变量、流程控制、函数、defer）
- `01_error.go`：错误处理（error、errors.Is、panic/recover）
- `02_slice.go`：切片（append、共享底层数组、copy）
- `03_map.go`：映射（读写、判断存在、删除、遍历）
- `04_string.go`：字符串与 UTF-8（bytes/runes）
- `05_struct.go`：结构体与嵌入字段
- `06_func_type.go`：函数类型与高阶函数
- `07_interface.go`：接口与多态

## 并发

- `08_goroutine.go`：goroutine 启动与同步退出
- `09_channel.go`：channel 基础收发与关闭
- `10_sync.go`：WaitGroup + Mutex 基础协作
- `11_context.go`：context 超时取消
- `12_worker_pool.go`：Worker Pool 模式
- `13_pipeline.go`：Pipeline 模式

## 进阶

- `14_closure.go`：闭包状态保持
- `15_generics.go`：泛型约束与泛型函数
- `16_pprof.go`：pprof 诊断入口
- `17_command.md`：Go 命令速查

## 使用建议

- 先运行 `00`~`07` 建立语法和抽象基础，再学习并发部分
- 每个文件都包含适量英文注释（English comments）
- 推荐边读边改：增加你的输入数据观察输出变化
