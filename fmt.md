# Go `fmt` Package

The `fmt` package implements formatted I/O with functions analogous to C's `printf` and `scanf`. It is one of the most commonly used packages in the Go standard library.

## Import

```go
import "fmt"
```

---

## Overview

The `fmt` package provides two categories of functions:

- **Print functions** — write output to stdout or a writer
- **Scan functions** — read input from stdin or a reader

---

## Format Verbs

Format verbs are placeholders used inside format strings. They begin with `%`.

### General

| Verb  | Description                          |
|-------|--------------------------------------|
| `%v`  | Default format                        |
| `%+v` | Struct fields with field names        |
| `%#v` | Go-syntax representation             |
| `%T`  | Type of the value                    |
| `%%`  | Literal percent sign                 |

### Boolean

| Verb  | Description         |
|-------|---------------------|
| `%t`  | `true` or `false`   |

### Integer

| Verb  | Description                      |
|-------|----------------------------------|
| `%d`  | Base-10 decimal                  |
| `%b`  | Base-2 binary                    |
| `%o`  | Base-8 octal                     |
| `%x`  | Base-16 hex (lowercase)          |
| `%X`  | Base-16 hex (uppercase)          |
| `%c`  | Unicode code point character     |
| `%U`  | Unicode format (`U+1234`)        |

### Floating Point

| Verb  | Description                                |
|-------|--------------------------------------------|
| `%f`  | Decimal, no exponent (`123.456`)           |
| `%e`  | Scientific notation (`1.234560e+02`)       |
| `%E`  | Scientific notation uppercase              |
| `%g`  | `%e` for large exponents, `%f` otherwise   |
| `%G`  | Same as `%g` but uppercase                 |

### String & Bytes

| Verb  | Description                            |
|-------|----------------------------------------|
| `%s`  | Unquoted string                        |
| `%q`  | Double-quoted, Go-escaped string       |
| `%x`  | Hex encoding of each byte (lowercase)  |
| `%X`  | Hex encoding of each byte (uppercase)  |

### Pointer

| Verb  | Description                    |
|-------|--------------------------------|
| `%p`  | Pointer address in hex         |

### Width and Precision

```
%[flags][width][.precision]verb
```

| Modifier | Description                         |
|----------|-------------------------------------|
| `5`      | Minimum width of 5                  |
| `.2`     | Precision of 2 decimal places       |
| `-`      | Left-align (default is right)       |
| `0`      | Pad with zeros instead of spaces    |
| `+`      | Always print sign for numbers       |
| ` `      | Space before positive numbers       |

```go
fmt.Printf("%8.2f", 3.14159)   // "    3.14"
fmt.Printf("%-8.2f", 3.14159)  // "3.14    "
fmt.Printf("%08.2f", 3.14159)  // "00003.14"
```

---

## Print Functions

### `fmt.Print`

Writes to stdout. Adds spaces between operands when neither is a string.

```go
fmt.Print("Hello", " ", "World")  // Hello World
fmt.Print(1, 2, 3)                // 1 2 3
```

### `fmt.Println`

Writes to stdout with spaces between operands and a trailing newline.

```go
fmt.Println("Hello", "World")  // Hello World\n
fmt.Println(1, 2, 3)           // 1 2 3\n
```

### `fmt.Printf`

Writes formatted output to stdout.

```go
name := "Alice"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Name: Alice, Age: 30
```

---

## Sprint Functions (Return Strings)

### `fmt.Sprint`

Returns a formatted string. Adds spaces between operands when neither is a string.

```go
s := fmt.Sprint("Hello", " ", "World")
// s == "Hello World"
```

### `fmt.Sprintln`

Returns a string with spaces between operands and a trailing newline.

```go
s := fmt.Sprintln("Hello", "World")
// s == "Hello World\n"
```

### `fmt.Sprintf`

Returns a formatted string using a format specifier.

```go
s := fmt.Sprintf("Pi is approximately %.4f", 3.14159)
// s == "Pi is approximately 3.1416"
```

---

## Fprint Functions (Write to io.Writer)

These write to any `io.Writer`, not just stdout.

### `fmt.Fprint`

```go
fmt.Fprint(os.Stderr, "an error occurred")
```

### `fmt.Fprintln`

```go
fmt.Fprintln(os.Stdout, "line output")
```

### `fmt.Fprintf`

```go
fmt.Fprintf(os.Stderr, "error: %v\n", err)
```

---

## Errorf

Creates a formatted error value.

```go
err := fmt.Errorf("file %s not found", filename)
```

### Wrapping Errors with `%w`

Since Go 1.13, use `%w` to wrap errors for use with `errors.Is` and `errors.As`.

