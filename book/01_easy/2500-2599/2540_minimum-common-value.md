# 2540 — Minimum Common Value

## Deskripsi

**Soal:** [2540. Minimum Common Value](https://leetcode.com/problems/minimum-common-value/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2540: Minimum Common Value
// https://leetcode.com/problems/minimum-common-value/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumCommonValue([]int{1, 2, 3}, []int{2, 4}))       // 2
	fmt.Println(MinimumCommonValue([]int{1, 2, 3, 6}, []int{2, 3, 4, 5})) // 2
}

func MinimumCommonValue(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
```
