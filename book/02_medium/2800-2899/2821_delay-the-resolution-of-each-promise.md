# 2821 — Delay The Resolution Of Each Promise

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func DelayTheResolutionOfEachPromise(functions []func() int, delay time.Duration) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

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
  // Alokasi slice
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
