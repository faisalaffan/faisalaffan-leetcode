# 2650 — Design Cancellable Function

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diminta untuk mendesain (merancang) sebuah struktur data kustom dengan operasi tertentu (insert, delete, search, update). Tugasmu adalah memilih representasi data yang tepat agar setiap operasi berjalan efisien — biasanya O(1) atau O(log n).

Ini adalah soal yang paling sering muncul di interview sistem desain. Kamu perlu memilih kombinasi struktur data yang tepat (HashMap + Heap + LinkedList) untuk mencapai kompleksitas yang diminta.

**Konsep kunci:** HashMap (O(1) lookup), Heap (priority), Doubly Linked List (O(1) remove), TreeMap (ordered keys).

**Fungsi yang perlu kamu implementasikan:**
```go
func newCancellable(ctx context.Context, gen func(func(int) error)) <-chan int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

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
