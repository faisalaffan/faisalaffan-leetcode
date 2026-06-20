# 3507 — Minimum Pair Removal To Sort Array I

## Deskripsi

**Soal:** [3507. Minimum Pair Removal To Sort Array I](https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2). Space: O(n).  
**Kompleksitas Ruang:** O(n).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3507: Minimum Pair Removal to Sort Array I
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumPairRemovalToSortArrayI([]int{5, 2, 3, 1}))
	fmt.Println(MinimumPairRemovalToSortArrayI([]int{1, 2, 3, 4}))
}

// MinimumPairRemovalToSortArrayI returns the minimum number of pairs to remove so the remaining array is sorted.
// A pair is two adjacent elements.
// Time: O(n^2). Space: O(n).
func MinimumPairRemovalToSortArrayI(nums []int) int {
	n := len(nums)
	// If already sorted
	sorted := true
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			sorted = false
			break
		}
	}
	if sorted {
		return 0
	}

	// Try removing pairs
	minOps := n
	var try func(arr []int, ops int)
	try = func(arr []int, ops int) {
		if ops >= minOps {
			return
		}
		// Check sorted
		ok := true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				ok = false
				break
			}
		}
		if ok {
			if ops < minOps {
				minOps = ops
			}
			return
		}
		if len(arr) < 2 {
			return
		}
		// Remove each possible adjacent pair
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(arr)-1; i++ {
  // Membuat slice untuk menyimpan hasil
			next := make([]int, 0, len(arr)-2)
			next = append(next, arr[:i]...)
			next = append(next, arr[i+2:]...)
			try(next, ops+1)
		}
	}
	try(nums, 0)
	return minOps
}
```
