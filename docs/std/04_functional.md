# Functional Programming

## Functions as First-Class Citizens

Functions can be assigned to variables, passed as arguments, and returned from other functions.

### Function Types

Define function types.

```go
type Operation func(int, int) int
```

### Assigning Functions

Assign functions to variables.

```go
add := func(a, b int) int { return a + b }
```

## Anonymous Functions

Functions without names, often used inline.

### Defining Anonymous Functions

Anonymous functions are defined with `func`.

```go
func() {
    fmt.Println("Anonymous")
}()
```

### Closures

Functions that capture variables from their environment.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Higher-Order Functions

Functions that take functions as arguments or return functions.

### Map Function

Apply a function to each element.

```go
func mapFunc(f func(int) int, slice []int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}
```

### Filter Function

Filter elements based on a predicate.

```go
func filter(f func(int) bool, slice []int) []int {
    var result []int
    for _, v := range slice {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}
```

## Function Composition and Currying

Combining functions and transforming multi-argument functions.

### Composition

Chain functions together.

```go
func compose(f, g func(int) int) func(int) int {
    return func(x int) int {
        return f(g(x))
    }
}
```

### Currying

Transform a function to take arguments one at a time.

```go
func add(a int) func(int) int {
    return func(b int) int {
        return a + b
    }
}
```

## Defer

Defer postpones function execution until the surrounding function returns.

### Basic Defer

Defer statements run in LIFO order.

```go
func example() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    fmt.Println("Main")
}
```

### Use Cases

Commonly used for cleanup.

```go
file, _ := os.Open("file.txt")
defer file.Close()
```

## Recursion

Functions can call themselves for divide-and-conquer problems.

### Basic Recursion

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}
```

### Stack Overflow Prevention

Be mindful of call depth for large datasets.

## Generic Functions

Functions that work with any type using generics.

### Defining Generic Functions

Use type parameters.

```go
func identity[T any](x T) T {
    return x
}
```

### Constraints

Constrain type parameters.

```go
func max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

## Error Handling in Functions

Use defer, panic, and recover for robust error handling.

### Defer for Cleanup

Ensure resources are released.

```go
func readFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    return os.ReadFile(path)
}
```

### Panic and Recover

Handle exceptional situations.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered from:", r)
    }
}()
if impossible {
    panic("This should not happen")
}
```
