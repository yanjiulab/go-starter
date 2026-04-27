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
