# Go Routines and Channels: Complete Guide

## Table of Contents
1. [Go Routines Basics](#go-routines-basics)
2. [Channels Fundamentals](#channels-fundamentals)
3. [Unbuffered Channels](#unbuffered-channels)
4. [Buffered Channels](#buffered-channels)
5. [Common Patterns](#common-patterns)
6. [Deadlocks and Errors](#deadlocks-and-errors)
7. [Best Practices](#best-practices)

---

## Go Routines Basics

### What is a Goroutine?

A **goroutine** is a lightweight thread managed by the Go runtime. It's a function or method that runs concurrently with other goroutines.

### Starting a Goroutine

Use the `go` keyword to start a goroutine:

```go
package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    fmt.Println(msg)
}

func main() {
    // Start goroutine
    go printMessage("Hello from goroutine")
    
    // Main continues
    fmt.Println("Main thread")
    
    time.Sleep(1 * time.Second)  // Wait for goroutine to finish
}
```

**Output:**
```
Main thread
Hello from goroutine
```

### Goroutine Memory Usage

- **Initial stack size:** ~2KB
- **OS threads:** 1-2MB per thread
- **Advantage:** Can spawn thousands of goroutines efficiently

```go
// Feasible with goroutines
for i := 0; i < 100000; i++ {
    go doSomething()  // Each uses ~2KB initially
}

// Would crash with OS threads
// (would need ~100-200GB of memory!)
```

### Goroutine Lifecycle

```go
func worker(id int) {
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(1 * time.Second)
    fmt.Printf("Worker %d finished\n", id)
}

func main() {
    // Start multiple goroutines
    for i := 1; i <= 3; i++ {
        go worker(i)
    }
    
    time.Sleep(2 * time.Second)  // Wait for all to finish
}
```

---

## Channels Fundamentals

### What is a Channel?

A **channel** is a typed conduit for communication between goroutines. It allows safe data exchange.

### Channel Declaration

```go
// Declare channel
var c chan string

// Create channel
c := make(chan string)

// Buffered channel (capacity 10)
c := make(chan string, 10)

// Channel of different types
intChan := make(chan int)
boolChan := make(chan bool)
structChan := make(chan MyStruct)
```

### Send and Receive

```go
c := make(chan string)

// Send data to channel
c <- "hello"

// Receive data from channel
msg := <-c

// Receive and ignore
<-c
```

### Channel Direction

```go
// Send-only channel
var sendChan chan<- string = make(chan string)

// Receive-only channel
var recvChan <-chan string = make(chan string)

// Bidirectional channel (default)
var biChan chan string = make(chan string)
```

---

## Unbuffered Channels

### What are Unbuffered Channels?

- **Capacity:** 0
- **Behavior:** Sender blocks until receiver is ready, and vice versa
- **Synchronous:** Both sender and receiver must be ready at the same time

### Creating Unbuffered Channels

```go
c := make(chan string)  // No buffer size specified
```

### Basic Example

```go
package main

import "fmt"

func main() {
    c := make(chan string)
    
    // Spawn goroutine to receive
    go func() {
        msg := <-c
        fmt.Println("Received:", msg)
    }()
    
    // Send data
    c <- "Hello World"
}
```

**Output:**
```
Received: Hello World
```

### How Unbuffered Channels Work

```
Timeline:
1. Goroutine A: ready to receive (<-c)
2. Goroutine B: sends data (c <- "hello")
3. Both proceed simultaneously
```

### Deadlock with Unbuffered Channels

```go
func main() {
    c := make(chan string)
    c <- "Hi there!"  // ❌ DEADLOCK - no one receiving!
}
```

**Error:**
```
fatal error: all goroutines are asleep - deadlock!
```

### Correct Pattern

```go
func main() {
    c := make(chan string)
    
    // Option 1: Spawn receiver first
    go func() {
        msg := <-c
        fmt.Println(msg)
    }()
    
    c <- "hello"
    
    // Option 2: Or send in goroutine, receive in main
    go func() {
        c <- "world"
    }()
    
    fmt.Println(<-c)
}
```

---

## Buffered Channels

### What are Buffered Channels?

- **Capacity:** N (specified when created)
- **Behavior:** Sender blocks only when buffer is full
- **Asynchronous:** Sender and receiver don't need to be synchronized

### Creating Buffered Channels

```go
c := make(chan string, 5)  // Buffer size = 5
```

### Basic Example

```go
package main

import "fmt"

func main() {
    c := make(chan string, 2)  // Buffer size = 2
    
    // Send data (no receiver needed yet)
    c <- "first"   // ✅ OK
    c <- "second"  // ✅ OK
    
    // Now receive
    fmt.Println(<-c)  // "first"
    fmt.Println(<-c)  // "second"
}
```

**Output:**
```
first
second
```

### Buffer Full Behavior

```go
func main() {
    c := make(chan int, 1)
    
    c <- 1       // ✅ OK (1st send)
    c <- 2       // ❌ BLOCKS (buffer full)
    fmt.Println(<-c)  // Removes 1st element
    c <- 3       // ✅ OK now (space available)
}
```

### How Many Times Can You Send?

```go
c := make(chan string, 5)

c <- "1"  // ✅ 1st send
c <- "2"  // ✅ 2nd send
c <- "3"  // ✅ 3rd send
c <- "4"  // ✅ 4th send
c <- "5"  // ✅ 5th send
c <- "6"  // ❌ BLOCKS (buffer full, need to receive)

// Receive one
<-c       // Removes "1"
c <- "6"  // ✅ OK now
```

### Buffered vs Unbuffered Comparison

| Feature | Unbuffered | Buffered |
|---------|-----------|----------|
| Capacity | 0 | N |
| Can send without receiver? | ❌ No | ✅ Yes (up to N) |
| Synchronization | Synchronous | Asynchronous |
| Use case | Coordination | Data queuing |

---

## Common Patterns

### Pattern 1: Worker Pool

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Get results
    for r := 0; r < 9; r++ {
        fmt.Println("Result:", <-results)
    }
}
```

### Pattern 2: Fan-Out / Fan-In

```go
package main

import (
    "fmt"
    "sync"
)

func produce(out chan<- int) {
    for i := 1; i <= 5; i++ {
        out <- i
    }
    close(out)
}

func square(in <-chan int, out chan<- int) {
    for num := range in {
        out <- num * num
    }
}

func main() {
    // Fan-out: 1 producer to N consumers
    numbers := make(chan int)
    go produce(numbers)
    
    // Fan-in: N producers to 1 consumer
    var wg sync.WaitGroup
    squared := make(chan int)
    
    for i := 0; i < 2; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            square(numbers, squared)
        }()
    }
    
    go func() {
        wg.Wait()
        close(squared)
    }()
    
    // Consume results
    for result := range squared {
        fmt.Println(result)
    }
}
```

### Pattern 3: Select Statement

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
    
    // Wait for first to complete
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Received", msg1)
        case msg2 := <-c2:
            fmt.Println("Received", msg2)
        }
    }
}
```

**Output:**
```
Received one
Received two
```

### Pattern 4: Using WaitGroup

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    c := make(chan string)
    
    // Add 3 goroutines to wait group
    wg.Add(3)
    
    go func() {
        defer wg.Done()
        c <- "task 1"
    }()
    
    go func() {
        defer wg.Done()
        c <- "task 2"
    }()
    
    go func() {
        defer wg.Done()
        c <- "task 3"
    }()
    
    // Receive all results
    go func() {
        for msg := range c {
            fmt.Println(msg)
        }
    }()
    
    // Wait for all goroutines to finish
    wg.Wait()
    close(c)
}
```

---

## Deadlocks and Errors

### Deadlock 1: Unbuffered Channel, No Goroutine

```go
func main() {
    c := make(chan string)
    c <- "hello"  // ❌ DEADLOCK
}
```

**Error:**
```
fatal error: all goroutines are asleep - deadlock!
```

**Fix:**
```go
func main() {
    c := make(chan string)
    go func() {
        msg := <-c
        fmt.Println(msg)
    }()
    c <- "hello"  // ✅ OK
}
```

### Deadlock 2: No Close, Infinite Range

```go
func main() {
    c := make(chan string)
    
    for _, link := range links {
        go checkLink(link, c)  // Goroutines send data
    }
    
    for l := range c {  // ❌ DEADLOCK - range never exits
        fmt.Println(l)  // Because channel never closes
    }
}
```

**Fix:**
```go
func main() {
    var wg sync.WaitGroup
    c := make(chan string)
    
    for _, link := range links {
        wg.Add(1)
        go func(link string) {
            defer wg.Done()
            checkLink(link, c)
        }(link)
    }
    
    go func() {
        wg.Wait()
        close(c)  // ✅ Close when all done
    }()
    
    for l := range c {
        fmt.Println(l)
    }
}
```

### Deadlock 3: Sending to Closed Channel

```go
func main() {
    c := make(chan string)
    close(c)
    c <- "hello"  // ❌ PANIC: send on closed channel
}
```

**Error:**
```
panic: send on closed channel
```

### Deadlock 4: Circular Wait

```go
func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    go func() {
        c1 <- <-c2  // Waiting for c2
    }()
    
    go func() {
        c2 <- <-c1  // Waiting for c1
    }()
    
    // ❌ DEADLOCK: both waiting for each other
}
```

---

## Best Practices

### 1. Close Channels Properly

```go
// Only the sender should close
func sender(c chan<- int) {
    for i := 0; i < 5; i++ {
        c <- i
    }
    close(c)  // ✅ Sender closes
}

func main() {
    c := make(chan int)
    go sender(c)
    
    for num := range c {  // Safe to range until closed
        fmt.Println(num)
    }
}
```

### 2. Use WaitGroup for Coordination

```go
func main() {
    var wg sync.WaitGroup
    c := make(chan string)
    
    // Track number of goroutines
    wg.Add(3)
    
    for i := 0; i < 3; i++ {
        go func(id int) {
            defer wg.Done()
            c <- fmt.Sprintf("Done %d", id)
        }(i)
    }
    
    // Close when all are done
    go func() {
        wg.Wait()
        close(c)
    }()
    
    for msg := range c {
        fmt.Println(msg)
    }
}
```

### 3. Use Buffered Channels for Decoupling

```go
// Good: Buffered channel for task queue
tasks := make(chan Task, 100)

// Send tasks without waiting
for _, task := range allTasks {
    tasks <- task
}

// Workers process at their own pace
for i := 0; i < numWorkers; i++ {
    go worker(tasks)
}
```

### 4. Use Receive-Only Channels in Functions

```go
func process(c <-chan int) {
    for num := range c {
        fmt.Println(num)
    }
}

func main() {
    c := make(chan int)
    go process(c)  // process can only receive
    c <- 42
    close(c)
}
```

### 5. Use Select with Timeout

```go
func main() {
    c := make(chan string)
    
    select {
    case msg := <-c:
        fmt.Println("Received:", msg)
    case <-time.After(2 * time.Second):
        fmt.Println("Timeout!")
    }
}
```

### 6. Avoid Common Mistakes

```go
// ❌ DON'T: Send to closed channel
c := make(chan int)
close(c)
c <- 1  // PANIC

// ✅ DO: Close only after all sends complete
// Use WaitGroup to ensure completion

// ❌ DON'T: Multiple senders closing
for i := 0; i < 3; i++ {
    go func() {
        c <- i
        close(c)  // Multiple closes = PANIC
    }()
}

// ✅ DO: One closer, multiple senders
go func() {
    wg.Wait()
    close(c)
}()

// ❌ DON'T: Receive from closed channel forever
for range c {  // After close, still iterates empty values
    // ...
}

// ✅ DO: Use ok to check if closed
if val, ok := <-c; ok {
    fmt.Println(val)
} else {
    fmt.Println("Channel closed")
}
```

---

## Summary Table

| Concept | Unbuffered | Buffered |
|---------|-----------|----------|
| Syntax | `make(chan T)` | `make(chan T, N)` |
| Capacity | 0 | N |
| Send blocks? | Always until received | Only if full |
| Receive blocks? | Always until sent | Only if empty |
| Sync type | Synchronous | Asynchronous |
| Use for | Coordination | Queuing |

### Quick Checklist

- [ ] Use `go` keyword to start goroutines
- [ ] Always close channels when done
- [ ] Use WaitGroup for synchronization
- [ ] Use select for multiple channels
- [ ] Prefer buffered channels for decoupling
- [ ] Make channels receive-only in function parameters
- [ ] Handle both cases in select (timeouts, multiple channels)
- [ ] Avoid sending to closed channels
- [ ] Use `ok` to check if channel is closed

---

## References

- [Go Concurrency](https://go.dev/tour/concurrency)
- [Effective Go - Concurrency](https://golang.org/doc/effective_go#concurrency)
- [Go Memory Model](https://golang.org/ref/mem)
