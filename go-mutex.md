# Go Mutex (Mutual Exclusion)

## What is a Mutex?

A Mutex (Mutual Exclusion) is a synchronization mechanism that ensures only one goroutine can access a shared resource at a time.

It is used to prevent **race conditions** when multiple goroutines read and write the same data.

Mutex is provided by the `sync` package.

```go
var mu sync.Mutex
```

---

## Why Do We Need Mutex?

Imagine multiple goroutines incrementing the same counter.

```go
counter++
```

This looks like a single operation but internally it performs:

```go
temp := counter // Read
temp = temp + 1 // Modify
counter = temp  // Write
```

If two goroutines execute these steps simultaneously, updates can be lost.

Example:

```text
counter = 5

Goroutine A reads 5
Goroutine B reads 5

A writes 6
B writes 6

Expected: 7
Actual: 6
```

This problem is called a **Race Condition**.

---

# Example Without Mutex

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
```

Possible output:

```text
941
972
998
```

The output is unpredictable because multiple goroutines modify the same variable simultaneously.

---

# Example With Mutex

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
```

Output:

```text
1000
```

---

# How Mutex Works

```go
mu.Lock()
counter++
mu.Unlock()
```

### Step 1

```go
mu.Lock()
```

Acquire the lock.

If another goroutine already owns the lock, the current goroutine waits.

### Step 2

```go
counter++
```

Execute the critical section.

### Step 3

```go
mu.Unlock()
```

Release the lock.

Waiting goroutines can now acquire it.

---

# Visualization

Without Mutex

```text
G1 ---> counter++
G2 ---> counter++
G3 ---> counter++

All run simultaneously
```

With Mutex

```text
G1 ---> Lock
      counter++
      Unlock

G2 waits

G2 ---> Lock
      counter++
      Unlock

G3 waits

G3 ---> Lock
      counter++
      Unlock
```

Only one goroutine can enter the critical section at a time.

---

# Critical Section

The code protected by a mutex is called the critical section.

```go
mu.Lock()

counter++

mu.Unlock()
```

Keep the critical section as small as possible.

Bad:

```go
mu.Lock()

time.Sleep(5 * time.Second)

counter++

mu.Unlock()
```

All other goroutines must wait 5 seconds.

Good:

```go
time.Sleep(5 * time.Second)

mu.Lock()
counter++
mu.Unlock()
```

---

# Using defer with Mutex

A common pattern:

```go
mu.Lock()
defer mu.Unlock()

counter++
```

Benefits:

- Cleaner code
- Prevents forgotten unlocks
- Safe during early returns

Example:

```go
func increment() {
	mu.Lock()
	defer mu.Unlock()

	counter++
}
```

---

# Protecting a Shared Map

Maps are not safe for concurrent writes.

Wrong:

```go
m := map[string]int{}

go func() {
	m["a"] = 1
}()

go func() {
	m["b"] = 2
}()
```

Can cause:

```text
fatal error: concurrent map writes
```

Correct:

```go
var mu sync.Mutex

m := map[string]int{}

go func() {
	mu.Lock()
	m["a"] = 1
	mu.Unlock()
}()

go func() {
	mu.Lock()
	m["b"] = 2
	mu.Unlock()
}()
```

---

# RWMutex

Go provides a special mutex:

```go
var mu sync.RWMutex
```

RWMutex supports:

- Multiple readers
- Single writer

---

## Read Lock

```go
mu.RLock()

value := cache["user"]

mu.RUnlock()
```

Many readers can run simultaneously.

---

## Write Lock

```go
mu.Lock()

cache["user"] = "Ajay"

mu.Unlock()
```

Only one writer can run.

During writing:

- Readers wait
- Writers wait

---

# RWMutex Example

```go
var (
	cache = map[string]string{}
	mu    sync.RWMutex
)

func Get(key string) string {
	mu.RLock()
	defer mu.RUnlock()

	return cache[key]
}

func Set(key, value string) {
	mu.Lock()
	defer mu.Unlock()

	cache[key] = value
}
```

---

# Mutex vs Channel

## Mutex

Use when protecting shared state.

Examples:

- Counters
- Maps
- Cache
- Shared structs

```go
mu.Lock()
counter++
mu.Unlock()
```

---

## Channel

Use when passing data between goroutines.

```go
ch <- data
```

Examples:

- Worker pools
- Pipelines
- Notifications
- Event processing

---

# Mutex Example with Struct

```go
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}
```

---

# Common Mistakes

## Forgetting Unlock

Wrong:

```go
mu.Lock()

counter++

return
```

Lock is never released.

Better:

```go
mu.Lock()
defer mu.Unlock()

counter++
```

---

## Locking Too Much Code

Wrong:

```go
mu.Lock()

time.Sleep(10 * time.Second)

counter++

mu.Unlock()
```

This blocks everyone.

---

## Double Unlock

Wrong:

```go
mu.Unlock()
```

without a corresponding Lock.

Result:

```text
fatal error: sync: unlock of unlocked mutex
```

---

# Race Detector

Go provides a race detector.

Run:

```bash
go run -race main.go
```

or

```bash
go test -race
```

Example output:

```text
WARNING: DATA RACE
```

Very useful for finding concurrency bugs.

---

# Summary

| Feature | Mutex |
|----------|--------|
| Purpose | Protect shared data |
| Package | sync |
| Lock Method | Lock() |
| Unlock Method | Unlock() |
| Multiple Readers | No |
| Multiple Writers | No |
| Read Optimization | Use RWMutex |
| Prevents Race Conditions | Yes |

---

# Rule of Thumb

Use a Mutex when:

- Multiple goroutines access the same data
- At least one goroutine modifies the data

Use Channels when:

- Goroutines need to communicate
- Data should flow through a pipeline
- Building worker pools

Remember:

> A mutex protects shared memory.
>
> A channel transfers ownership of data between goroutines.
