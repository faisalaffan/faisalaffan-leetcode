# 0950 — Reveal Cards In Increasing Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func deckRevealedIncreasing(deck []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #950: Reveal Cards In Increasing Order
// https://leetcode.com/problems/reveal-cards-in-increasing-order/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(n)
func deckRevealedIncreasing(deck []int) []int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(deck)
	n := len(deck)
  // Alokasi slice integer
	q := make([]int, n)
  // Range loop: iterasi dengan indeks + nilai
	for i := range q {
		q[i] = i
	}

  // Alokasi slice integer
	ans := make([]int, n)
	for _, card := range deck {
		idx := q[0]
		q = q[1:]
		ans[idx] = card
		if len(q) > 0 {
			q = append(q, q[0])
			q = q[1:]
		}
	}
	return ans
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
	fmt.Println(deckRevealedIncreasing([]int{1, 1000}))
}
```
