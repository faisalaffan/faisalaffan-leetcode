# 3431 — Minimum Unlocked Indices To Sort Nums

## Deskripsi

**Soal:** [3431. Minimum Unlocked Indices To Sort Nums](https://leetcode.com/problems/minimum-unlocked-indices-to-sort-nums/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2) Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minUnlockedIndices(nums []int, locked []int) int`

## Solusi Go

```go
package main

// LeetCode #3431: Minimum Unlocked Indices to Sort Nums
// https://leetcode.com/problems/minimum-unlocked-indices-to-sort-nums/
// Difficulty: Medium [Paid]
// Time: O(n^2) Space: O(n)

import (
	"fmt"
	"sort"
)

func minUnlockedIndices(nums []int, locked []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)

	unlockCost := 0
	for i := 0; i < n; i++ {
		if nums[i] != sorted[i] && locked[i] == 1 {
			unlockCost++
		}
	}
	return unlockCost
}

func main() {
	fmt.Println(minUnlockedIndices([]int{1, 3, 2, 4}, []int{1, 1, 0, 1})) // 1
	fmt.Println(minUnlockedIndices([]int{1, 2, 3}, []int{1, 1, 1}))       // 0
}
```
