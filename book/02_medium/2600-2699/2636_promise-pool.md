# 2636 — Promise Pool

## Deskripsi

**Soal:** [2636. Promise Pool](https://leetcode.com/problems/promise-pool/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func promisePool(tasks []func(int) (int, error), limit int) []int`

## Solusi Go

```go
package main

// LeetCode #2636: Promise Pool
// https://leetcode.com/problems/promise-pool/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sync"
	"time"
)

func promisePool(tasks []func(int) (int, error), limit int) []int {
	var wg sync.WaitGroup
  // Membuat slice untuk menyimpan hasil
	results := make([]int, len(tasks))
	sem := make(chan struct{}, limit)

	for i, task := range tasks {
		wg.Add(1)
		sem <- struct{}{}
		go func(idx int, t func(int) (int, error)) {
			defer wg.Done()
			defer func() { <-sem }()
			result, _ := t(idx)
			results[idx] = result
		}(i, task)
	}

	wg.Wait()
	return results
}

func main() {
	tasks := []func(int) (int, error){
		func(i int) (int, error) { time.Sleep(50 * time.Millisecond); return i * 2, nil },
		func(i int) (int, error) { time.Sleep(30 * time.Millisecond); return i * 3, nil },
		func(i int) (int, error) { time.Sleep(10 * time.Millisecond); return i * 4, nil },
	}

	results := promisePool(tasks, 2)
	fmt.Println("Test 1:", results)
	// Expected: [0, 3, 8]

	// Test case 2: empty tasks
	fmt.Println("Test 2:", promisePool([]func(int) (int, error){}, 5))
	// Expected: []

	// Test case 3: single task
	single := []func(int) (int, error){
		func(i int) (int, error) { return 42, nil },
	}
	fmt.Println("Test 3:", promisePool(single, 1))
	// Expected: [42]
}
```
