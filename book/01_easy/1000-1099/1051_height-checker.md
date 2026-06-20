# 1051 — Height Checker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func heightChecker(heights []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1051: Height Checker
// https://leetcode.com/problems/height-checker/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3})) // 3
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))    // 5
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 5}))    // 0
}

// LeetCode submission: heightChecker
func heightChecker(heights []int) int {
  // Alokasi slice integer
	expected := make([]int, len(heights))
	copy(expected, heights)
  // Urutkan secara ascending — O(n log n)
	sort.Ints(expected)
	count := 0
  // Range loop: iterasi dengan indeks + nilai
	for i := range heights {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}
```
