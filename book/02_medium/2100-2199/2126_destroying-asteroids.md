# 2126 — Destroying Asteroids

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func asteroidsDestroyed(mass int, asteroids []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2126: Destroying Asteroids
// https://leetcode.com/problems/destroying-asteroids/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(asteroids)
	current := int64(mass)

	for _, a := range asteroids {
		if current < int64(a) {
			return false
		}
		current += int64(a)
	}

	return true
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", asteroidsDestroyed(10, []int{3, 9, 19, 5, 21}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", asteroidsDestroyed(5, []int{4, 9, 23, 4}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", asteroidsDestroyed(1, []int{1, 1, 1, 1}))
	// Expected: true
}
```
