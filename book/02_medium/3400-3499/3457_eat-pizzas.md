# 3457 — Eat Pizzas

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxWeight(pizzas []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n) Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3457: Eat Pizzas!
// https://leetcode.com/problems/eat-pizzas/
// Difficulty: Medium
// Time: O(n log n) Space: O(log n)

import (
	"fmt"
	"sort"
)

func maxWeight(pizzas []int) int64 {
  // Custom sort dengan comparator
	sort.Slice(pizzas, func(i, j int) bool {
		return pizzas[i] > pizzas[j]
	})
	days := len(pizzas) / 4
	odd := (days + 1) / 2
	ans := int64(0)
	for i := 0; i < odd; i++ {
		ans += int64(pizzas[i])
	}
	for i := odd + 1; i < odd+days/2*2; i += 2 {
		ans += int64(pizzas[i])
	}
	return ans
}

func main() {
	fmt.Println(maxWeight([]int{1, 2, 3, 4, 5, 6, 7, 8})) // 14
	fmt.Println(maxWeight([]int{2, 2, 2, 2, 2, 2, 2, 2})) // 8
}
```
