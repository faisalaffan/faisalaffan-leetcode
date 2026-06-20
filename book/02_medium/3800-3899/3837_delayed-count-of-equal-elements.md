# 3837 — Delayed Count Of Equal Elements

## Deskripsi

**Soal:** [3837. Delayed Count Of Equal Elements](https://leetcode.com/problems/delayed-count-of-equal-elements/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

**Fungsi Solusi:** `func DelayedCountOfEqualElements(nums []int) int64`

> **Ide Kunci:** Track frequency of each value and count pairs where values are equal.

## Solusi Go

```go
package main

// LeetCode #3837: Delayed Count of Equal Elements
// https://leetcode.com/problems/delayed-count-of-equal-elements/
// Difficulty: Medium [Paid]
// Time: O(N) | Space: O(N)
// Approach: Track frequency of each value and count pairs where values are equal.

import "fmt"

func DelayedCountOfEqualElements(nums []int) int64 {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int64)
	var ans int64

	for _, v := range nums {
		ans += freq[v]
		freq[v]++
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 1, 2, 1})) // Expected: 4

	// Example 2
	fmt.Println(DelayedCountOfEqualElements([]int{1, 1, 1, 1})) // Expected: 6

	// Example 3
	fmt.Println(DelayedCountOfEqualElements([]int{1, 2, 3})) // Expected: 0
}
```
