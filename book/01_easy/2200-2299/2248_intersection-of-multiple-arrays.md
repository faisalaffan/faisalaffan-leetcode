# 2248 — Intersection Of Multiple Arrays

## Deskripsi

**Soal:** [2248. Intersection Of Multiple Arrays](https://leetcode.com/problems/intersection-of-multiple-arrays/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n * m), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2248: Intersection of Multiple Arrays
// https://leetcode.com/problems/intersection-of-multiple-arrays/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(IntersectionOfMultipleArrays([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}})) // [3 4]
	fmt.Println(IntersectionOfMultipleArrays([][]int{{1, 2, 3}, {4, 5, 6}}))                        // []
}

// Time: O(n * m), Space: O(n)
func IntersectionOfMultipleArrays(nums [][]int) []int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return []int{}
	}

  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, v := range nums[0] {
		freq[v] = 1
	}

	for i := 1; i < len(nums); i++ {
  // Membuat map untuk pencarian O(1): key → value
		seen := make(map[int]bool)
		for _, v := range nums[i] {
			if !seen[v] {
				freq[v]++
				seen[v] = true
			}
		}
	}

	var result []int
	for v, c := range freq {
		if c == len(nums) {
			result = append(result, v)
		}
	}
	sort.Ints(result)
	return result
}
```
