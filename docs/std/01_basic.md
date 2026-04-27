# Basic

## Hello World

This section covers the fundamentals of writing and running a simple Go program, including compilation and comments.

### Basic Compilation

To compile a Go program, use the `go build` command. For example, if you have a file named `main.go`, run `go build main.go` to create an executable. You can also use `go run main.go` to compile and run the program in one step.

### Comments

Comments in Go are used to add explanatory notes to the code. Single-line comments start with `//`, and multi-line comments are enclosed in `/* */`.

```go
// This is a single-line comment
/*
This is a multi-line comment
*/
```

## Constants & Variables

Constants and variables are fundamental concepts in Go for storing data.

### Constants

Constants are immutable values declared with the `const` keyword. They must be assigned a value at declaration and cannot be changed later.

```go
const Pi = 3.14159
const Greeting = "Hello, World!"
```

### Iota

Use `iota` to create sequences of constants.

```go
const (
    A = iota // 0
    B = iota // 1
    C = iota // 2
)
```

### Variables

Variables are declared using the `var` keyword or short variable declaration `:=`. They can be reassigned.

```go
var name string = "Go"
age := 10
```

## Basic Types & Built-in Functions

Go has several basic types and useful built-in functions for common operations.

### Basic Types

Common basic types include `int`, `float64`, `string`, `bool`, and `rune`. The zero value is the default value for a type when a variable is declared without initialization.

```go
var x int        // 0
var s string     // ""
var b bool       // false
var by byte      // 0 (same as uint8)
var r rune       // 0 (same as int32, represents Unicode)
```

### Named Types

Create custom types based on existing types.

```go
type Meter float64
type Point struct { X, Y int }
```

### Type Conversion

Convert between compatible types explicitly.

```go
var i int = 42
var f float64 = float64(i)
```

### Built-in Functions

Common built-in functions include `len`, `cap`, `make`, `new`, `append`, `copy`, and `delete`.

```go
s := make([]int, 0, 10)
len(s)
cap(s)
```

## Flow Control

Flow control structures manage the execution path of a program.

### If Statements

The `if` statement executes code based on a condition.

```go
if x > 0 {
    fmt.Println("Positive")
}

if age >= 3 {
    fmt.Println(name, "is growing with Go")
} else {
    fmt.Println(name, "is too young")
}
```

### Loops

Go uses `for` loops for iteration.

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### Range Loops

Use `range` to iterate over arrays, slices, maps, and strings.

```go
for i, v := range slice {
    fmt.Println(i, v)
}
```

### Switch Statements

Switch provides multi-way branching.

```go
switch day {
case "Monday":
    fmt.Println("Start of work week")
default:
    fmt.Println("Other day")
}
```

## Error Handling

Go handles errors explicitly using the `error` type.

### Returning Errors

Functions can return errors as their last return value.

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Handling Errors

Use `if` to check for errors.

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
}
```

## Function

Functions are blocks of code that perform specific tasks.

### Defining Functions

Functions are defined with the `func` keyword.

```go
func add(a, b int) int {
    return a + b
}
```

### Function Parameters and Return Values

Functions can take parameters and return values.

```go
func greet(name string) string {
    return "Hello, " + name
}
```

## Pointer

Pointers hold the memory address of a value.

### Declaring Pointers

Use `*` to declare a pointer type.

```go
var p *int
```

### Dereferencing

Use `*` to access the value at the pointer.

```go
x := 10
p = &x
fmt.Println(*p) // 10
```

### Nil Pointers

Pointers have a zero value of `nil`.

```go
var p *int // nil
if p != nil {
    fmt.Println(*p)
}
```

## Data Structure

Go provides built-in data structures like arrays, slices, and maps.

### Arrays

Fixed-size sequences of elements.

```go
var arr [5]int
arr[0] = 1
fmt.Println(arr[0])
```

### Slices

Dynamic arrays that can grow and shrink.

#### Create

```go
slice := []int{1, 2, 3}
```

#### Read

```go
fmt.Println(slice[1])
```

#### Update

```go
slice[0] = 10
```

#### Delete

Use slicing to remove an element.

```go
slice = append(slice[:1], slice[2:]...)
```

### Maps

Key-value pairs with dynamic resizing.

#### Create

```go
m := make(map[string]int)
```

#### Read

```go
value := m["key"]
```

#### Update

```go
m["key"] = 42
```

#### Delete

```go
delete(m, "key")
```

#### Existence Check

```go
value, ok := m["key"]
if ok {
    fmt.Println("exists", value)
}
```

## String Operations

The `strings` package provides common string operations.

### Common String Functions

```go
import "strings"

strings.Contains("hello", "ell")     // true
strings.ToUpper("hello")             // "HELLO"
strings.Split("a,b,c", ",")          // ["a" "b" "c"]
strings.Join([]string{"a", "b"}, "-") // "a-b"
strings.HasPrefix("hello", "he")     // true
```

## Panic & Recover

Handle unexpected errors or exceptional situations.

### Panic

Panic terminates the program flow.

```go
if x < 0 {
    panic("x must be positive")
}
```

### Recover

Recover catches panics and allows continued execution.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
panic("something went wrong")
```

## OOP

Go supports object-oriented programming through structs and methods.

### Structs

Define custom types.

```go
type Person struct {
    Name string
    Age  int
}
```

### Methods

Functions attached to types.

```go
func (p Person) greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

## Goroutine

Goroutines are lightweight threads for concurrent execution.

### Starting Goroutines

Use `go` keyword to start a goroutine.

```go
go func() {
    fmt.Println("Hello from goroutine")
}()
```

### Synchronization

Use channels or sync package for synchronization.

```go
ch := make(chan int)
go func() { ch <- 42 }()
value := <-ch
```
