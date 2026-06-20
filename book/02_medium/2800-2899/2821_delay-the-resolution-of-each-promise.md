# 2821 — Delay The Resolution Of Each Promise

## Deskripsi

**Soal:** [2821. Delay The Resolution Of Each Promise](https://leetcode.com/problems/delay-the-resolution-of-each-promise/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func DelayTheResolutionOfEachPromise(functions []func() int, delay time.Duration) []int`

## Solusi Go

```go
package main

// LeetCode #2821: Delay the Resolution of Each Promise
// https://leetcode.com/problems/delay-the-resolution-of-each-promise/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

func DelayTheResolutionOfEachPromise(functions []func() int, delay time.Duration) []int {
	n := len(functions)
  // Membuat slice untuk menyimpan hasil
	results := make([]int, n)
	var wg sync.WaitGroup

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			time.Sleep(delay)
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	start := time.Now()
	results := DelayTheResolutionOfEachPromise([]func() int{
		func() int { return 10 },
		func() int { return 20 },
		func() int { return 30 },
	}, 5*time.Millisecond)
	fmt.Println(results)
	fmt.Println("Took:", time.Since(start).Round(time.Millisecond))
}
```
