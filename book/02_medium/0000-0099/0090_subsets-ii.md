# 0090 — Subsets Ii

## Deskripsi

**Soal:** [0090. Subsets Ii](https://leetcode.com/problems/subsets-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * 2^n)  
**Kompleksitas Ruang:** O(n * 2^n)

**Algoritma:** —

**Fungsi Solusi:** `func subsetsWithDup(nums []int) [][]int`

## Solusi Go

```go
package main

// LeetCode #90: Subsets II
// https://leetcode.com/problems/subsets-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	start := 0

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		n := len(result)
		begin := 0
		if i > 0 && nums[i] == nums[i-1] {
			begin = start
		}
		start = n
		for j := begin; j < n; j++ {
  // Membuat slice untuk menyimpan hasil
			newSubset := make([]int, len(result[j])+1)
			copy(newSubset, result[j])
			newSubset[len(result[j])] = nums[i]
			result = append(result, newSubset)
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	// [[] [1] [2] [1 2] [2 2] [1 2 2]]

	// Test case 2
	fmt.Println(subsetsWithDup([]int{0})) // [[] [0]]
}

// Time: O(n * 2^n) | Space: O(n * 2^n)
```
