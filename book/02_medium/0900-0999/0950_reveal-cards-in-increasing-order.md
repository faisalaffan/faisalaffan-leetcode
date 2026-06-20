# 0950 — Reveal Cards In Increasing Order

## Deskripsi

**Soal:** [0950. Reveal Cards In Increasing Order](https://leetcode.com/problems/reveal-cards-in-increasing-order/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func deckRevealedIncreasing(deck []int) []int`

## Solusi Go

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
	sort.Ints(deck)
	n := len(deck)
  // Membuat slice untuk menyimpan hasil
	q := make([]int, n)
  // Iterasi seluruh elemen
	for i := range q {
		q[i] = i
	}

  // Membuat slice untuk menyimpan hasil
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
