# 1033 — Moving Stones Until Consecutive

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numMovesStones(a int, b int, c int) []int
```

> **💡 Hint:** Sort the positions. Min moves = 0, 1, or 2 based on gaps.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1033: Moving Stones Until Consecutive
// https://leetcode.com/problems/moving-stones-until-consecutive/
// Difficulty: Medium
//
// Approach: Sort the positions. Min moves = 0, 1, or 2 based on gaps.
//           Max moves = distance between extremes - 2.
// Time: O(1)
// Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numMovesStones(1, 2, 5)) // [1,2]
	fmt.Println(numMovesStones(4, 3, 2)) // [0,0]
	fmt.Println(numMovesStones(3, 5, 1)) // [1,2]
}

func numMovesStones(a int, b int, c int) []int {
	stones := []int{a, b, c}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(stones)
	x, y, z := stones[0], stones[1], stones[2]

	minMoves := 2
	if z-x == 2 {
		minMoves = 0
	} else if z-y <= 2 || y-x <= 2 {
		minMoves = 1
	}

	maxMoves := (z - x - 2)

	return []int{minMoves, maxMoves}
}
```
