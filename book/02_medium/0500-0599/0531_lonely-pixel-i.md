# 0531 — Lonely Pixel I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func FindLonelyPixel(picture [][]byte) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(m * n)  |  **Ruang:** O(m + n)


## 💻 Solusi Go

```go
package main

// LeetCode #531: Lonely Pixel I
// https://leetcode.com/problems/lonely-pixel-i/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m + n)

import "fmt"

func main() {
	picture := [][]byte{
		{'W', 'W', 'B'},
		{'W', 'B', 'W'},
		{'B', 'W', 'W'},
	}
	fmt.Println(FindLonelyPixel(picture))
}

func FindLonelyPixel(picture [][]byte) int {
	m, n := len(picture), len(picture[0])
  // Alokasi slice
	rows := make([]int, m)
  // Alokasi slice
	cols := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && rows[i] == 1 && cols[j] == 1 {
				count++
			}
		}
	}

	return count
}
```
