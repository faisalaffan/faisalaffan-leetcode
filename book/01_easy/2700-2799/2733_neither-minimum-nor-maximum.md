# 2733 — Neither Minimum Nor Maximum

## Deskripsi

**Soal:** [2733. Neither Minimum Nor Maximum](https://leetcode.com/problems/neither-minimum-nor-maximum/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2733: Neither Minimum nor Maximum
// https://leetcode.com/problems/neither-minimum-nor-maximum/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NeitherMinimumNorMaximum([]int{3, 2, 1, 4}))
	fmt.Println(NeitherMinimumNorMaximum([]int{1, 2}))
}

func NeitherMinimumNorMaximum(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	for _, v := range nums {
		if v != minVal && v != maxVal {
			return v
		}
	}

	return -1
}
```
