# 0593 — Valid Square

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ValidSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #593: Valid Square
// https://leetcode.com/problems/valid-square/
// Difficulty: Medium
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(ValidSquare([]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}))
	fmt.Println(ValidSquare([]int{0, 0}, []int{1, 2}, []int{2, 1}, []int{1, 0}))
}

func ValidSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	points := [][]int{p1, p2, p3, p4}
	dists := []int{}

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			dist := distSq(points[i], points[j])
			dists = append(dists, dist)
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(dists)

	// All 4 sides equal and non-zero, and 2 diagonals equal
	return dists[0] > 0 && dists[0] == dists[1] && dists[1] == dists[2] && dists[2] == dists[3] &&
		dists[4] == dists[5]
}

func distSq(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}
```
