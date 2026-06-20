# 2517 — Maximum Tastiness Of Candy Basket

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumTastiness(price []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting

**Waktu:** O(n log n + n log max)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2517: Maximum Tastiness of Candy Basket
// https://leetcode.com/problems/maximum-tastiness-of-candy-basket/
// Difficulty: Medium
// Time: O(n log n + n log max) | Space: O(1)
// Binary search on answer. Check: can we pick k candies with min diff >= mid?

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3)) // 8
	fmt.Println(maximumTastiness([]int{1, 3, 1}, 2))             // 2
}

func maximumTastiness(price []int, k int) int {
  // Sort O(n log n)
	sort.Ints(price)
	lo, hi := 0, price[len(price)-1]-price[0]

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canPick(price, k, mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func canPick(price []int, k, minDiff int) bool {
	count, last := 1, price[0]
	for i := 1; i < len(price); i++ {
		if price[i]-last >= minDiff {
			count++
			last = price[i]
		}
	}
	return count >= k
}
```
