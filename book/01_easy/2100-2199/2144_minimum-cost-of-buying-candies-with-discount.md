# 2144 — Minimum Cost Of Buying Candies With Discount

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumCostOfBuyingCandiesWithDiscount(cost []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(1) ignoring sort  
**Kompleksitas Ruang:** O(1) ignoring sort

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2144: Minimum Cost of Buying Candies With Discount
// https://leetcode.com/problems/minimum-cost-of-buying-candies-with-discount/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{1, 2, 3}))    // 5
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{6, 5, 7, 9, 2, 2})) // 23
	fmt.Println(MinimumCostOfBuyingCandiesWithDiscount([]int{5, 5}))        // 10
}

// Time: O(n log n), Space: O(1) ignoring sort
func MinimumCostOfBuyingCandiesWithDiscount(cost []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cost)))
	total := 0
	for i, c := range cost {
		if i%3 != 2 { // buy 2, get 1 free (the cheapest = every 3rd item)
			total += c
		}
	}
	return total
}
```
