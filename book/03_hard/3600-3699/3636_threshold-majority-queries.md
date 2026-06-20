# 3636 — Threshold Majority Queries

## Deskripsi

**Soal:** [3636. Threshold Majority Queries](https://leetcode.com/problems/threshold-majority-queries/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner), LIS (Longest Increasing Subsequence)

> **Ide Kunci:** For each value, store sorted list of positions. For

## Solusi Go

```go
package main

// LeetCode #3636: Threshold Majority Queries
// https://leetcode.com/problems/threshold-majority-queries/
// Difficulty: Hard
//
// Given array nums, for each query [l, r, threshold], find the
// smallest value that appears more than threshold times in the
// subarray nums[l:r+1]. The candidate must appear > threshold times.
//
// Approach: For each value, store sorted list of positions. For
// each query, iterate over values with freq > threshold in the
// range using binary search to count occurrences.

import "fmt"
import "sort"

func main() {
	// Example 1
	fmt.Println(thresholdMajorityQueries([]int{1, 1, 2, 2, 1, 1}, [][]int{{0, 5, 2}, {2, 3, 1}}))
	// Example 2
	fmt.Println(thresholdMajorityQueries([]int{1, 2, 3, 4}, [][]int{{0, 3, 1}}))
	// Edge: single element
	fmt.Println(thresholdMajorityQueries([]int{5}, [][]int{{0, 0, 0}}))
}

func thresholdMajorityQueries(nums []int, queries [][]int) []int {
  // Membuat map untuk pencarian O(1): key → value
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r, threshold := q[0], q[1], q[2]
		best := -1

		// Try each value's positions
		for val, positions := range pos {
			cnt := sort.SearchInts(positions, r+1) - sort.SearchInts(positions, l)
			if cnt > threshold {
				if best == -1 || val < best {
					best = val
				}
			}
		}
		ans[qi] = best
	}

	return ans
}
```
