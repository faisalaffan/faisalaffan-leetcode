# 3189 — Minimum Moves To Get A Peaceful Board

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minMoves(rooks [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3189: Minimum Moves to Get a Peaceful Board
// https://leetcode.com/problems/minimum-moves-to-get-a-peaceful-board/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minMoves(rooks [][]int) int {
	n := len(rooks)
  // Alokasi slice integer
	rows := make([]int, n)
  // Alokasi slice integer
	cols := make([]int, n)
	for i, r := range rooks {
		rows[i] = r[0]
		cols[i] = r[1]
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(rows)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(cols)

	moves := 0
	for i := 0; i < n; i++ {
		moves += abs(rows[i] - i)
		moves += abs(cols[i] - i)
	}
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minMoves([][]int{{0, 0}, {1, 1}, {2, 2}}))    // Expected: 0
	fmt.Println(minMoves([][]int{{0, 0}, {0, 2}, {2, 0}}))    // Expected: 2
	fmt.Println(minMoves([][]int{{2, 2}, {0, 0}, {1, 1}}))    // Expected: 0
}
```
