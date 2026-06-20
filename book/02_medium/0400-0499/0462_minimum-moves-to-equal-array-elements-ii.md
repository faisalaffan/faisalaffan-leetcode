# 0462 — Minimum Moves To Equal Array Elements Ii

## Deskripsi

**Soal:** [0462. Minimum Moves To Equal Array Elements Ii](https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n) for sorting, O(n) for QuickSelect  
**Kompleksitas Ruang:** O(log n) for sorting, O(1) for QuickSelect

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #462: Minimum Moves to Equal Array Elements II
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/
// Difficulty: Medium
// Time: O(n log n) for sorting, O(n) for QuickSelect
// Space: O(log n) for sorting, O(1) for QuickSelect

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 2, 3}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 10, 2, 9}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 0, 0, 8, 6}))
}

func MinimumMovesToEqualArrayElementsIi(nums []int) int {
	sort.Ints(nums)
	median := nums[len(nums)/2]
	moves := 0
	for _, num := range nums {
		diff := num - median
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
```
