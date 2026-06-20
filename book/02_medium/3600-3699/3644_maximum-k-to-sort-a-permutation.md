# 3644 — Maximum K To Sort A Permutation

## Deskripsi

**Soal:** [3644. Maximum K To Sort A Permutation](https://leetcode.com/problems/maximum-k-to-sort-a-permutation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maximumKToSortAPermutation(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #3644: Maximum K to Sort a Permutation
// https://leetcode.com/problems/maximum-k-to-sort-a-permutation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumKToSortAPermutation(nums []int) int {
	ans := -1 // all bits set to 1 (identity for AND)
	for i, x := range nums {
		if i != x {
			ans &= x
		}
	}
	if ans < 0 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(maximumKToSortAPermutation([]int{0, 3, 2, 1}))
	fmt.Println(maximumKToSortAPermutation([]int{3, 2, 1, 0}))
	fmt.Println(maximumKToSortAPermutation([]int{0, 1, 2, 3}))
}
```
