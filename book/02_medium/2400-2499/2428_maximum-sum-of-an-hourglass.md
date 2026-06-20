# 2428 — Maximum Sum Of An Hourglass

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSum(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2428: Maximum Sum of an Hourglass
// https://leetcode.com/problems/maximum-sum-of-an-hourglass/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)
// Sliding hourglass over the grid.

import "fmt"

func main() {
	fmt.Println(maxSum([][]int{{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9}})) // 30
	fmt.Println(maxSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))                         // 35
}

func maxSum(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	ans := 0
	for i := 0; i+2 < r; i++ {
		for j := 0; j+2 < c; j++ {
			sum := grid[i][j] + grid[i][j+1] + grid[i][j+2] + // top row
				grid[i+1][j+1] + // middle
				grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2] // bottom row
			if sum > ans {
				ans = sum
			}
		}
	}
	return ans
}
```
