# 1208 — Get Equal Substrings Within Budget

## Deskripsi

**Soal:** [1208. Get Equal Substrings Within Budget](https://leetcode.com/problems/get-equal-substrings-within-budget/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Sliding Window (jendela geser)

**Fungsi Solusi:** `func equalSubstring(s string, t string, maxCost int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1208: Get Equal Substrings Within Budget
// https://leetcode.com/problems/get-equal-substrings-within-budget/
// Difficulty: Medium

// Sliding window: maxCost - cost[i] = difference between s[i] and t[i].
// Find longest substring with total cost <= maxCost.

// Time: O(n)
// Space: O(1)

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		c := int(s[i]) - int(t[i])
		if c < 0 {
			c = -c
		}
		cost[i] = c
	}

	left := 0
	currCost := 0
	maxLen := 0

	for right := 0; right < n; right++ {
		currCost += cost[right]
		for currCost > maxCost {
			currCost -= cost[left]
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Printf("%d (expected: 3)\n", equalSubstring("abcd", "bcdf", 3))
	fmt.Printf("%d (expected: 1)\n", equalSubstring("abcd", "cdef", 3))
	fmt.Printf("%d (expected: 1)\n", equalSubstring("abcd", "acde", 0))
}
```
