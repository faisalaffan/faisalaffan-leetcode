# 3652 — Best Time To Buy And Sell Stock Using Strategy

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func bestTimeToBuyAndSellStockUsingStrategy(prices []int, strategy []int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3652: Best Time to Buy and Sell Stock using Strategy
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-using-strategy/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func bestTimeToBuyAndSellStockUsingStrategy(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	half := k / 2

	var base int64 = 0
	for i := 0; i < n; i++ {
		base += int64(prices[i] * strategy[i])
	}

	// sm = sum of prices in right half of window
	// so = sum of strategy[i]*prices[i] in full window
	var sm, so, diff int64 = 0, 0, 0
	for i := 0; i < n; i++ {
		sm += int64(prices[i])
		if i >= half {
			sm -= int64(prices[i-half])
		}

		so += int64(prices[i] * strategy[i])
		if i >= k {
			so -= int64(prices[i-k] * strategy[i-k])
		}

		if i+1 >= k {
			delta := sm - so
			if delta > diff {
				diff = delta
			}
		}
	}

	return base + diff
}

func main() {
	fmt.Println(bestTimeToBuyAndSellStockUsingStrategy([]int{1, 2, 3, 4}, []int{-1, 0, 1, 1}, 2))
	fmt.Println(bestTimeToBuyAndSellStockUsingStrategy([]int{5, 3, 2, 6}, []int{-1, 0, 1, -1}, 2))
	fmt.Println(bestTimeToBuyAndSellStockUsingStrategy([]int{1, 2, 3}, []int{-1, 0, 1}, 2))
}
```
