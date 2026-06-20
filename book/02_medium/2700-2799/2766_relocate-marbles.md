# 2766 — Relocate Marbles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2766: Relocate Marbles
// https://leetcode.com/problems/relocate-marbles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func RelocateMarbles(nums []int, moveFrom []int, moveTo []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	positions := make(map[int]bool)
	for _, n := range nums {
		positions[n] = true
	}

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(moveFrom); i++ {
		delete(positions, moveFrom[i])
		positions[moveTo[i]] = true
	}

  // Alokasi slice integer
	result := make([]int, 0, len(positions))
	for p := range positions {
		result = append(result, p)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(result)
	return result
}

func main() {
	fmt.Println(RelocateMarbles([]int{1, 2, 3}, []int{1}, []int{4}))
	fmt.Println(RelocateMarbles([]int{1, 1, 2, 2}, []int{1, 2}, []int{3, 4}))
}
```
