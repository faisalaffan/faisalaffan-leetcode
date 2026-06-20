# 3128 — Right Triangles

## Deskripsi

**Soal:** [3128. Right Triangles](https://leetcode.com/problems/right-triangles/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m + n)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfRightTriangles(grid [][]int) int64`

## Solusi Go

```go
package main

// LeetCode #3128: Right Triangles
// https://leetcode.com/problems/right-triangles/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m + n)

import "fmt"

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat slice untuk menyimpan hasil
	rowSum := make([]int, m)
  // Membuat slice untuk menyimpan hasil
	colSum := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowSum[i]++
				colSum[j]++
			}
		}
	}

	var ans int64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans += int64(rowSum[i]-1) * int64(colSum[j]-1)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfRightTriangles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}})) // Expected: 2
	fmt.Println(numberOfRightTriangles([][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}})) // Expected: 2
}
```
