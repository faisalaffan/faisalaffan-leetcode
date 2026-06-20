# 3510 — Minimum Pair Removal To Sort Array Ii

## Deskripsi

**Soal:** [3510. Minimum Pair Removal To Sort Array Ii](https://leetcode.com/problems/minimum-pair-removal-to-sort-array-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Merge pairs from left to right. For each element,

## Solusi Go

```go
package main

// LeetCode #3510: Minimum Pair Removal to Sort Array II
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-ii/
// Difficulty: Hard
//
// In one operation, remove a pair of adjacent elements and replace
// with their sum. Minimum operations to make the resulting array
// non-decreasing.
//
// Approach: Merge pairs from left to right. For each element,
// if it's smaller than the previous, merge it with previous.
// Continue merging until sorted.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minimumPairRemoval([]int{1, 3, 2, 4}))
	// Example 2
	fmt.Println(minimumPairRemoval([]int{5, 4, 3, 2, 1}))
	// Edge: already sorted
	fmt.Println(minimumPairRemoval([]int{1, 2, 3}))
}

func minimumPairRemoval(nums []int) int {
	// Simulate the process: repeatedly find and merge the first
	// adjacent pair where left > right (inversion), then merge them
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	arr := make([]int64, n)
	for i, v := range nums {
		arr[i] = int64(v)
	}

	ops := 0
	for {
		sorted := true
		mergeIdx := -1
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				sorted = false
				if mergeIdx == -1 {
					mergeIdx = i
				}
				break
			}
		}
		if sorted {
			break
		}

		// Merge mergeIdx and mergeIdx+1
		arr[mergeIdx] = arr[mergeIdx] + arr[mergeIdx+1]
		arr = append(arr[:mergeIdx+1], arr[mergeIdx+2:]...)
		ops++
	}

	return ops
}
```
