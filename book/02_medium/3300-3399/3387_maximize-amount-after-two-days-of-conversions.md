# 3387 — Maximize Amount After Two Days Of Conversions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n1 + n2) Space: O(currencies)  |  **Ruang:** O(currencies)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3387: Maximize Amount After Two Days of Conversions
// https://leetcode.com/problems/maximize-amount-after-two-days-of-conversions/
// Difficulty: Medium
// Time: O(n1 + n2) Space: O(currencies)

import "fmt"

func main() {
	fmt.Println(maxAmount("EUR", [][]string{{"EUR", "USD"}}, []float64{2.0}, [][]string{{"USD", "EUR"}}, []float64{0.5}))
	// 1.0 EUR -> 2.0 USD day1 -> 1.0 EUR day2 = 1.0
}

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) float64 {
	// Day 1: Bellman-Ford to find max amount of each currency
  // HashMap: O(1) lookup
	amounts1 := make(map[string]float64)
	amounts1[initialCurrency] = 1.0

	// Run Bellman-Ford (or just process all pairs repeatedly)
  // Linear scan O(n)
	for i := 0; i < len(pairs1); i++ {
		updated := false
		for j, p := range pairs1 {
			from, to := p[0], p[1]
			r := rates1[j]
			if v, ok := amounts1[from]; ok {
				if v*r > amounts1[to] {
					amounts1[to] = v * r
					updated = true
				}
			}
			if v, ok := amounts1[to]; ok {
				if v/r > amounts1[from] {
					amounts1[from] = v / r
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	// Day 2: start with day1 amounts, find max back to initial
  // HashMap: O(1) lookup
	amounts2 := make(map[string]float64)
	for k, v := range amounts1 {
		amounts2[k] = v
	}

  // Linear scan O(n)
	for i := 0; i < len(pairs2); i++ {
		updated := false
		for j, p := range pairs2 {
			from, to := p[0], p[1]
			r := rates2[j]
			if v, ok := amounts2[from]; ok {
				if v*r > amounts2[to] {
					amounts2[to] = v * r
					updated = true
				}
			}
			if v, ok := amounts2[to]; ok {
				if v/r > amounts2[from] {
					amounts2[from] = v / r
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	return amounts2[initialCurrency]
}
```
