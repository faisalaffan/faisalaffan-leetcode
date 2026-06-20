# 2089 — Find Target Indices After Sorting Array

## Deskripsi

**Soal:** [2089. Find Target Indices After Sorting Array](https://leetcode.com/problems/find-target-indices-after-sorting-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(1) ignoring sort  
**Kompleksitas Ruang:** O(1) ignoring sort

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2089: Find Target Indices After Sorting Array
// https://leetcode.com/problems/find-target-indices-after-sorting-array/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 2)) // [1 2]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 3)) // [3]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 5)) // [4]
}

// Time: O(n log n), Space: O(1) ignoring sort
func FindTargetIndicesAfterSortingArray(nums []int, target int) []int {
	sort.Ints(nums)
	var result []int
	for i, v := range nums {
		if v == target {
			result = append(result, i)
		}
	}
	return result
}
```
