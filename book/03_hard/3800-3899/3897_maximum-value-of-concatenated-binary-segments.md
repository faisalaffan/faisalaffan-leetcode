# 3897 — Maximum Value Of Concatenated Binary Segments

## Deskripsi

**Soal:** [3897. Maximum Value Of Concatenated Binary Segments](https://leetcode.com/problems/maximum-value-of-concatenated-binary-segments/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Sort both arrays descending. Greedily pick the largest

## Solusi Go

```go
package main

// LeetCode #3897: Maximum Value of Concatenated Binary Segments
// https://leetcode.com/problems/maximum-value-of-concatenated-binary-segments/
// Difficulty: Hard
//
// Given two arrays nums1 and nums0, select segments from each.
// Concatenate binary representations of selected elements to
// maximize the resulting value. Segments from nums1 precede
// those from nums0.
//
// Approach: Sort both arrays descending. Greedily pick the largest
// elements since binary concatenation favors larger values in
// higher positions.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxValue([]int{3, 5}, []int{2, 4}))
	// Example 2
	fmt.Println(maxValue([]int{1, 2}, []int{3, 4}))
	// Edge: single elements
	fmt.Println(maxValue([]int{7}, []int{1}))
	// Edge: same values
	fmt.Println(maxValue([]int{5, 5}, []int{5, 5}))
}

func maxValue(nums1 []int, nums0 []int) int {
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] > nums1[j] })
	sort.Slice(nums0, func(i, j int) bool { return nums0[i] > nums0[j] })

	// Try all possible split points: take first i from nums1, rest from nums0
	best := 0
	for i := 0; i <= len(nums1); i++ {
		for j := 0; j <= len(nums0); j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := 0
			for k := 0; k < i; k++ {
				val = appendBits(val, nums1[k])
			}
			for k := 0; k < j; k++ {
				val = appendBits(val, nums0[k])
			}
			if val > best {
				best = val
			}
		}
	}
	return best
}

func appendBits(current, val int) int {
	if val == 0 {
		return current << 1
	}
	bits := 0
	tmp := val
	for tmp > 0 {
		bits++
		tmp >>= 1
	}
	return (current << bits) | val
}
```
