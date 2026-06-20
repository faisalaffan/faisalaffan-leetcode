# 2389 — Longest Subsequence With Limited Sum

## Deskripsi

**Soal:** [2389. Longest Subsequence With Limited Sum](https://leetcode.com/problems/longest-subsequence-with-limited-sum/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

```go
package main

// LeetCode #2389: Longest Subsequence With Limited Sum
// https://leetcode.com/problems/longest-subsequence-with-limited-sum/
// Difficulty: Easy
// Time O((n+m) log n) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LongestSubsequenceWithLimitedSum([]int{4, 5, 2, 1}, []int{3, 10, 21})) // [2,3,4]
	fmt.Println(LongestSubsequenceWithLimitedSum([]int{2, 3, 4, 5}, []int{1}))          // [0]
}

func LongestSubsequenceWithLimitedSum(nums []int, queries []int) []int {
	sort.Ints(nums)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, len(nums))
	sum := 0
	for i, n := range nums {
		sum += n
		prefix[i] = sum
	}

  // Membuat slice untuk menyimpan hasil
	res := make([]int, len(queries))
	for i, q := range queries {
		// Binary search for rightmost index where prefix <= q
		res[i] = sort.SearchInts(prefix, q+1)
	}
	return res
}
```
