# 0888 — Fair Candy Swap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func fairCandySwap(aliceSizes []int, bobSizes []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n + m). Space: O(m).  |  **Ruang:** O(m).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #888: Fair Candy Swap
// https://leetcode.com/problems/fair-candy-swap/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(fairCandySwap([]int{1, 1}, []int{2, 2}))       // [1,2]
	fmt.Println(fairCandySwap([]int{1, 2}, []int{2, 3}))       // [1,2]
	fmt.Println(fairCandySwap([]int{2}, []int{1, 3}))          // [2,3]
}

// fairCandySwap finds a candy swap that makes both Alice and Bob have equal candy.
// Time: O(n + m). Space: O(m).
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumA, sumB := 0, 0
  // HashMap: O(1) lookup
	bSet := make(map[int]bool)
	for _, v := range aliceSizes {
		sumA += v
	}
	for _, v := range bobSizes {
		sumB += v
		bSet[v] = true
	}
	diff := (sumB - sumA) / 2
	for _, a := range aliceSizes {
		if bSet[a+diff] {
			return []int{a, a + diff}
		}
	}
	return nil
}
```
