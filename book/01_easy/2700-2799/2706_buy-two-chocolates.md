# 2706 — Buy Two Chocolates

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BuyTwoChocolates(prices []int, money int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

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
  // Sort O(n log n)
	sort.Ints(prices)
	if len(prices) >= 2 && prices[0]+prices[1] <= money {
		return money - prices[0] - prices[1]
	}
	return money
}
```
