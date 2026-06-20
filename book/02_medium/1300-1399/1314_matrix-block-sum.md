# 1314 — Matrix Block Sum

## Deskripsi

**Soal:** [1314. Matrix Block Sum](https://leetcode.com/problems/matrix-block-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n) where m,n are matrix dimensions  
**Kompleksitas Ruang:** O(m*n) for prefix sum matrix

**Algoritma:** Prefix Sum (jumlah kumulatif)

## Solusi Go

```go
package main

// LeetCode #1314: Matrix Block Sum
// https://leetcode.com/problems/matrix-block-sum/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	// [[12,21,16],[27,45,33],[24,39,28]]

	// Test case 2
	fmt.Println(matrixBlockSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 2))
	// [[45,45,45],[45,45,45],[45,45,45]]

	// Test case 3
	fmt.Println(matrixBlockSum([][]int{{1}}, 1))
	// [[1]]
}

// Time: O(m*n) where m,n are matrix dimensions
// Space: O(m*n) for prefix sum matrix
func matrixBlockSum(mat [][]int, k int) [][]int {
	m, n := len(mat), len(mat[0])

	// Build 2D prefix sum (1-indexed)
  // Membuat slice 2D untuk DP/tabel
	prefix := make([][]int, m+1)
  // Iterasi seluruh elemen
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = mat[i][j] + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
		}
	}

	// Calculate block sums
  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, m)
  // Iterasi seluruh elemen
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			r1, c1 := max(0, i-k), max(0, j-k)
			r2, c2 := min(m-1, i+k), min(n-1, j+k)
			result[i][j] = prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
