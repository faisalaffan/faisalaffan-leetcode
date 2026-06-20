# 2579 — Count Total Number Of Colored Cells

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func coloredCells(n int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2579: Count Total Number of Colored Cells
// https://leetcode.com/problems/count-total-number-of-colored-cells/
// Difficulty: Medium
// Time: O(1) | Space: O(1)

import "fmt"

func coloredCells(n int) int64 {
	// Formula: 1 + 2*n*(n-1) = 2n^2 - 2n + 1
	return int64(1 + 2*n*(n-1))
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", coloredCells(1))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", coloredCells(2))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", coloredCells(3))
	// Expected: 13
}
```
