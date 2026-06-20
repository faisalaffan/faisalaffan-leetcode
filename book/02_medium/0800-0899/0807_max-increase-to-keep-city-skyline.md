# 0807 — Max Increase To Keep City Skyline

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxIncreaseKeepingSkyline(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #807: Max Increase to Keep City Skyline
// https://leetcode.com/problems/max-increase-to-keep-city-skyline/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	fmt.Println(maxIncreaseKeepingSkyline(grid))
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
  // Alokasi slice integer
	rowMax := make([]int, n)
  // Alokasi slice integer
	colMax := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > rowMax[i] {
				rowMax[i] = grid[i][j]
			}
			if grid[i][j] > colMax[j] {
				colMax[j] = grid[i][j]
			}
		}
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			limit := min(rowMax[i], colMax[j])
			total += limit - grid[i][j]
		}
	}

	return total
}
```
