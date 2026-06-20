# 3674 — Minimum Operations To Equalize Array

## Deskripsi

**Soal:** [3674. Minimum Operations To Equalize Array](https://leetcode.com/problems/minimum-operations-to-equalize-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3674: Minimum Operations to Equalize Array
// https://leetcode.com/problems/minimum-operations-to-equalize-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToEqualizeArray([]int{1, 2, 3}))
	fmt.Println(MinimumOperationsToEqualizeArray([]int{5, 5, 5}))
}

// Time: O(n)
// Space: O(1)
func MinimumOperationsToEqualizeArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[0] {
			return 1
		}
	}
	return 0
}
```
