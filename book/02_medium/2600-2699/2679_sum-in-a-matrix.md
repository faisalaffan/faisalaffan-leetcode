# 2679 — Sum In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func matrixSum(nums [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2679: Sum in a Matrix
// https://leetcode.com/problems/sum-in-a-matrix/
// Difficulty: Medium
// Time: O(m * n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	m := len(nums)
	if m == 0 {
		return 0
	}
	n := len(nums[0])

	// Sort each row
	for i := 0; i < m; i++ {
  // Urutkan secara ascending — O(n log n)
		sort.Ints(nums[i])
	}

	ans := 0
	for col := 0; col < n; col++ {
		maxVal := 0
		for row := 0; row < m; row++ {
			if nums[row][col] > maxVal {
				maxVal = nums[row][col]
			}
		}
		ans += maxVal
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", matrixSum([][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}))
	// Expected: 15

	// Test case 2
	fmt.Println("Test 2:", matrixSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// Expected: 18

	// Test case 3
	fmt.Println("Test 3:", matrixSum([][]int{{1}}))
	// Expected: 1
}
```
