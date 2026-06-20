# 1296 — Divide Array In Sets Of K Consecutive Numbers

## Deskripsi

**Soal:** [1296. Divide Array In Sets Of K Consecutive Numbers](https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func isPossibleDivide(nums []int, k int) bool`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1296: Divide Array in Sets of K Consecutive Numbers
// https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
// Difficulty: Medium

// Sort array, greedily form groups of size k.
// Use frequency map to track available numbers.

// Time: O(n log n)
// Space: O(n)

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	sort.Ints(nums)
	for _, v := range nums {
		if freq[v] == 0 {
			continue
		}
		for i := 0; i < k; i++ {
			if freq[v+i] == 0 {
				return false
			}
			freq[v+i]--
		}
	}

	return true
}

func main() {
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 3, 4, 4, 5, 6}, 4))
	fmt.Printf("%t (expected: false)\n", isPossibleDivide([]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3))
	fmt.Printf("%t (expected: true)\n", isPossibleDivide([]int{1, 2, 3, 4}, 2))
}
```
