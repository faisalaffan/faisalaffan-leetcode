# 1217 — Minimum Cost To Move Chips To The Same Position

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minCostToMoveChips(position []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1217: Minimum Cost to Move Chips to The Same Position
// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))    // 1
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3})) // 2
}

// LeetCode submission: minCostToMoveChips
func minCostToMoveChips(position []int) int {
	even, odd := 0, 0
	for _, p := range position {
		if p%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even < odd {
		return even
	}
	return odd
}
```
