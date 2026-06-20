# 3287 — Find The Maximum Sequence Value Of Array

## Deskripsi

**Soal:** [3287. Find The Maximum Sequence Value Of Array](https://leetcode.com/problems/find-the-maximum-sequence-value-of-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Bitmask (representasi himpunan dengan bit)

> **Ide Kunci:** //  1. For each split point, consider elements before the split (left part)

## Solusi Go

```go
package main

// LeetCode #3287: Find the Maximum Sequence Value of Array
// https://leetcode.com/problems/find-the-maximum-sequence-value-of-array/
// Difficulty: Hard
//
// Given an array nums and integer k, select a subsequence of exactly 2k+1
// elements. The value of the subsequence is (OR of first k+1 elements)
// XOR (OR of last k elements). Maximize this value.
//
// Approach:
//  1. For each split point, consider elements before the split (left part)
//     and after the split (right part).
//  2. DP to compute all possible OR values achievable with exactly `cnt`
//     elements from a prefix or suffix.
//  3. nums[i] <= 127 so OR values fit in 7 bits (0..127).
//  4. Combine left and right OR values to find max XOR.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(maxSequenceValue([]int{2, 6, 7}, 1))
	// Example 2
	fmt.Println(maxSequenceValue([]int{4, 2, 5, 6, 7}, 2))
	// Example 3: single element with k=0 => 2k+1 = 1
	fmt.Println(maxSequenceValue([]int{5}, 0))
	// Example 4
	fmt.Println(maxSequenceValue([]int{1, 2, 3, 4, 5, 6, 7}, 2))
	// Example 5
	fmt.Println(maxSequenceValue([]int{10, 20, 30, 40, 50}, 1))
}

func maxSequenceValue(nums []int, k int) int {
	n := len(nums)
	leftSize := k + 1
	rightSize := k
	maxOr := 128 // 7 bits (0..127)

	// leftDP[i][j] = bitmask of OR values achievable using j elements
	// from prefix nums[0..i-1].
  // Membuat slice 2D untuk DP/tabel
	leftDP := make([][]uint64, n+1)
  // Iterasi seluruh elemen
	for i := range leftDP {
		leftDP[i] = make([]uint64, leftSize+1)
	}
	leftDP[0][0] = 1 << 0 // OR value 0 is achievable with 0 elements

	for i := 0; i < n; i++ {
		v := nums[i]
		for j := 0; j <= leftSize; j++ {
			// Don't take nums[i].
			leftDP[i+1][j] |= leftDP[i][j]
			// Take nums[i].
			if j+1 <= leftSize {
				mask := leftDP[i][j]
				// For each achievable OR value, compute new OR with v.
				var newMask uint64
				for or := 0; or < maxOr; or++ {
					if mask&(1<<or) != 0 {
						newMask |= 1 << (or | v)
					}
				}
				leftDP[i+1][j+1] |= newMask
			}
		}
	}

	// rightDP[i][j] = bitmask of OR values achievable using j elements
	// from suffix nums[i..n-1].
  // Membuat slice 2D untuk DP/tabel
	rightDP := make([][]uint64, n+1)
  // Iterasi seluruh elemen
	for i := range rightDP {
		rightDP[i] = make([]uint64, rightSize+1)
	}
	rightDP[n][0] = 1 << 0 // OR value 0 with 0 elements

	for i := n - 1; i >= 0; i-- {
		v := nums[i]
		for j := 0; j <= rightSize; j++ {
			// Don't take nums[i].
			rightDP[i][j] |= rightDP[i+1][j]
			// Take nums[i].
			if j+1 <= rightSize {
				mask := rightDP[i+1][j]
				var newMask uint64
				for or := 0; or < maxOr; or++ {
					if mask&(1<<or) != 0 {
						newMask |= 1 << (or | v)
					}
				}
				rightDP[i][j+1] |= newMask
			}
		}
	}

	// For each split point, combine left and right.
	// Left takes leftSize elements from prefix ending at split-1.
	// Right takes rightSize elements from suffix starting at split.
	ans := 0
	for split := leftSize; split <= n-rightSize; split++ {
		leftMask := leftDP[split][leftSize]
		rightMask := rightDP[split][rightSize]
		if leftMask == 0 || rightMask == 0 {
			continue
		}
		for lOr := 0; lOr < maxOr; lOr++ {
			if leftMask&(1<<lOr) == 0 {
				continue
			}
			for rOr := 0; rOr < maxOr; rOr++ {
				if rightMask&(1<<rOr) == 0 {
					continue
				}
				xor := lOr ^ rOr
				if xor > ans {
					ans = xor
				}
			}
		}
	}

	return ans
}
```
