# 2721 — Execute Asynchronous Functions In Parallel

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ExecuteAsynchronousFunctionsInParallel(functions []func() int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2721: Execute Asynchronous Functions in Parallel
// https://leetcode.com/problems/execute-asynchronous-functions-in-parallel/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"sync"
)

func ExecuteAsynchronousFunctionsInParallel(functions []func() int) []int {
	var wg sync.WaitGroup
  // Alokasi slice
	results := make([]int, len(functions))

	for i, fn := range functions {
		wg.Add(1)
		go func(idx int, f func() int) {
			defer wg.Done()
			results[idx] = f()
		}(i, fn)
	}

	wg.Wait()
	return results
}

func main() {
	results := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
	})
	fmt.Println(results)

	results2 := ExecuteAsynchronousFunctionsInParallel([]func() int{
		func() int { return 42 },
	})
	fmt.Println(results2)
}
```
