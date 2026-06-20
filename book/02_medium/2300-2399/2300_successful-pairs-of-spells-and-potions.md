# 2300 — Successful Pairs Of Spells And Potions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func successfulPairs(spells []int, potions []int, success int64) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O((n + m) log m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2300: Successful Pairs of Spells and Potions
// https://leetcode.com/problems/successful-pairs-of-spells-and-potions/
// Difficulty: Medium
// Time: O((n + m) log m) | Space: O(1)

import (
	"fmt"
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(potions)
	m := len(potions)
  // Alokasi slice integer
	result := make([]int, len(spells))

	for i, s := range spells {
		need := (int(success) + s - 1) / s // ceiling division
		idx := sort.SearchInts(potions, need)
		result[i] = m - idx
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(successfulPairs([]int{5, 1, 3}, []int{1, 2, 3, 4, 5}, 7))
	// Expected: [4, 0, 3]

	// Test case 2
	fmt.Println(successfulPairs([]int{3, 1, 2}, []int{8, 5, 8}, 16))
	// Expected: [2, 0, 2]
}
```
