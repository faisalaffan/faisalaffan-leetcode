# 3769 — Sort Integers By Binary Reflection

## Deskripsi

**Soal:** [3769. Sort Integers By Binary Reflection](https://leetcode.com/problems/sort-integers-by-binary-reflection/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3769: Sort Integers by Binary Reflection
// https://leetcode.com/problems/sort-integers-by-binary-reflection/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortIntegersByBinaryReflection([]int{3, 6, 5}))
	fmt.Println(SortIntegersByBinaryReflection([]int{1, 2, 3}))
}

// Time: O(n log n)
// Space: O(n)
func SortIntegersByBinaryReflection(nums []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	reflections := make(map[int]int)
	for _, v := range nums {
		reflections[v] = reverseBits(v)
	}

	sort.Slice(nums, func(i, j int) bool {
		if reflections[nums[i]] != reflections[nums[j]] {
			return reflections[nums[i]] < reflections[nums[j]]
		}
		return nums[i] < nums[j]
	})
	return nums
}

func reverseBits(n int) int {
	rev := 0
	for n > 0 {
		rev = (rev << 1) | (n & 1)
		n >>= 1
	}
	return rev
}
```
