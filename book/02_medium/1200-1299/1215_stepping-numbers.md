# 1215 — Stepping Numbers

## Deskripsi

**Soal:** [1215. Stepping Numbers](https://leetcode.com/problems/stepping-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(2^n) where n = number of digits in high  
**Kompleksitas Ruang:** O(2^n)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func countSteppingNumbers(low int, high int) []int`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1215: Stepping Numbers
// https://leetcode.com/problems/stepping-numbers/
// Difficulty: Medium [Paid]

// Stepping numbers are numbers where adjacent digits differ by 1.
// Return all stepping numbers in range [low, high] sorted.

// Time: O(2^n) where n = number of digits in high
// Space: O(2^n)

func countSteppingNumbers(low int, high int) []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, 0)

	var dfs func(num int)
	dfs = func(num int) {
		if num > high {
			return
		}
		if num >= low {
			result = append(result, num)
		}
		lastDigit := num % 10
		if lastDigit > 0 {
			dfs(num*10 + (lastDigit - 1))
		}
		if lastDigit < 9 {
			dfs(num*10 + (lastDigit + 1))
		}
	}

	for i := 0; i <= 9; i++ {
		dfs(i)
	}

	sort.Ints(result)
	return result
}

func main() {
	fmt.Printf("%v (expected: [0 1 2 3 4 5 6 7 8 9 10 12])\n", countSteppingNumbers(0, 12))
	fmt.Printf("%v (expected: [10 12])\n", countSteppingNumbers(10, 14))
}
```
