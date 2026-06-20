# 0436 — Find Right Interval

## Deskripsi

**Soal:** [0436. Find Right Interval](https://leetcode.com/problems/find-right-interval/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func findRightInterval(intervals [][]int) []int`

## Solusi Go

```go
package main

// LeetCode #436: Find Right Interval
// https://leetcode.com/problems/find-right-interval/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	// Create array of (start, index) pairs
  // Membuat slice untuk menyimpan hasil
	starts := make([][2]int, n)
	for i, iv := range intervals {
		starts[i] = [2]int{iv[0], i}
	}
	sort.Slice(starts, func(i, j int) bool {
		return starts[i][0] < starts[j][0]
	})

  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	for i, iv := range intervals {
		target := iv[1]
		// Binary search
		idx := sort.Search(n, func(j int) bool {
			return starts[j][0] >= target
		})
		if idx < n {
			result[i] = starts[idx][1]
		} else {
			result[i] = -1
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findRightInterval([][]int{{1, 2}}))
	// Expected: [-1]

	// Test case 2
	fmt.Println("Test 2:", findRightInterval([][]int{{3, 4}, {2, 3}, {1, 2}}))
	// Expected: [-1, 0, 1]

	// Test case 3
	fmt.Println("Test 3:", findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
	// Expected: [-1, 2, -1]
}
```
