# 3195 — Find The Minimum Area To Cover All Ones I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumArea(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3195: Find the Minimum Area to Cover All Ones I
// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-i/
// Difficulty: Medium
// Time: O(m * n) | Space: O(1)

import "fmt"

func minimumArea(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	top, bottom, left, right := m, -1, n, -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if i < top {
					top = i
				}
				if i > bottom {
					bottom = i
				}
				if j < left {
					left = j
				}
				if j > right {
					right = j
				}
			}
		}
	}

	if bottom == -1 {
		return 0
	}
	return (bottom - top + 1) * (right - left + 1)
}

func main() {
	fmt.Println(minimumArea([][]int{{0, 1, 0}, {1, 0, 1}})) // Expected: 4
	fmt.Println(minimumArea([][]int{{0, 0}, {1, 1}}))        // Expected: 2
}
```
