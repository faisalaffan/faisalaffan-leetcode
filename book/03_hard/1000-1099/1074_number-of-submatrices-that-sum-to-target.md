# 1074 — Number Of Submatrices That Sum To Target

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func numSubmatrixSumTarget(matrix [][]int, target int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1074: Number of Submatrices That Sum to Target
// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/
// Difficulty: Hard
//
// Fix top row, expand bottom row, accumulate column-wise sums, then for each
// row-pair use prefix sum map (same as subarray sum equals target) across
// columns to count submatrices.

import "fmt"

func main() {
	matrix := [][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}
	fmt.Println(numSubmatrixSumTarget(matrix, 0))
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	ans := 0

	for top := 0; top < m; top++ {
  // Alokasi slice integer
		colSum := make([]int, n)
		for bottom := top; bottom < m; bottom++ {
			for c := 0; c < n; c++ {
				colSum[c] += matrix[bottom][c]
			}
			// Count subarrays in colSum that sum to target
			countMap := map[int]int{0: 1}
			prefix := 0
			for c := 0; c < n; c++ {
				prefix += colSum[c]
				ans += countMap[prefix-target]
				countMap[prefix]++
			}
		}
	}

	return ans
}
```
