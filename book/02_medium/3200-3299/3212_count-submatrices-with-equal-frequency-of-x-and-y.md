# 3212 — Count Submatrices With Equal Frequency Of X And Y

## Deskripsi

**Soal:** [3212. Count Submatrices With Equal Frequency Of X And Y](https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

**Algoritma:** —

**Fungsi Solusi:** `func numberOfSubmatrices(grid [][]byte) int`

## Solusi Go

```go
package main

// LeetCode #3212: Count Submatrices With Equal Frequency of X and Y
// https://leetcode.com/problems/count-submatrices-with-equal-frequency-of-x-and-y/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func numberOfSubmatrices(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat slice 2D untuk DP/tabel
	prefX := make([][]int, m+1)
  // Membuat slice 2D untuk DP/tabel
	prefY := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range prefX {
		prefX[i] = make([]int, n+1)
		prefY[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefX[i+1][j+1] = prefX[i][j+1] + prefX[i+1][j] - prefX[i][j]
			prefY[i+1][j+1] = prefY[i][j+1] + prefY[i+1][j] - prefY[i][j]
			if grid[i][j] == 'X' {
				prefX[i+1][j+1]++
			} else if grid[i][j] == 'Y' {
				prefY[i+1][j+1]++
			}
		}
	}

	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			x := prefX[i][j]
			y := prefY[i][j]
			if x > 0 && x == y {
				ans++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'Y', '.'}, {'Y', '.', '.'}})) // Expected: 3
	fmt.Println(numberOfSubmatrices([][]byte{{'X', 'X'}, {'Y', 'Y'}}))          // Expected: 0
}
```
