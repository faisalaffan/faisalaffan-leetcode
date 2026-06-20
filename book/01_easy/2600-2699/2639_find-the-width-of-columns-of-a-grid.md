# 2639 — Find The Width Of Columns Of A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheWidthOfColumnsOfAGrid(grid [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2639: Find the Width of Columns of a Grid
// https://leetcode.com/problems/find-the-width-of-columns-of-a-grid/
// Difficulty: Easy
// Time: O(m * n) | Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{1}, {22}, {333}}))
	fmt.Println(FindTheWidthOfColumnsOfAGrid([][]int{{-15, 1, 3}, {15, 7, 12}, {5, 6, -2}}))
}

func FindTheWidthOfColumnsOfAGrid(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
  // Alokasi slice integer
	ans := make([]int, n)
	for j := 0; j < n; j++ {
		maxLen := 0
		for i := 0; i < m; i++ {
			width := len(strconv.Itoa(grid[i][j]))
			if width > maxLen {
				maxLen = width
			}
		}
		ans[j] = maxLen
	}
	return ans
}
```
