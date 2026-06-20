# 1402 — Reducing Dishes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSatisfaction(satisfaction []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1402: Reducing Dishes
// https://leetcode.com/problems/reducing-dishes/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	// Sort descending
  // Custom sort dengan comparator
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})

	curr := 0
	maxVal := 0

	for _, s := range satisfaction {
		curr += s
		if curr > 0 {
			maxVal += curr
		}
	}

	return maxVal
}

func main() {
	// Example 1
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	// Expected: 14

	// Example 2
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	// Expected: 20

	// Example 3
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
	// Expected: 0
}
```
