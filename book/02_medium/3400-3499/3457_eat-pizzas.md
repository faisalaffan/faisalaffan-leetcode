# 3457 — Eat Pizzas

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxWeight(pizzas []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n) Space: O(log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

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
  // Custom sort
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
