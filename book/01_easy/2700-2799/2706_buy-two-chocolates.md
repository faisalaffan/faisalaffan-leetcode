# 2706 — Buy Two Chocolates

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BuyTwoChocolates(prices []int, money int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2706: Buy Two Chocolates
// https://leetcode.com/problems/buy-two-chocolates/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BuyTwoChocolates([]int{1, 2, 2}, 3))
	fmt.Println(BuyTwoChocolates([]int{3, 2, 3}, 3))
}

func BuyTwoChocolates(prices []int, money int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(prices)
	if len(prices) >= 2 && prices[0]+prices[1] <= money {
		return money - prices[0] - prices[1]
	}
	return money
}
```
