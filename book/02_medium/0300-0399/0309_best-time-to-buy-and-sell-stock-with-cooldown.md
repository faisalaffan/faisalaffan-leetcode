# 0309 — Best Time To Buy And Sell Stock With Cooldown

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxProfit(prices []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #309: Best Time to Buy and Sell Stock with Cooldown
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	hold := -prices[0]
	cool := 0
	sell := 0

	for i := 1; i < len(prices); i++ {
		prevSell := sell
		sell = max(sell, hold+prices[i])
		hold = max(hold, cool-prices[i])
		cool = max(cool, prevSell)
	}

	return max(sell, cool)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
	fmt.Println(maxProfit([]int{1}))
	fmt.Println(maxProfit([]int{1, 2, 4}))
}
```
