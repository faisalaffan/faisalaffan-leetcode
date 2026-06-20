# 2099 — Find Subsequence Of Length K With The Largest Sum

## Deskripsi

**Soal:** [2099. Find Subsequence Of Length K With The Largest Sum](https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2099: Find Subsequence of Length K With the Largest Sum
// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{2, 1, 3, 3}, 2))       // [3 3]
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{-1, -2, 3, 4}, 3))    // [-1 3 4]
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{3, 4, 3, 3}, 2))      // [3 4]
}

// Time: O(n log n), Space: O(n)
func FindSubsequenceOfLengthKWithTheLargestSum(nums []int, k int) []int {
	type pair struct {
		val int
		idx int
	}

  // Membuat slice untuk menyimpan hasil
	pairs := make([]pair, len(nums))
	for i, v := range nums {
		pairs[i] = pair{v, i}
	}

	// Sort by value descending
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val > pairs[j].val
	})

	// Take top k
	selected := pairs[:k]

	// Sort by original index to preserve order
	sort.Slice(selected, func(i, j int) bool {
		return selected[i].idx < selected[j].idx
	})

  // Membuat slice untuk menyimpan hasil
	result := make([]int, k)
	for i, p := range selected {
		result[i] = p.val
	}
	return result
}
```
