# 1964 — Find The Longest Valid Obstacle Course At Each Position

## Deskripsi

**Soal:** [1964. Find The Longest Valid Obstacle Course At Each Position](https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func longestObstacleCourseAtEachPosition(obstacles []int) []int`

## Solusi Go

```go
package main

// LeetCode #1964: Find the Longest Valid Obstacle Course at Each Position
// https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/
// Difficulty: Hard
// LIS variant with non-decreasing constraint.
// Patience sorting: binary search for first element > h, replace with h.

import (
	"fmt"
	"sort"
)

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, n)
  // Membuat slice untuk menyimpan hasil
	tails := make([]int, 0, n)

	for i, h := range obstacles {
		// Find first element in tails > h (strictly greater)
		// sort.Search finds first index where predicate is true
		idx := sort.Search(len(tails), func(k int) bool { return tails[k] > h })
		if idx == len(tails) {
			tails = append(tails, h)
		} else {
			tails[idx] = h
		}
		ans[i] = idx + 1
	}
	return ans
}

func main() {
	fmt.Println(longestObstacleCourseAtEachPosition([]int{1, 2, 3, 2}))          // Expected: [1 2 3 3]
	fmt.Println(longestObstacleCourseAtEachPosition([]int{2, 2, 1}))             // Expected: [1 2 1]
	fmt.Println(longestObstacleCourseAtEachPosition([]int{3, 1, 5, 6, 4, 2}))   // Expected: [1 1 2 3 2 2]
}
```
