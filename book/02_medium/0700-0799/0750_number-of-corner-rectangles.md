# 0750 — Number Of Corner Rectangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func countCornerRectangles(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(R * C^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #750: Number of Corner Rectangles
// https://leetcode.com/problems/number-of-corner-rectangles/
// Difficulty: Medium [Paid]
// Time: O(R * C^2)
// Space: O(1)

import "fmt"

func main() {
	grid := [][]int{
		{1, 0, 0, 1, 0},
		{0, 0, 1, 0, 1},
		{0, 0, 0, 1, 0},
		{1, 0, 1, 0, 1},
	}
	fmt.Println(countCornerRectangles(grid))
}

func countCornerRectangles(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	for c1 := 0; c1 < cols; c1++ {
		for c2 := c1 + 1; c2 < cols; c2++ {
			pairs := 0
			for r := 0; r < rows; r++ {
				if grid[r][c1] == 1 && grid[r][c2] == 1 {
					pairs++
				}
			}
			count += pairs * (pairs - 1) / 2
		}
	}

	return count
}
```
