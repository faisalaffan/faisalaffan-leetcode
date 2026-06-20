# 1599 — Maximum Profit Of Operating A Centennial Wheel

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1599: Maximum Profit of Operating a Centennial Wheel
// https://leetcode.com/problems/maximum-profit-of-operating-a-centennial-wheel/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinOperationsMaxProfit([]int{8, 3}, 5, 6))
	fmt.Println(MinOperationsMaxProfit([]int{10, 9, 6}, 6, 4))
	fmt.Println(MinOperationsMaxProfit([]int{3, 4, 0, 5, 1}, 1, 92))
}

func MinOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	// Time: O(N), Space: O(1)
	// Each rotation can board up to 4 customers, costs runningCost, earns boardingCost per customer
	maxProfit := -1
	maxRotation := -1
	waiting := 0
	profit := 0
	rotation := 0

  // Linear scan O(n)
	for i := 0; i < len(customers) || waiting > 0; i++ {
		if i < len(customers) {
			waiting += customers[i]
		}

		// Board up to 4 customers
		boarded := 4
		if waiting < 4 {
			boarded = waiting
		}
		waiting -= boarded

		profit += boarded*boardingCost - runningCost
		rotation++

		if profit > maxProfit {
			maxProfit = profit
			maxRotation = rotation
		}
	}

	return maxRotation
}
```
