# 2623 — Memoize

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func memoize(fn func(...int) int) MemoizedFn`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1) amortized per call  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2623: Memoize
// https://leetcode.com/problems/memoize/
// Difficulty: Medium
// Time: O(1) amortized per call | Space: O(n)

import (
	"fmt"
	"sync"
)

type MemoizedFn func(args ...int) int

func memoize(fn func(...int) int) MemoizedFn {
	var mu sync.Mutex
  // HashMap: O(1) lookup
	cache := make(map[string]int)

	return func(args ...int) int {
		// Create a key from args
		key := ""
		for i, arg := range args {
			if i > 0 {
				key += ","
			}
			key += fmt.Sprintf("%d", arg)
		}

		mu.Lock()
		defer mu.Unlock()

		if val, ok := cache[key]; ok {
			return val
		}
		result := fn(args...)
		cache[key] = result
		return result
	}
}

var callCount int

func add(a, b int) int {
	callCount++
	return a + b
}

func main() {
	callCount = 0
	memoizedAdd := memoize(func(args ...int) int {
		return add(args[0], args[1])
	})

	// Test case 1
	fmt.Println("Test 1:", memoizedAdd(1, 2))
	// Expected: 3

	// Test case 2: same args, should use cache
	fmt.Println("Test 2:", memoizedAdd(1, 2))
	// Expected: 3

	// Test case 3: different args
	fmt.Println("Test 3:", memoizedAdd(2, 3))
	// Expected: 5

	fmt.Println("Calls:", callCount)
	// Expected: 2 (not 3)
}
```
