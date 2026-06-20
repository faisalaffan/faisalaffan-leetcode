# 0240 — Search A 2D Matrix Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func searchMatrix(matrix [][]int, target int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m+n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #240: Search a 2D Matrix II
// https://leetcode.com/problems/search-a-2d-matrix-ii/
// Difficulty: Medium
// Time: O(m+n), Space: O(1)

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1

	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(matrix, 5))
	fmt.Println(searchMatrix(matrix, 20))
	fmt.Println(searchMatrix(matrix, 30))
}
```
