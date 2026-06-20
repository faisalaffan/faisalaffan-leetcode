# 1637 — Widest Vertical Area Between Two Points Containing No Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxWidthOfVerticalArea(points [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(log n) (for sorting)  
**Kompleksitas Ruang:** O(log n) (for sorting)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1637: Widest Vertical Area Between Two Points Containing No Points
// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func MaxWidthOfVerticalArea(points [][]int) int {
  // Alokasi slice integer
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(xs)
	maxWidth := 0
	for i := 1; i < len(xs); i++ {
		if xs[i]-xs[i-1] > maxWidth {
			maxWidth = xs[i] - xs[i-1]
		}
	}
	return maxWidth
}

func main() {
	fmt.Println(MaxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}))
	fmt.Println(MaxWidthOfVerticalArea([][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}))
}
```
