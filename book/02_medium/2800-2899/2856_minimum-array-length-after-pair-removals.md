# 2856 — Minimum Array Length After Pair Removals

## Deskripsi

**Soal:** [2856. Minimum Array Length After Pair Removals](https://leetcode.com/problems/minimum-array-length-after-pair-removals/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumArrayLengthAfterPairRemovals(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2856: Minimum Array Length After Pair Removals
// https://leetcode.com/problems/minimum-array-length-after-pair-removals/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func MinimumArrayLengthAfterPairRemovals(nums []int) int {
	n := len(nums)
	// Find max frequency
	maxFreq := 0
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
		if freq[v] > maxFreq {
			maxFreq = freq[v]
		}
	}

	remaining := n - maxFreq
	if maxFreq > remaining {
		return maxFreq - remaining
	}
	if n%2 == 0 {
		return 0
	}
	return 1
}

func main() {
	fmt.Println(MinimumArrayLengthAfterPairRemovals([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(MinimumArrayLengthAfterPairRemovals([]int{1, 1, 2, 3}))
}
```
