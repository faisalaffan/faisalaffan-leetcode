# 2035 — Partition Array Into Two Arrays To Minimize Sum Difference

## Deskripsi

**Soal:** [2035. Partition Array Into Two Arrays To Minimize Sum Difference](https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func minimumDifference(nums []int) int`

> **Ide Kunci:** Meet-in-the-Middle

## Solusi Go

```go
package main

// LeetCode #2035: Partition Array Into Two Arrays to Minimize Sum Difference
// https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/
// Difficulty: Hard
// Approach: Meet-in-the-Middle

import (
	"fmt"
	"math"
	"sort"
)

func minimumDifference(nums []int) int {
	n := len(nums) / 2 // each half size
	total := 0
	for _, v := range nums {
		total += v
	}

	// Generate all subset sums for each half, grouped by subset size
  // Membuat slice 2D untuk DP/tabel
	leftSums := make([][]int, n+1)
  // Membuat slice 2D untuk DP/tabel
	rightSums := make([][]int, n+1)

	// Generate combinations for left half
	for mask := 0; mask < (1 << n); mask++ {
		sum := 0
		size := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += nums[i]
				size++
			}
		}
		leftSums[size] = append(leftSums[size], sum)
	}

	// Generate combinations for right half
	for mask := 0; mask < (1 << n); mask++ {
		sum := 0
		size := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum += nums[n+i]
				size++
			}
		}
		rightSums[size] = append(rightSums[size], sum)
	}

	// Sort each group in right half for binary search
	for i := 0; i <= n; i++ {
		sort.Ints(rightSums[i])
	}

	ans := math.MaxInt32

	// For each possible size from left half, find complement from right half
	for leftSize := 0; leftSize <= n; leftSize++ {
		rightSize := n - leftSize
		for _, leftSum := range leftSums[leftSize] {
			// We want leftSum + rightSum as close to total/2 as possible
			target := total/2 - leftSum
			rightArr := rightSums[rightSize]
			if len(rightArr) == 0 {
				continue
			}
			// Binary search for closest
			idx := sort.SearchInts(rightArr, target)
			if idx < len(rightArr) {
				sum := leftSum + rightArr[idx]
				diff := total - 2*sum
				if diff < 0 {
					diff = -diff
				}
				if diff < ans {
					ans = diff
				}
			}
			if idx > 0 {
				idx--
				sum := leftSum + rightArr[idx]
				diff := total - 2*sum
				if diff < 0 {
					diff = -diff
				}
				if diff < ans {
					ans = diff
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println("2035. Partition Array Into Two Arrays to Minimize Sum Difference")

	// Example 1
	nums1 := []int{3, 9, 7, 3}
	fmt.Printf("nums=%v → %d (expected 2)\n", nums1, minimumDifference(nums1))

	// Example 2
	nums2 := []int{-36, 36}
	fmt.Printf("nums=%v → %d (expected 72)\n", nums2, minimumDifference(nums2))

	// Example 3
	nums3 := []int{2, -1, 0, 4, -2, -9}
	fmt.Printf("nums=%v → %d (expected 0)\n", nums3, minimumDifference(nums3))
}
```
