# OOP

## Struct

Structs are composite data types that group together variables under a single name.

### Defining Structs

Structs are defined using the `type` keyword.

```go
type Person struct {
    Name string
    Age  int
}
```

### Struct Fields

Fields can be accessed using dot notation.

```go
p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name) // Alice
```

## Method

Methods are functions associated with a particular type.

### Defining Methods

Methods are defined with a receiver.

```go
func (p Person) greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

### Value vs Pointer Receivers

Value receivers work on copies, pointer receivers modify the original.

```go
func (p *Person) setAge(age int) {
    p.Age = age
}
```

## Interface

Interfaces define a set of method signatures.

### Defining Interfaces

Interfaces are defined with method signatures.

```go
type Speaker interface {
    Speak() string
}
```

### Implementing Interfaces

Types implement interfaces by having the required methods.

```go
func (p Person) Speak() string {
    return "Hello!"
}
```

## Composition

Go uses composition instead of inheritance.

### Embedding Structs

Embed structs to reuse fields and methods.

```go
type Employee struct {
    Person
    Salary int
}
```

## Value Passing & Reference Passing

Go passes values by default, but pointers allow reference passing.

### Pass by Value

Functions receive copies of arguments.

```go
func modify(x int) {
    x = 10 // Doesn't affect original
}
```

### Pass by Reference

Use pointers to modify originals.

```go
func modifyPtr(x *int) {
    *x = 10
}
```

## Type Assertion

Type assertions extract the underlying value from an interface.

### Basic Type Assertion

```go
var i interface{} = "hello"
s := i.(string)       // "hello"
s, ok := i.(string)   // "hello", true
```

### Type Switch

Type switch handles multiple types.

```go
switch v := i.(type) {
case string:
    fmt.Println("String:", v)
case int:
    fmt.Println("Int:", v)
default:
    fmt.Println("Unknown")
}
```

## Struct Tags

Struct tags provide metadata for fields.

### Using Tags

Common tags for JSON marshaling and field behavior.

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
    Age   int    `json:"-"`
}
```

## Empty Interface

The `interface{}` type can hold any value.

```go
var x interface{}
x = "hello"
x = 42
x = []int{1, 2, 3}
```

## Reflection

Reflection allows inspection of types at runtime.

### Using reflect Package

The `reflect` package provides reflection capabilities.

```go
t := reflect.TypeOf(p)
fmt.Println(t.Name())
```

## Generics

Generics allow writing flexible, reusable code.

### Generic Types

Define types with type parameters.

```go
type Stack[T any] struct {
    items []T
}
```

### Generic Functions

Functions can be generic.

```go
func add[T any](a, b T) T {
    // Implementation depends on T
}
```
