# 3546 — Equal Sum Grid Partition I

## Deskripsi

**Soal:** [3546. Equal Sum Grid Partition I](https://leetcode.com/problems/equal-sum-grid-partition-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Prefix Sum (jumlah kumulatif)

## Solusi Go

```go
package main

// LeetCode #3546: Equal Sum Grid Partition I
// https://leetcode.com/problems/equal-sum-grid-partition-i/
// Difficulty: Medium
// Complexity: O(n*m) time, O(n*m) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Test 1:", EqualSumGridPartitionI(grid))
	// Test case 2
	grid2 := [][]int{{1, 1, 1}, {1, 1, 1}}
	fmt.Println("Test 2:", EqualSumGridPartitionI(grid2))
	// Test case 3
	grid3 := [][]int{{5}}
	fmt.Println("Test 3:", EqualSumGridPartitionI(grid3))
}

func EqualSumGridPartitionI(grid [][]int) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}
	m, n := len(grid), len(grid[0])
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			total += grid[i][j]
		}
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2

	// prefix sums for rows
  // Membuat slice untuk menyimpan hasil
	rowSum := make([]int, m)
	for i := 0; i < m; i++ {
		s := 0
		for j := 0; j < n; j++ {
			s += grid[i][j]
		}
		rowSum[i] = s
	}

	// Try horizontal split
	sum := 0
	for i := 0; i < m-1; i++ {
		sum += rowSum[i]
		if sum == target {
			return true
		}
	}

	// Try vertical split
	for j := 0; j < n-1; j++ {
		sum := 0
		for i := 0; i < m; i++ {
			sum += grid[i][j]
		}
		if sum == target {
			return true
		}
	}

	return false
}
```
