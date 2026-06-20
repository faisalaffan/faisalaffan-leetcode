# 1222 — Queens That Can Attack The King

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func queensAttacktheKing(queens [][]int, king []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n) where n = number of queens  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1222: Queens That Can Attack the King
// https://leetcode.com/problems/queens-that-can-attack-the-king/
// Difficulty: Medium

// Queens can attack the king if no other piece blocks.
// From king's position, check 8 directions for nearest queen.

// Time: O(n) where n = number of queens
// Space: O(1)

func queensAttacktheKing(queens [][]int, king []int) [][]int {
  // Membuat map (HashMap) — pencarian O(1)
	queenSet := make(map[[2]int]bool)
	for _, q := range queens {
		queenSet[[2]int{q[0], q[1]}] = true
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

  // Membuat matriks/slice 2D untuk DP
	result := make([][]int, 0)

	for _, d := range dirs {
		r, c := king[0]+d[0], king[1]+d[1]
		for r >= 0 && r < 8 && c >= 0 && c < 8 {
			if queenSet[[2]int{r, c}] {
				result = append(result, []int{r, c})
				break
			}
			r += d[0]
			c += d[1]
		}
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [[0 1] [1 0] [3 3]])\n",
		queensAttacktheKing([][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}},
			[]int{0, 0}))

	fmt.Printf("%v (expected: [[2 2] [4 2]])\n",
		queensAttacktheKing([][]int{{0, 0}, {2, 2}, {4, 2}, {5, 5}},
			[]int{3, 2}))
}
```
