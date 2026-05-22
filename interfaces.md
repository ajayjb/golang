# Interfaces in Go

## Introduction

Interfaces are one of the most powerful features in Go. They define a set of methods that a type must implement, enabling polymorphism and flexible, loosely-coupled code. Unlike languages like Java, Go uses implicit interface implementation — a type doesn't need to explicitly declare that it implements an interface.

## Basic Syntax

### Defining an Interface

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

An interface is defined using the `type` keyword followed by the interface name and the `interface` keyword. It contains method signatures (name, parameters, and return types) but no implementation.

### Empty Interface

The empty interface `interface{}` specifies no methods and can hold values of any type:

```go
var x interface{} = "hello"
var y interface{} = 42
var z interface{} = struct{}{}
```

In Go 1.18+, `any` is an alias for `interface{}`:

```go
var x any = "hello"
```

## Implementing Interfaces

In Go, a type implements an interface by implementing all of its methods. There's no explicit "implements" keyword.

### Example

```go
package main

import "fmt"

type Writer interface {
    Write(p []byte) (n int, err error)
}

type FileWriter struct {
    filename string
}

func (fw *FileWriter) Write(p []byte) (int, error) {
    fmt.Printf("Writing to %s: %s\n", fw.filename, string(p))
    return len(p), nil
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(p []byte) (int, error) {
    fmt.Printf("Console: %s\n", string(p))
    return len(p), nil
}

func main() {
    var w Writer
    
    w = &FileWriter{filename: "output.txt"}
    w.Write([]byte("Hello from file"))
    
    w = ConsoleWriter{}
    w.Write([]byte("Hello from console"))
}
```

**Key Points:**
- Methods must have the same signature as defined in the interface
- The receiver can be a pointer or a value
- A type automatically implements all interfaces it has methods for

## Method Receivers

When implementing interface methods, the choice between value and pointer receivers matters:

```go
type Speaker interface {
    Speak() string
}

type Dog struct {
    name string
}

// Value receiver - implements Speaker for Dog values
func (d Dog) Speak() string {
    return d.name + " says Woof!"
}

type Cat struct {
    name string
}

// Pointer receiver - implements Speaker for *Cat pointers
func (c *Cat) Speak() string {
    return c.name + " says Meow!"
}

func main() {
    var s Speaker
    
    // Works: Dog value implements Speaker
    s = Dog{name: "Buddy"}
    fmt.Println(s.Speak()) // Buddy says Woof!
    
    // Works: *Cat pointer implements Speaker
    s = &Cat{name: "Whiskers"}
    fmt.Println(s.Speak()) // Whiskers says Meow!
}
```

## Type Assertions

Type assertions extract the underlying concrete value from an interface:

```go
var i interface{} = "hello"

// Extract as string
s := i.(string)
fmt.Println(s) // hello

// Type assertion with error check (recommended)
s, ok := i.(string)
if ok {
    fmt.Println("String:", s)
} else {
    fmt.Println("Not a string")
}

// Panics if i is not a string
// s := i.(string)
```

### Type Switches

For multiple type assertions, use a type switch:

```go
func PrintType(i interface{}) {
    switch v := i.(type) {
    case string:
        fmt.Printf("String: %s\n", v)
    case int:
        fmt.Printf("Integer: %d\n", v)
    case bool:
        fmt.Printf("Boolean: %v\n", v)
    case nil:
        fmt.Println("Nil")
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

func main() {
    PrintType("hello")      // String: hello
    PrintType(42)           // Integer: 42
    PrintType(true)         // Boolean: true
    PrintType(3.14)         // Unknown type: float64
}
```

## Interface Composition

Interfaces can embed other interfaces:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

type BufferedReadWriter struct{}

func (brw BufferedReadWriter) Read(p []byte) (int, error) {
    return len(p), nil
}

func (brw BufferedReadWriter) Write(p []byte) (int, error) {
    return len(p), nil
}

func main() {
    var rw ReadWriter = BufferedReadWriter{}
    _ = rw
}
```

## Common Interface Patterns

### Stringer Interface

The `Stringer` interface provides custom string representation:

```go
type Stringer interface {
    String() string
}

