# 1755 — Closest Subsequence Sum

## Deskripsi

**Soal:** [1755. Closest Subsequence Sum](https://leetcode.com/problems/closest-subsequence-sum/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func minAbs(x, y int) int`

> **Ide Kunci:** Meet-in-the-middle.

## Solusi Go

```go
package main

// LeetCode #1755: Closest Subsequence Sum
// https://leetcode.com/problems/closest-subsequence-sum/
// Difficulty: Hard
//
// Approach: Meet-in-the-middle.
// Split array into two halves. Generate all possible subset sums for each half.
// Sort the second half. For each sum in the first half, binary search the second
// half for the value closest to (goal - sum).

import (
	"fmt"
	"sort"
)

func minAbs(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if x < y {
		return x
	}
	return y
}

func min(vals ...int) int {
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func generateSums(arr []int) []int {
	n := len(arr)
	sums := []int{0}
	for i := 0; i < n; i++ {
		curLen := len(sums)
		for j := 0; j < curLen; j++ {
			sums = append(sums, sums[j]+arr[i])
		}
	}
	return sums
}

func minAbsDifference(nums []int, goal int) int {
	n := len(nums)
	mid := n / 2

	left := generateSums(nums[:mid])
	right := generateSums(nums[mid:])

	sort.Ints(right)

	best := 1 << 60
	for _, s := range left {
		target := goal - s
		// Binary search for closest
		idx := sort.SearchInts(right, target)
		if idx < len(right) {
			diff := target - right[idx]
			if diff < 0 {
				diff = -diff
			}
			if diff < best {
				best = diff
			}
		}
		if idx > 0 {
			diff := target - right[idx-1]
			if diff < 0 {
				diff = -diff
			}
			if diff < best {
				best = diff
			}
		}
		if best == 0 {
			break
		}
	}
	return best
}

func main() {
	// Example test case
	fmt.Println("nums=[5,-7,3,5],goal=6 →", minAbsDifference([]int{5, -7, 3, 5}, 6)) // Expected: 0

	// Additional tests
	fmt.Println("nums=[1,2,3],goal=10 →", minAbsDifference([]int{1, 2, 3}, 10))       // Expected: 4 (6 vs 10)
	fmt.Println("nums=[-1,-2,-3],goal=0 →", minAbsDifference([]int{-1, -2, -3}, 0))    // Expected: 0 (sum=0 via empty)
	fmt.Println("nums=[7,-9,15,-2],goal=-5 →", minAbsDifference([]int{7, -9, 15, -2}, -5)) // Expected: 0
}
```
