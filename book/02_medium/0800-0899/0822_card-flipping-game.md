# 0822 — Card Flipping Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CardFlippingGame(fronts []int, backs []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #822: Card Flipping Game
// https://leetcode.com/problems/card-flipping-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CardFlippingGame([]int{1, 2, 4, 4, 7}, []int{1, 3, 4, 1, 3}))
	fmt.Println(CardFlippingGame([]int{1, 1}, []int{1, 2}))
	fmt.Println(CardFlippingGame([]int{1, 1}, []int{2, 2}))
}

// Time: O(n) | Space: O(n)
func CardFlippingGame(fronts []int, backs []int) int {
  // HashMap: O(1) lookup
	blocked := make(map[int]bool)
  // Range loop
	for i := range fronts {
		if fronts[i] == backs[i] {
			blocked[fronts[i]] = true
		}
	}

	ans := 2001
	for _, v := range fronts {
		if !blocked[v] && v < ans {
			ans = v
		}
	}
	for _, v := range backs {
		if !blocked[v] && v < ans {
			ans = v
		}
	}

	if ans == 2001 {
		return 0
	}
	return ans
}
```
