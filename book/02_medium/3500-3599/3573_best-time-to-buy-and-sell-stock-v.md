# 3573 — Best Time To Buy And Sell Stock V

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BestTimeToBuyAndSellStockV(prices []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3573: Best Time to Buy and Sell Stock V
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-v/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", BestTimeToBuyAndSellStockV([]int{7, 1, 5, 3, 6, 4}))
	// Test case 2
	fmt.Println("Test 2:", BestTimeToBuyAndSellStockV([]int{7, 6, 4, 3, 1}))
	// Test case 3
	fmt.Println("Test 3:", BestTimeToBuyAndSellStockV([]int{1, 2, 3, 4, 5}))
}

func BestTimeToBuyAndSellStockV(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// With at most 2 transactions
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		if -prices[i] > buy1 {
			buy1 = -prices[i]
		}
		if buy1+prices[i] > sell1 {
			sell1 = buy1 + prices[i]
		}
		if sell1-prices[i] > buy2 {
			buy2 = sell1 - prices[i]
		}
		if buy2+prices[i] > sell2 {
			sell2 = buy2 + prices[i]
		}
	}
	return sell2
}
```
