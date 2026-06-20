# 3022 — Minimize Or Of Remaining Elements Using Operations

## Deskripsi

**Soal:** [3022. Minimize Or Of Remaining Elements Using Operations](https://leetcode.com/problems/minimize-or-of-remaining-elements-using-operations/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal)

**Fungsi Solusi:** `func minOrAfterOperations(nums []int, k int) int`

> **Ide Kunci:** Greedy bit-by-bit (high to low)

## Solusi Go

```go
package main

// LeetCode #3022: Minimize OR of Remaining Elements Using Operations
// https://leetcode.com/problems/minimize-or-of-remaining-elements-using-operations/
// Difficulty: Hard
//
// Given an array nums and integer k. We can perform at most k operations.
// Each operation: pick two adjacent elements, replace them with their AND.
// After all operations, n-k elements remain. Minimize the OR of all remaining elements.
//
// Approach: Greedy bit-by-bit (high to low)
//   For each bit from 29 down to 0, try to make it 0 in the final result.
//   A bit can be cleared if we can partition the array into segments where
//   each segment's AND (restricted to the "candidate zero bits") is 0.
//   Each segment requires one operation to collapse, so we need at most k
//   operations => at most k+1 segments where AND has no forbidden bits.
//   We count segments where cumulative AND != 0 (still has forbidden bits);
//   if this count <= k, the bit can be cleared.

import "fmt"

func minOrAfterOperations(nums []int, k int) int {
	ans := 0       // bits that MUST be 1
	tryMask := 0   // bits we are trying to make 0
	for b := 29; b >= 0; b-- {
		tryMask |= 1 << b
		cnt := 0
		and := -1 // all bits set
		for _, x := range nums {
			and &= x & tryMask
			if and != 0 {
				cnt++
			} else {
				and = -1
			}
		}
		if cnt > k {
			ans |= 1 << b
			tryMask ^= 1 << b
		}
	}
	return ans
}

func main() {
	// Example: [3,5,3,2,7], k=2 -> 3
	fmt.Println("Test 1:", minOrAfterOperations([]int{3, 5, 3, 2, 7}, 2))

	// Example: [7,3,15,14,2,8], k=4 -> 2
	fmt.Println("Test 2:", minOrAfterOperations([]int{7, 3, 15, 14, 2, 8}, 4))

	// Single element
	fmt.Println("Test 3:", minOrAfterOperations([]int{5}, 0))

	// k = 0 (no operations allowed)
	fmt.Println("Test 4:", minOrAfterOperations([]int{1, 2, 4}, 0))

	// All same
	fmt.Println("Test 5:", minOrAfterOperations([]int{7, 7, 7, 7}, 2))

	// Large k
	fmt.Println("Test 6:", minOrAfterOperations([]int{8, 4, 2, 1}, 3))

	// Edge: empty or single
	fmt.Println("Test 7:", minOrAfterOperations([]int{0}, 0))
}
```
