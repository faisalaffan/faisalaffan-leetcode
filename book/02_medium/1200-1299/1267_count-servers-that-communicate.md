# 1267 — Count Servers That Communicate

## Deskripsi

**Soal:** [1267. Count Servers That Communicate](https://leetcode.com/problems/count-servers-that-communicate/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m+n)

**Algoritma:** —

**Fungsi Solusi:** `func countServers(grid [][]int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1267: Count Servers that Communicate
// https://leetcode.com/problems/count-servers-that-communicate/
// Difficulty: Medium

// Count servers that can communicate with at least one other server
// in the same row or column.

// Time: O(m*n)
// Space: O(m+n)

func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
  // Membuat slice untuk menyimpan hasil
	rowCount := make([]int, m)
  // Membuat slice untuk menyimpan hasil
	colCount := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCount[i]++
				colCount[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("%d (expected: 3)\n",
		countServers([][]int{{1, 0}, {0, 1}}))

	fmt.Printf("%d (expected: 4)\n",
		countServers([][]int{{1, 0}, {1, 1}}))

	fmt.Printf("%d (expected: 0)\n",
		countServers([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}))
}
```
