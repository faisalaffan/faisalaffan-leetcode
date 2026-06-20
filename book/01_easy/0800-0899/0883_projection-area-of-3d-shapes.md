# 0883 — Projection Area Of 3D Shapes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func projectionArea(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #883: Projection Area of 3D Shapes
// https://leetcode.com/problems/projection-area-of-3d-shapes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}})) // 17
	fmt.Println(projectionArea([][]int{{2}}))             // 5
	fmt.Println(projectionArea([][]int{{1, 0}, {0, 2}})) // 8
}

// projectionArea returns the total area of the 3D shape projection.
// Time: O(n^2). Space: O(1).
func projectionArea(grid [][]int) int {
	n := len(grid)
	xy := 0
	xz := 0
	yz := 0
	for i := 0; i < n; i++ {
		maxRow, maxCol := 0, 0
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				xy++
			}
			if grid[i][j] > maxRow {
				maxRow = grid[i][j]
			}
			if grid[j][i] > maxCol {
				maxCol = grid[j][i]
			}
		}
		xz += maxRow
		yz += maxCol
	}
	return xy + xz + yz
}
```
