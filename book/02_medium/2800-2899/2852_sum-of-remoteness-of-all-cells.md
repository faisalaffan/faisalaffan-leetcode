# 2852 — Sum Of Remoteness Of All Cells

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func SumOfRemotenessOfAllCells(grid [][]int) []int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2852: Sum of Remoteness of All Cells
// https://leetcode.com/problems/sum-of-remoteness-of-all-cells/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func SumOfRemotenessOfAllCells(grid [][]int) []int64 {
	n := len(grid)
  // Membuat matriks/slice 2D untuk DP
	visited := make([][]bool, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var totalSum int64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				totalSum += int64(grid[i][j])
			}
		}
	}

  // Alokasi slice integer
	result := make([]int64, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				result[i*n+j] = totalSum - int64(grid[i][j])
			}
		}
	}

	return result
}

func main() {
	fmt.Println(SumOfRemotenessOfAllCells([][]int{{1, 2}, {3, 4}}))
	fmt.Println(SumOfRemotenessOfAllCells([][]int{{5, 0}, {0, 5}}))
}
```
