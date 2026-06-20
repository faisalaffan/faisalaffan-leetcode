# 1139 — Largest 1 Bordered Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func largest1BorderedSquare(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n * min(m, n))  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1139: Largest 1-Bordered Square
// https://leetcode.com/problems/largest-1-bordered-square/
// Difficulty: Medium

// Precompute horizontal/vertical consecutive 1s, then check each cell as
// bottom-right corner.

// Time: O(m * n * min(m, n))
// Space: O(m * n)

func largest1BorderedSquare(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
  // Membuat matriks/slice 2D untuk DP
	hor := make([][]int, m)
  // Membuat matriks/slice 2D untuk DP
	ver := make([][]int, m)
	for i := 0; i < m; i++ {
		hor[i] = make([]int, n)
		ver[i] = make([]int, n)
	}
	maxSide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				if j == 0 {
					hor[i][j] = 1
				} else {
					hor[i][j] = hor[i][j-1] + 1
				}
				if i == 0 {
					ver[i][j] = 1
				} else {
					ver[i][j] = ver[i-1][j] + 1
				}
				minSide := hor[i][j]
				if ver[i][j] < minSide {
					minSide = ver[i][j]
				}
				for s := minSide; s > maxSide; s-- {
					if hor[i-s+1][j] >= s && ver[i][j-s+1] >= s {
						maxSide = s
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func main() {
	fmt.Printf("%d (expected: 9)\n", largest1BorderedSquare([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Printf("%d (expected: 1)\n", largest1BorderedSquare([][]int{{1, 1, 0, 0}}))
	fmt.Printf("%d (expected: 0)\n", largest1BorderedSquare([][]int{{0}}))
}
```
