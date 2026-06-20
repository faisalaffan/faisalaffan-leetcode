# 2279 — Maximum Bags With Full Capacity Of Rocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumBags(capacity []int, rocks []int, additionalRocks int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2279: Maximum Bags With Full Capacity of Rocks
// https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	n := len(capacity)
  // Alokasi slice integer
	needed := make([]int, n)
	for i := 0; i < n; i++ {
		needed[i] = capacity[i] - rocks[i]
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(needed)

	full := 0
	for _, n := range needed {
  // Edge case: input kosong — langsung return
		if n == 0 {
			full++
		} else if n <= additionalRocks {
			additionalRocks -= n
			full++
		} else {
			break
		}
	}
	return full
}

func main() {
	// Test case 1
	fmt.Println(maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2))
	// Expected: 3

	// Test case 2
	fmt.Println(maximumBags([]int{10, 2, 2}, []int{2, 2, 0}, 100))
	// Expected: 3
}
```
