# 2303 — Calculate Amount Paid In Taxes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func CalculateAmountPaidInTaxes(brackets [][]int, income int) float64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2303: Calculate Amount Paid in Taxes
// https://leetcode.com/problems/calculate-amount-paid-in-taxes/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(CalculateAmountPaidInTaxes([][]int{{3, 50}, {7, 10}, {12, 25}}, 10)) // 2.65
	fmt.Println(CalculateAmountPaidInTaxes([][]int{{1, 0}, {4, 25}, {5, 50}}, 2))    // 0.25
}

func CalculateAmountPaidInTaxes(brackets [][]int, income int) float64 {
	tax := 0.0
	prev := 0
	for _, b := range brackets {
		upper := b[0]
		percent := float64(b[1]) / 100.0
		taxable := min(upper, income) - prev
		if taxable > 0 {
			tax += float64(taxable) * percent
		}
		prev = upper
		if income <= upper {
			break
		}
	}
	return tax
}
```
