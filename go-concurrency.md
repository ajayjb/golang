# Go Concurrency: Complete Beginner to Advanced Guide

## Table of Contents

1. Concurrency vs Parallelism
2. Goroutines
3. Channels
4. Buffered vs Unbuffered Channels
5. Channel Lifecycle
6. WaitGroup
7. Select
8. Semaphore
9. Worker Pool Pattern
10. Pipeline Pattern
11. Context Cancellation
12. Common Mistakes
13. Mental Models
14. Interview Questions

---

# 1. Concurrency vs Parallelism

Many beginners confuse these concepts.

## Concurrency

Concurrency means:

> Multiple tasks are making progress during the same period.

Example:

```text
Task A
Task B
Task A
Task B
Task A
Task B
```

Only one CPU core may be involved.

The runtime switches between tasks.

---

## Parallelism

Parallelism means:

> Multiple tasks are executing at exactly the same time.

Example:

```text
CPU Core 1 -> Task A
CPU Core 2 -> Task B
```

---

## Real World Example

Imagine a chef.

### Concurrency

One chef:

```text
Cut vegetables
Boil water
Check oven
Mix ingredients
```

Switching between tasks.

---

### Parallelism

Four chefs:

```text
Chef 1 -> Cut vegetables
Chef 2 -> Prepare sauce
Chef 3 -> Bake bread
Chef 4 -> Prepare dessert
```

Working simultaneously.

---

## In Go

```go
go task1()
go task2()
```

creates concurrency.

The Go scheduler decides whether they run in parallel.

---

# 2. Goroutines

A goroutine is a lightweight thread managed by Go.

## Creating a Goroutine

```go
go doWork()
```

Example:

```go
package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func main() {
	go hello()

	fmt.Println("Main")
}
```

Possible output:

```text
Main
```

Why?

Because main exits before hello runs.

---

## Waiting

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Worker")
	}()

	time.Sleep(time.Second)
}
```

Output:

```text
Worker
```

---

## Why Goroutines Are Cheap

OS Thread:

```text
~1 MB stack
```

Goroutine:

```text
~2 KB initial stack
```

Therefore:

```text
Thousands or millions possible
```

---

# 3. Channels

Channels allow goroutines to communicate safely.

Think:

```text
Producer --> Channel --> Consumer
```

---

## Creating

```go
ch := make(chan int)
```

---

## Sending

```go
ch <- 10
```

---

## Receiving

```go
value := <-ch
```

---

## Example

```go
package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello"
	}()

	fmt.Println(<-ch)
}
```

Output:

```text
hello
```

---

# 4. Unbuffered Channels

```go
ch := make(chan int)
```

No buffer.

---

## Rule

Sender waits for receiver.

Receiver waits for sender.

Example:

```go
go func() {
	ch <- 10
}()
```

Execution:

```text
Worker wants to send

No receiver yet

Worker blocks
```

---

Main:

```go
fmt.Println(<-ch)
```

Now:

```text
Receiver ready

Value transferred

Both continue
```

---

## Mental Model

Imagine a handshake.

```text
Sender ----- Receiver
```

Both must meet.

---

# 5. Buffered Channels

```go
ch := make(chan int, 3)
```

Capacity = 3.

---

## Example

```go
ch <- 1
ch <- 2
ch <- 3
```

Buffer:

```text
[1][2][3]
```

Full.

Next send:

```go
ch <- 4
```

blocks.

---

## Why Use Buffered Channels?

Useful for:

* Job queues
* Worker pools
* Event processing

---

# 6. Channel Lifecycle

## Creation

```go
ch := make(chan int)
```

---

## Send

```go
ch <- 10
```

---

## Receive

```go
v := <-ch
```

---

## Close

```go
close(ch)
```

---

# Important Rule

Only the sender should close the channel.

Bad:

```go
Receiver closes channel
```

May panic.

---

# 7. Reading Until Closed

```go
for v := range ch {
	fmt.Println(v)
}
```

Equivalent to:

```go
for {
	v, ok := <-ch

	if !ok {
		break
	}

	fmt.Println(v)
}
```

---

## Example

```go
go func() {
	ch <- 1
	ch <- 2
	close(ch)
}()
```

Output:

```text
1
2
```

Loop exits automatically.

---

# 8. WaitGroup

WaitGroup waits for goroutines to finish.

Think:

```text
Counter
```

---

## Add

```go
wg.Add(1)
```

Counter:

```text
0 -> 1
```

---

## Done

```go
wg.Done()
```

Counter:

```text
1 -> 0
```

---

## Wait

```go
wg.Wait()
```

Blocks until:

```text
Counter == 0
```

---

## Example

```go
var wg sync.WaitGroup

wg.Add(2)

go func() {
	defer wg.Done()
}()

go func() {
	defer wg.Done()
}()

wg.Wait()
```

---

## WaitGroup.Go (Go 1.25+)

```go
wg.Go(func() {
	doWork()
})
```

Equivalent:

```go
wg.Add(1)

