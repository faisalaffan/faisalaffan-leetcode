# 0462 — Minimum Moves To Equal Array Elements Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumMovesToEqualArrayElementsIi(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) for sorting, O(n) for QuickSelect  
**Kompleksitas Ruang:** O(log n) for sorting, O(1) for QuickSelect

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #462: Minimum Moves to Equal Array Elements II
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/
// Difficulty: Medium
// Time: O(n log n) for sorting, O(n) for QuickSelect
// Space: O(log n) for sorting, O(1) for QuickSelect

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 2, 3}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 10, 2, 9}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 0, 0, 8, 6}))
}

func MinimumMovesToEqualArrayElementsIi(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	median := nums[len(nums)/2]
	moves := 0
	for _, num := range nums {
		diff := num - median
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
```
