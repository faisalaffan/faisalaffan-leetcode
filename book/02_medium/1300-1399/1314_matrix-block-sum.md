# 1314 — Matrix Block Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func matrixBlockSum(mat [][]int, k int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(m*n) where m,n are matrix dimensions  
**Kompleksitas Ruang:** O(m*n) for prefix sum matrix

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Membuat matriks/slice 2D untuk DP
	prefix := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = mat[i][j] + prefix[i][j+1] + prefix[i+1][j] - prefix[i][j]
		}
	}

	// Calculate block sums
  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
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
