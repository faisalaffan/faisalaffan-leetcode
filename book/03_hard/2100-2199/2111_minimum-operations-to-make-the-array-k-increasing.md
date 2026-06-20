# 2111 — Minimum Operations To Make The Array K Increasing

## Deskripsi

**Soal:** [2111. Minimum Operations To Make The Array K Increasing](https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

> **Ide Kunci:** LIS per mod-k subsequence.

## Solusi Go

```go
package main

// LeetCode #2111: Minimum Operations to Make the Array K-Increasing
// https://leetcode.com/problems/minimum-operations-to-make-the-array-k-increasing/
// Difficulty: Hard
//
// Approach: LIS per mod-k subsequence.
// Split array into k subsequences: arr[i], arr[i+k], arr[i+2k], ...
// For each subsequence, compute length of longest non-decreasing subsequence (LIS).
// Min operations = len(seq) - LIS(seq). Sum over all k subsequences.

import (
	"fmt"
	"sort"
)

func main() {
	// Example from problem statement
	arr1 := []int{5, 4, 3, 2, 1}
	k1 := 1
	fmt.Printf("minOperations(%v, %d) = %d (expected 4)\n", arr1, k1, minOperations(arr1, k1))

	// Additional tests
	arr2 := []int{4, 1, 5, 2, 6, 2}
	k2 := 2
	fmt.Printf("minOperations(%v, %d) = %d (expected 0)\n", arr2, k2, minOperations(arr2, k2))

	arr3 := []int{1, 2, 3, 4, 5, 6}
	k3 := 2
	fmt.Printf("minOperations(%v, %d) = %d\n", arr3, k3, minOperations(arr3, k3))

	arr4 := []int{5, 3, 1, 4, 2}
	k4 := 2
	fmt.Printf("minOperations(%v, %d) = %d\n", arr4, k4, minOperations(arr4, k4))
}

func minOperations(arr []int, k int) int {
	n := len(arr)
	total := 0

	for i := 0; i < k; i++ {
		seq := []int{}
		for j := i; j < n; j += k {
			seq = append(seq, arr[j])
		}
		total += len(seq) - lengthOfLIS(seq)
	}

	return total
}

// lengthOfLIS returns the length of the longest non-decreasing subsequence.
func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, x := range nums {
		// Find first element > x (since we want non-decreasing, equal values can extend)
		idx := sort.SearchInts(tails, x+1)
		if idx == len(tails) {
			tails = append(tails, x)
		} else {
			tails[idx] = x
		}
	}
	return len(tails)
}
```
