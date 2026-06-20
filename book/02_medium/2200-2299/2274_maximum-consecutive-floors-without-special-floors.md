# 2274 — Maximum Consecutive Floors Without Special Floors

## Deskripsi

**Soal:** [2274. Maximum Consecutive Floors Without Special Floors](https://leetcode.com/problems/maximum-consecutive-floors-without-special-floors/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxConsecutive(bottom int, top int, special []int) int`

## Solusi Go

```go
package main

// LeetCode #2274: Maximum Consecutive Floors Without Special Floors
// https://leetcode.com/problems/maximum-consecutive-floors-without-special-floors/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)

	maxGap := special[0] - bottom
	for i := 1; i < len(special); i++ {
		gap := special[i] - special[i-1] - 1
		if gap > maxGap {
			maxGap = gap
		}
	}
	gap := top - special[len(special)-1]
	if gap > maxGap {
		maxGap = gap
	}
	return maxGap
}

func main() {
	// Test case 1
	fmt.Println(maxConsecutive(2, 9, []int{4, 6}))
	// Expected: 3

	// Test case 2
	fmt.Println(maxConsecutive(6, 8, []int{7, 6, 8}))
	// Expected: 0

	// Test case 3
	fmt.Println(maxConsecutive(1, 1000000000, []int{1, 1000000000}))
	// Expected: 999999998
}
```