```go
originalErr := errors.New("connection refused")
wrappedErr := fmt.Errorf("database error: %w", originalErr)

errors.Is(wrappedErr, originalErr) // true
```

---

## Scan Functions

### `fmt.Scan`

Reads space-separated values from stdin.

```go
var name string
var age int
fmt.Scan(&name, &age)
```

### `fmt.Scanln`

Like `fmt.Scan` but stops at a newline.

```go
var s string
fmt.Scanln(&s)
```

### `fmt.Scanf`

Reads input using a format string.

```go
var name string
var age int
fmt.Scanf("%s %d", &name, &age)
```

### `fmt.Sscan` / `fmt.Sscanf` / `fmt.Sscanln`

Read from a string rather than stdin.

```go
var x, y int
fmt.Sscan("10 20", &x, &y)
// x == 10, y == 20

fmt.Sscanf("Name: Alice", "Name: %s", &name)
```

### `fmt.Fscan` / `fmt.Fscanf` / `fmt.Fscanln`

Read from any `io.Reader`.

```go
var val int
fmt.Fscan(reader, &val)
```

---

## The `Stringer` Interface

If a type implements the `Stringer` interface, `fmt` will use its `String()` method automatically with `%v` and `%s`.

```go
type Point struct {
    X, Y int
}

func (p Point) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

p := Point{3, 4}
fmt.Println(p)         // (3, 4)
fmt.Printf("%v\n", p) // (3, 4)
```

---

## The `GoStringer` Interface

Implement `GoString()` to control the output of `%#v`.

```go
func (p Point) GoString() string {
    return fmt.Sprintf("Point{X:%d, Y:%d}", p.X, p.Y)
}

fmt.Printf("%#v\n", p) // Point{X:3, Y:4}
```

---

## The `Formatter` Interface

For full control over formatting, implement `fmt.Formatter`:

```go
type MyType struct{}

func (m MyType) Format(f fmt.State, verb rune) {
    switch verb {
    case 'v', 's':
        fmt.Fprint(f, "MyType")
    default:
        fmt.Fprintf(f, "%%!%c(MyType)", verb)
    }
}
```

---

## Common Patterns and Examples

### Printing Struct Values

```go
type User struct {
    Name string
    Age  int
}

u := User{"Bob", 25}
fmt.Printf("%v\n", u)   // {Bob 25}
fmt.Printf("%+v\n", u)  // {Name:Bob Age:25}
fmt.Printf("%#v\n", u)  // main.User{Name:"Bob", Age:25}
```

### Padding and Alignment

```go
fmt.Printf("|%10s|\n", "right")    // |     right|
fmt.Printf("|%-10s|\n", "left")    // |left      |
fmt.Printf("|%10d|\n", 42)         // |        42|
fmt.Printf("|%-10d|\n", 42)        // |42        |
```

### Printing Multiple Types

```go
fmt.Printf("%T\n", 42)          // int
fmt.Printf("%T\n", 3.14)        // float64
fmt.Printf("%T\n", "hello")     // string
fmt.Printf("%T\n", []int{1, 2}) // []int
```

### Hex Dump

```go
data := []byte("hello")
fmt.Printf("%x\n", data)   // 68656c6c6f
fmt.Printf("% x\n", data)  // 68 65 6c 6c 6f
fmt.Printf("%X\n", data)   // 68656C6C6F
```

---

## Summary Table

| Function       | Destination  | Format String | Returns    |
|----------------|--------------|---------------|------------|
| `Print`        | stdout       | No            | n, err     |
| `Println`      | stdout       | No            | n, err     |
| `Printf`       | stdout       | Yes           | n, err     |
| `Sprint`       | string       | No            | string     |
| `Sprintln`     | string       | No            | string     |
| `Sprintf`      | string       | Yes           | string     |
| `Fprint`       | io.Writer    | No            | n, err     |
| `Fprintln`     | io.Writer    | No            | n, err     |
| `Fprintf`      | io.Writer    | Yes           | n, err     |
| `Errorf`       | error value  | Yes           | error      |
| `Scan`         | stdin        | No            | n, err     |
| `Scanf`        | stdin        | Yes           | n, err     |
| `Scanln`       | stdin        | No            | n, err     |
| `Sscan`        | string       | No            | n, err     |
| `Sscanf`       | string       | Yes           | n, err     |
| `Fscan`        | io.Reader    | No            | n, err     |
| `Fscanf`       | io.Reader    | Yes           | n, err     |

---

## References

- [Official Go Documentation — fmt package](https://pkg.go.dev/fmt)
- [Go by Example — String Formatting](https://gobyexample.com/string-formatting)
- [Effective Go](https://go.dev/doc/effective_go)
