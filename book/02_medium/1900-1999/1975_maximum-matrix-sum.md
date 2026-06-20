# 1975 — Maximum Matrix Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxMatrixSum(matrix [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1975: Maximum Matrix Sum
// https://leetcode.com/problems/maximum-matrix-sum/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaxMatrixSum([][]int{{1, -1}, {-1, 1}}))
	fmt.Println(MaxMatrixSum([][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}))
}

// Time: O(m*n), Space: O(1)
func MaxMatrixSum(matrix [][]int) int64 {
	total := int64(0)
	negCount := 0
	minAbs := int64(1 << 31)

	for _, row := range matrix {
		for _, val := range row {
			if val < 0 {
				negCount++
			}
			abs := int64(val)
			if abs < 0 {
				abs = -abs
			}
			total += abs
			if abs < minAbs {
				minAbs = abs
			}
		}
	}

	if negCount%2 == 1 {
		total -= 2 * minAbs
	}
	return total
}
```