type Person struct {
    name string
    age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d years old)", p.name, p.age)
}

func main() {
    p := Person{name: "Alice", age: 30}
    fmt.Println(p) // Alice (30 years old)
}
```

### Error Interface

The `error` interface is fundamental in Go:

```go
type error interface {
    Error() string
}

type CustomError struct {
    message string
    code    int
}

func (ce CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", ce.code, ce.message)
}

func doSomething() error {
    return CustomError{message: "Something went wrong", code: 500}
}

func main() {
    if err := doSomething(); err != nil {
        fmt.Println(err) // Error 500: Something went wrong
    }
}
```

### io.Reader and io.Writer

Standard library interfaces for I/O:

```go
import (
    "io"
    "strings"
)

func CountWords(r io.Reader) (int, error) {
    // Process any Reader
    data := make([]byte, 1024)
    n, _ := r.Read(data)
    return len(strings.Fields(string(data[:n]))), nil
}

func main() {
    reader := strings.NewReader("hello world go programming")
    count, _ := CountWords(reader)
    fmt.Println(count) // 4
}
```

## Interface Segregation Principle

Keep interfaces small and focused:

```go
// Good: Small, focused interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

// Compose them when needed
type ReadCloser interface {
    Reader
    Closer
}

// Bad: Large interface with many unrelated methods
type DataProcessor interface {
    Read(p []byte) (n int, err error)
    Write(p []byte) (n int, err error)
    Close() error
    Compress()
    Encrypt()
    Upload()
}
```

## Nil Interfaces

An interface value is `nil` only if both its type and value are `nil`:

```go
var i interface{} = nil
fmt.Println(i == nil) // true

var p *int = nil
i = p
fmt.Println(i == nil) // false! Type is *int, even though value is nil

// Check properly
if i != nil {
    // Type assertion
    _ = i
}
```

## Practical Example: Plugin System

```go
package main

import "fmt"

// Plugin interface
type Plugin interface {
    Name() string
    Execute(input string) string
}

// Plugin implementations
type UppercasePlugin struct{}

func (up UppercasePlugin) Name() string {
    return "Uppercase"
}

func (up UppercasePlugin) Execute(input string) string {
    return strings.ToUpper(input)
}

type ReversePlugin struct{}

func (rp ReversePlugin) Name() string {
    return "Reverse"
}

func (rp ReversePlugin) Execute(input string) string {
    runes := []rune(input)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// Manager uses plugins
type PluginManager struct {
    plugins map[string]Plugin
}

func (pm *PluginManager) Register(plugin Plugin) {
    pm.plugins[plugin.Name()] = plugin
}

func (pm *PluginManager) Execute(name string, input string) (string, error) {
    plugin, exists := pm.plugins[name]
    if !exists {
        return "", fmt.Errorf("plugin %s not found", name)
    }
    return plugin.Execute(input), nil
}

func main() {
    manager := &PluginManager{plugins: make(map[string]Plugin)}
    
    manager.Register(UppercasePlugin{})
    manager.Register(ReversePlugin{})
    
    result, _ := manager.Execute("Uppercase", "hello")
    fmt.Println(result) // HELLO
    
    result, _ = manager.Execute("Reverse", "hello")
    fmt.Println(result) // olleh
}
```

## Best Practices

1. **Define interfaces near their point of use** — Don't define large interfaces upfront
2. **Accept interfaces, return concrete types** — Flexible inputs, predictable outputs
3. **Keep interfaces small** — Ideally 1-3 methods
4. **Use empty interface sparingly** — It bypasses type safety
5. **Document interface behavior** — Not just method signatures
6. **Check for nil carefully** — Remember nil type vs nil value
7. **Leverage composition** — Combine small interfaces for flexibility

## Summary

Interfaces in Go enable clean, flexible code through implicit implementation and composition. They're fundamental to writing idiomatic Go and are used extensively in the standard library. Understanding interfaces is key to mastering Go.