go func() {
	defer wg.Done()
	doWork()
}()
```

---

# 9. Select

Select waits on multiple channel operations.

Think:

```text
switch for channels
```

---

## Example

```go
select {
case msg := <-ch1:
	fmt.Println(msg)

case msg := <-ch2:
	fmt.Println(msg)
}
```

Whichever channel becomes ready first wins.

---

## Timeout Example

```go
select {
case msg := <-ch:
	fmt.Println(msg)

case <-time.After(5 * time.Second):
	fmt.Println("timeout")
}
```

---

## Context Example

```go
select {
case <-ctx.Done():
	return

case msg := <-ch:
	process(msg)
}
```

---

# 10. Semaphore

Semaphore limits concurrency.

Think:

```text
Parking lot
```

Capacity:

```text
4 parking spots
```

Only four cars can enter.

---

## Go Semaphore

```go
sem := make(chan struct{}, 4)
```

---

## Acquire

```go
sem <- struct{}{}
```

Occupy a slot.

---

## Release

```go
<-sem
```

Free a slot.

---

## Example

```go
sem := make(chan struct{}, 2)

sem <- struct{}{}
sem <- struct{}{}
```

Buffer:

```text
[X][X]
```

Full.

Another send blocks.

---

# 11. Worker Pool Pattern

Problem:

```go
for _, job := range jobs {
	go process(job)
}
```

100,000 jobs:

```text
100,000 goroutines
```

Bad.

---

## Solution

Use a semaphore.

```go
sem := make(chan struct{}, 4)
```

Only four jobs process simultaneously.

---

## Complete Example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := []int{1,2,3,4,5,6,7,8}

	const workers = 4

	sem := make(chan struct{}, workers)

	var wg sync.WaitGroup

	for _, job := range jobs {

		sem <- struct{}{}

		wg.Go(func() {
			defer func() {
				<-sem
			}()

			fmt.Printf("Starting job %d\n", job)

			time.Sleep(2 * time.Second)

			fmt.Printf("Finished job %d\n", job)
		})
	}

	wg.Wait()

	fmt.Println("All jobs complete")
}
```

---

## Execution Timeline

```text
Time 0

Job1
Job2
Job3
Job4

Semaphore full

Job5 waits
Job6 waits
Job7 waits
Job8 waits
```

---

```text
Time 2

Job2 finishes

Slot released

Job5 starts
```

---

```text
Time 4

Job1 finishes

Slot released

Job6 starts
```

---

Eventually:

```text
All jobs finish

WaitGroup counter becomes 0

wg.Wait() returns
```

---

# 12. Pipeline Pattern

A pipeline is a sequence of stages.

Example:

```text
Mongo
  |
Read
  |
Transform
  |
Write Parquet
```

Each stage passes data to the next.

---

## Example

```go
Source
  |
  v
Channel
  |
  v
Transformer
  |
  v
Sink
```

---

# 13. Your Pipeline

```go
func (p *Pipeline[T, R]) Run(ctx context.Context) {

	ch := p.source.Read(ctx)

	const workers = 4

	sem := make(chan struct{}, workers)

	var wg sync.WaitGroup

	for data := range ch {

		sem <- struct{}{}

		wg.Go(func() {
			defer func() {
				<-sem
			}()

			transformed :=
				p.transformer.Transform(data)

			p.sink.Write(transformed)
		})
	}

	wg.Wait()
}
```

---

## What Happens?

### Step 1

Source sends batches.

```text
Batch1
Batch2
Batch3
...
```

---

### Step 2

Pipeline receives Batch1.

```text
Acquire semaphore slot
Start worker
```

---

### Step 3

After four batches:

```text
Semaphore full
```

---

### Step 4

Batch5 waits.

---

### Step 5

Worker finishes.

```text
Release semaphore slot
```

---

### Step 6

Batch5 starts.

---

### Step 7

Source channel closes.

Range loop exits.

---

### Step 8

WaitGroup waits for remaining workers.

---

### Step 9

Run returns.

---

# 14. Common Mistakes

## Forgetting Done

```go
wg.Add(1)

go func() {
	// forgot Done
}()
```

Deadlock.

---

## Add After Starting Goroutine

Bad:

```go
go worker()

wg.Add(1)
```

Race condition.

---

Correct:

```go
wg.Add(1)

go worker()
```

---

## Closing From Receiver

Bad:

```go
close(ch)
```

inside consumer.

Sender should close.

---

## Reading Closed Channel Without ok

```go
v := <-ch
```

Returns zero value forever.

Use:

```go
v, ok := <-ch
```

---

# 15. Mental Models

## Goroutine

```text
Lightweight worker
```

---

## Channel

```text
Pipe carrying data
```

---

## WaitGroup

```text
How many workers still running?
```

---

## Semaphore

```text
How many workers may run simultaneously?
```

---

## Select

```text
Wait for multiple channel events
```

---

# Interview Summary

## WaitGroup

```text
Add()  -> Start work
Done() -> Finish work
Wait() -> Wait for all work
```

## Channel

```text
Sender sends
Receiver receives
Sender closes
```

## Select

```text
Wait on multiple channels
```

## Semaphore

```text
Limit concurrency
```

## Worker Pool

```text
Semaphore limits workers

WaitGroup waits for workers
```

The combination of:

```go
WaitGroup + Channel + Semaphore + Select
```

forms the foundation of most production-grade Go concurrency patterns.
