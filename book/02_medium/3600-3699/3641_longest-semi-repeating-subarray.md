# 3641 — Longest Semi Repeating Subarray

## Deskripsi

**Soal:** [3641. Longest Semi Repeating Subarray](https://leetcode.com/problems/longest-semi-repeating-subarray/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func longestSemiRepeatingSubarray(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #3641: Longest Semi-Repeating Subarray
// https://leetcode.com/problems/longest-semi-repeating-subarray/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func longestSemiRepeatingSubarray(nums []int, k int) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	dupCount := 0
	maxLen := 0
	l := 0

	for r, x := range nums {
		freq[x]++
		if freq[x] == 2 {
			dupCount++
		}

		for dupCount > k {
			left := nums[l]
			freq[left]--
			if freq[left] == 1 {
				dupCount--
			}
			l++
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 2, 3, 1, 2, 3, 4}, 2))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 4))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 0))
}
```
