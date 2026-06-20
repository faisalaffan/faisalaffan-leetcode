# 3111 — Minimum Rectangles To Cover Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minRectanglesToCoverPoints(points [][]int, w int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3111: Minimum Rectangles to Cover Points
// https://leetcode.com/problems/minimum-rectangles-to-cover-points/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) int {
  // Alokasi slice integer
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(xs)

	ans := 0
	i := 0
	for i < len(xs) {
		ans++
		end := xs[i] + w
		for i < len(xs) && xs[i] <= end {
			i++
		}
	}
	return ans
}

func main() {
	fmt.Println(minRectanglesToCoverPoints([][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1)) // Expected: 2
	fmt.Println(minRectanglesToCoverPoints([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, 2)) // Expected: 2
}
```
