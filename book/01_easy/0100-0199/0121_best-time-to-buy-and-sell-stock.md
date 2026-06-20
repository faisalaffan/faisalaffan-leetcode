# 0121 — Best Time To Buy And Sell Stock

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxProfit(prices []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #121: Best Time to Buy and Sell Stock
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func MaxProfit(prices []int) int {
	minPrice, maxProfit := prices[0], 0
	for _, p := range prices[1:] {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}

func main() {
	fmt.Println(MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(MaxProfit([]int{7, 6, 4, 3, 1}))
}
```
