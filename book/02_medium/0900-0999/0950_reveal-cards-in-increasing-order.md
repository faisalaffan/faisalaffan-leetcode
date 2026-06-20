# 0950 — Reveal Cards In Increasing Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func deckRevealedIncreasing(deck []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #950: Reveal Cards In Increasing Order
// https://leetcode.com/problems/reveal-cards-in-increasing-order/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(n)
func deckRevealedIncreasing(deck []int) []int {
  // Sort O(n log n)
	sort.Ints(deck)
	n := len(deck)
  // Alokasi slice
	q := make([]int, n)
  // Range loop
	for i := range q {
		q[i] = i
	}

  // Alokasi slice
	ans := make([]int, n)
	for _, card := range deck {
		idx := q[0]
		q = q[1:]
		ans[idx] = card
		if len(q) > 0 {
			q = append(q, q[0])
			q = q[1:]
		}
	}
	return ans
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
	fmt.Println(deckRevealedIncreasing([]int{1, 1000}))
}
```
