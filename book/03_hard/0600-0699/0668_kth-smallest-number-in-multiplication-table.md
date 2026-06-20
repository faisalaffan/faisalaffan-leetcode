# 0668 — Kth Smallest Number In Multiplication Table

## Deskripsi

**Soal:** [0668. Kth Smallest Number In Multiplication Table](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #668: Kth Smallest Number in Multiplication Table
// https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/
// Difficulty: Hard

import "fmt"

func main() {
	// Test cases
	testCases := []struct {
		m    int
		n    int
		k    int
		want int
	}{
		{3, 3, 5, 3},
		{2, 3, 6, 6},
		{1, 1, 1, 1},
		{2, 2, 1, 1},
		{2, 2, 2, 2},
		{2, 2, 3, 2},
		{2, 2, 4, 4},
		{3, 3, 1, 1},
		{3, 3, 9, 9},
		{5, 5, 10, 5},
		{10, 10, 75, 45},
		{42, 34, 401, 126},
	}

	for _, tc := range testCases {
		got := findKthNumber(tc.m, tc.n, tc.k)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: findKthNumber(%d, %d, %d) = %d (want %d)\n", status, tc.m, tc.n, tc.k, got, tc.want)
	}
}

func findKthNumber(m int, n int, k int) int {
	left, right := 1, m*n

  // Loop two-pointer: kiri vs kanan
	for left < right {
		mid := left + (right-left)/2
		count := countLessOrEqual(m, n, mid)
		if count >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func countLessOrEqual(m, n, x int) int {
	count := 0
	for i := 1; i <= m; i++ {
		count += min(x/i, n)
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
