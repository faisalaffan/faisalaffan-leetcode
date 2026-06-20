# 3684 — Maximize Sum Of At Most K Distinct Elements

## Deskripsi

**Soal:** [3684. Maximize Sum Of At Most K Distinct Elements](https://leetcode.com/problems/maximize-sum-of-at-most-k-distinct-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3684: Maximize Sum of At Most K Distinct Elements
// https://leetcode.com/problems/maximize-sum-of-at-most-k-distinct-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 90}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 93}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{1, 1, 1, 2, 2, 2}, 6))
}

// Time: O(n log n)
// Space: O(n)
func MaximizeSumOfAtMostKDistinctElements(nums []int, k int) []int {
  // Membuat map untuk pencarian O(1): key → value
	seen := make(map[int]bool)
  // Membuat slice untuk menyimpan hasil
	unique := make([]int, 0)
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	size := k
	if size > len(unique) {
		size = len(unique)
	}
	return unique[:size]
}
```
