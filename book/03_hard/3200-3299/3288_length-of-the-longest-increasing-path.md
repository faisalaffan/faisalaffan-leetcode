# 3288 — Length Of The Longest Increasing Path

## Deskripsi

**Soal:** [3288. Length Of The Longest Increasing Path](https://leetcode.com/problems/length-of-the-longest-increasing-path/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner), LIS (Longest Increasing Subsequence)

> **Ide Kunci:** //  1. Sort points by (x, y) ascending.

## Solusi Go

```go
package main

// LeetCode #3288: Length of the Longest Increasing Path
// https://leetcode.com/problems/length-of-the-longest-increasing-path/
// Difficulty: Hard
//
// Given an array of 2D points (x, y), find the length of the longest path
// where you can move from one point to another if both x and y strictly
// increase. This is equivalent to the Longest Increasing Subsequence (LIS)
// on y after sorting by x, handling duplicate x values carefully.
//
// Approach:
//  1. Sort points by (x, y) ascending.
//  2. Group points by x. Within each group, process y values in descending
//     order to prevent using two points with the same x in the path.
//  3. Use patience sorting (binary search on tails array) to find LIS length.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(longestIncreasingPath([][]int{{1, 2}, {2, 3}, {3, 4}}))
	// Example 2
	fmt.Println(longestIncreasingPath([][]int{{1, 1}, {2, 2}, {2, 3}, {3, 4}}))
	// Example 3: no valid path (x doesn't increase)
	fmt.Println(longestIncreasingPath([][]int{{3, 1}, {2, 2}, {1, 3}}))
	// Example 4: single point
	fmt.Println(longestIncreasingPath([][]int{{5, 5}}))
	// Example 5: points with same x
	fmt.Println(longestIncreasingPath([][]int{{1, 1}, {1, 2}, {1, 3}, {2, 4}}))
	// Example 6: complex
	fmt.Println(longestIncreasingPath([][]int{{1, 5}, {2, 3}, {3, 4}, {4, 2}, {5, 1}}))
}

func longestIncreasingPath(coordinates [][]int) int {
	n := len(coordinates)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Sort by x ascending, then y ascending.
	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i][0] != coordinates[j][0] {
			return coordinates[i][0] < coordinates[j][0]
		}
		return coordinates[i][1] < coordinates[j][1]
	})

	// Patience sorting (LIS) on y.
	// Process points grouped by x. Within each group, process y in descending
	// order to avoid taking two points from the same x.
  // Membuat slice untuk menyimpan hasil
	tails := make([]int, 0, n)

	i := 0
	for i < n {
		j := i
		// Find the group of points with the same x.
		for j < n && coordinates[j][0] == coordinates[i][0] {
			j++
		}

		// Process this group's y values in descending order,
		// so that same-x y values don't chain into each other.
		// For each y, find its position in the LIS tails.
  // Membuat slice untuk menyimpan hasil
		updates := make([]struct {
			pos int
			val int
		}, 0, j-i)
		for k := j - 1; k >= i; k-- {
			y := coordinates[k][1]
			// Find first tails[pos] >= y (lower_bound for strictly increasing).
			pos := lowerBound(tails, y)
			updates = append(updates, struct {
				pos int
				val int
			}{pos, y})
		}

		// Apply the updates.
		for _, upd := range updates {
			if upd.pos < len(tails) {
				if upd.val < tails[upd.pos] {
					tails[upd.pos] = upd.val
				}
			} else {
				tails = append(tails, upd.val)
			}
		}

		i = j
	}

	return len(tails)
}

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
```
