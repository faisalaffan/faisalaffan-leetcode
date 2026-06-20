# 3942 — Minimum Operations To Sort A Permutation

## Deskripsi

**Soal:** [3942. Minimum Operations To Sort A Permutation](https://leetcode.com/problems/minimum-operations-to-sort-a-permutation/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumOperationsToSortAPermutation(nums []int) int`

> **Ide Kunci:** Available ops are reverse and rotate left. Sorted array must be

## Solusi Go

```go
package main

// LeetCode #3942: Minimum Operations to Sort a Permutation
// https://leetcode.com/problems/minimum-operations-to-sort-a-permutation/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Available ops are reverse and rotate left. Sorted array must be
// [0,1,...,n-1] or a rotation of it. Check all rotations of sorted array vs
// all rotations of array and its reverse. Return min ops.

import "fmt"

func MinimumOperationsToSortAPermutation(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Expected sorted array
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	for i := 0; i < n; i++ {
		sorted[i] = i
	}

	// Find position of 0 in nums. In sorted, 0 is at index 0.
	// After rotation, 0 moves. Sorted rotated by k: 0 at index (n-k)%n.
	// Actually: rotate left by k means element at index k moves to 0.
	// sorted left-rotated by k: [k, k+1, ..., n-1, 0, 1, ..., k-1]
	// In this rotation, 0 is at index n-k.

	// Check all rotations of sorted vs all rotations of nums
	// Ops to get from nums to a target rotation of sorted:
	// 1. Maybe reverse first, then rotate
	// 2. Maybe rotate first, then reverse

	// Strategy: try all possible rotations.

	// First, find pos of 0 in nums
	pos0 := -1
	for i, v := range nums {
		if v == 0 {
			pos0 = i
			break
		}
	}

	ans := -1

	// Try: rotate nums to match sorted (no reverse)
	// Need to rotate so that sorted[0]=0 aligns with nums[pos0]=0
	// Rotate left by pos0: nums becomes [nums[pos0], nums[pos0+1], ..., nums[(pos0-1+n)%n]]
	// This puts '0' at front, which matches sorted[0]=0
	// Then check if everything matches
	matches := true
	for i := 0; i < n; i++ {
		if nums[(pos0+i)%n] != i {
			matches = false
			break
		}
	}
	if matches {
		ans = pos0 // just rotate, no reverse
	}

	// Try: reverse then rotate to match sorted
	// Reverse: [nums[n-1], nums[n-2], ..., nums[0]]
  // Membuat slice untuk menyimpan hasil
	reversed := make([]int, n)
	for i := 0; i < n; i++ {
		reversed[i] = nums[n-1-i]
	}

	// Find pos of 0 in reversed
	pos0Rev := -1
	for i, v := range reversed {
		if v == 0 {
			pos0Rev = i
			break
		}
	}

	matches = true
	for i := 0; i < n; i++ {
		if reversed[(pos0Rev+i)%n] != i {
			matches = false
			break
		}
	}
	if matches {
		ops := 1 + pos0Rev // 1 reverse + pos0Rev rotations
		if ans == -1 || ops < ans {
			ans = ops
		}
	}

	// Try: rotate then reverse
	// But which rotation? We need sorted after all ops.
	// sorted rotated by some amount k.
	// We can also rotate nums (by r), then reverse.
	// sorted left-rotated by k: element k, k+1, ..., n-1, 0, ..., k-1
	// After reversal of (nums rotated by r): we get reversed_rotated
	// reversed_rotated[i] = nums[(r-1-i+n)%n] (since rotate left by r, then reverse)

	for k := 0; k < n; k++ {
		// Check if (nums rotated by r) reversed equals sorted rotated by k
		// for some r

		for r := 0; r < n; r++ {
			matches = true
			for i := 0; i < n; i++ {
				// rot = nums[(r+i)%n]
				// rev_rot = rot[n-1-i] = nums[(r+(n-1-i))%n]
				// should equal sorted[(k+i)%n]
				expected := (k + i) % n
				actual := nums[(r + n - 1 - i) % n]
				if actual != expected {
					matches = false
					break
				}
			}
			if matches {
				ops := r + 1 + k // r rotations, 1 reverse, k rotations
				if ans == -1 || ops < ans {
					ans = ops
				}
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToSortAPermutation([]int{0, 2, 1})) // Expected: 2

	// Example 2
	fmt.Println(MinimumOperationsToSortAPermutation([]int{1, 0, 2})) // Expected: 2

	// Example 3
	fmt.Println(MinimumOperationsToSortAPermutation([]int{2, 0, 1, 3})) // Expected: -1
}
```
