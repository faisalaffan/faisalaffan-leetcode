# 1887 — Reduction Operations To Make The Array Elements Equal

## Deskripsi

**Soal:** [1887. Reduction Operations To Make The Array Elements Equal](https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1887: Reduction Operations to Make the Array Elements Equal
// https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(ReductionOperations([]int{5, 1, 3}))
	fmt.Println(ReductionOperations([]int{1, 1, 1}))
	fmt.Println(ReductionOperations([]int{1, 1, 2, 2, 3}))
}

// Time: O(n log n), Space: O(1)
func ReductionOperations(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	ops := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			ops += n - i
		}
	}
	return ops
}
```
