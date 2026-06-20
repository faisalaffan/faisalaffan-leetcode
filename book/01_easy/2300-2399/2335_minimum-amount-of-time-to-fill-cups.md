# 2335 — Minimum Amount Of Time To Fill Cups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumAmountOfTimeToFillCups(amount []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2335: Minimum Amount of Time to Fill Cups
// https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups/
// Difficulty: Easy
// Time O(1) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumAmountOfTimeToFillCups([]int{1, 4, 2})) // 4
	fmt.Println(MinimumAmountOfTimeToFillCups([]int{5, 4, 4})) // 5
}

func MinimumAmountOfTimeToFillCups(amount []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(amount)
	a, b, c := amount[0], amount[1], amount[2]
	if a+b <= c {
		return c
	}
	return (a+b+c+1)/2
}
```
