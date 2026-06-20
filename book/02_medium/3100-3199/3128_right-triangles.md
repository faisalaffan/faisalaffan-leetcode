# 3128 — Right Triangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfRightTriangles(grid [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m + n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3128: Right Triangles
// https://leetcode.com/problems/right-triangles/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m + n)

import "fmt"

func numberOfRightTriangles(grid [][]int) int64 {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Alokasi slice integer
	rowSum := make([]int, m)
  // Alokasi slice integer
	colSum := make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowSum[i]++
				colSum[j]++
			}
		}
	}

	var ans int64
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ans += int64(rowSum[i]-1) * int64(colSum[j]-1)
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfRightTriangles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}})) // Expected: 2
	fmt.Println(numberOfRightTriangles([][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}})) // Expected: 2
}
```
