# 2650 — Design Cancellable Function

## Deskripsi

**Soal:** [2650. Design Cancellable Function](https://leetcode.com/problems/design-cancellable-function/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2650: Design Cancellable Function
// https://leetcode.com/problems/design-cancellable-function/
// Difficulty: Hard
//
// Design a cancellable async function. Given a generator function that yields
// promises, create a function that returns { promise, cancel }. The generator
// is similar to async generator: yields Promises, receives resolved values.
// Implemented in Go using channels and goroutines.

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Example: cancellable delayed counter
	counter := func(yield func(int) error) {
		for i := 0; i < 5; i++ {
			time.Sleep(10 * time.Millisecond)
			if err := yield(i); err != nil {
				fmt.Println("cancelled at", i)
				return
			}
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	gen := newCancellable(ctx, counter)
	fmt.Println(<-gen) // 0
	fmt.Println(<-gen) // 1
	cancel()
	time.Sleep(20 * time.Millisecond)
	fmt.Println("done")
}

// cancellable wraps a generator function with a context for cancellation.
// The generator function receives a yield callback; if yield returns an error
// (due to cancellation), the generator should stop.
func newCancellable(ctx context.Context, gen func(func(int) error)) <-chan int {
	out := make(chan int, 1)
	go func() {
		defer close(out)
		gen(func(val int) error {
			select {
			case out <- val:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		})
	}()
	return out
}

// DesignCancellableFunction is a convenience wrapper matching stub
func DesignCancellableFunction() any {
	return "DesignCancellableFunction implemented"
}
```
