# 0892 — Surface Area Of 3D Shapes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func surfaceArea(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #892: Surface Area of 3D Shapes
// https://leetcode.com/problems/surface-area-of-3d-shapes/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(surfaceArea([][]int{{2}}))                                   // 10
	fmt.Println(surfaceArea([][]int{{1, 2}, {3, 4}}))                       // 34
}

// surfaceArea calculates the surface area of a 3D shape on a grid.
// Time: O(n^2). Space: O(1).
func surfaceArea(grid [][]int) int {
	n := len(grid)
	area := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				// Top + bottom
				area += 2
				// Four sides - subtract adjacent overlaps
				area += grid[i][j] * 4
				if i > 0 {
					area -= min(grid[i][j], grid[i-1][j]) * 2
				}
				if j > 0 {
					area -= min(grid[i][j], grid[i][j-1]) * 2
				}
			}
		}
	}
	return area
}
```
