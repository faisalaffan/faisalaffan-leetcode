# 2721 — Execute Asynchronous Functions In Parallel

## Deskripsi

**Soal:** [2721. Execute Asynchronous Functions In Parallel](https://leetcode.com/problems/execute-asynchronous-functions-in-parallel/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func ExecuteAsynchronousFunctionsInParallel(functions []func() int) []int`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
