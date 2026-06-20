# 3207 — Maximum Points After Enemy Battles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumPoints(enemyEnergies []int, currentEnergy int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3207: Maximum Points After Enemy Battles
// https://leetcode.com/problems/maximum-points-after-enemy-battles/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
  // Sort O(n log n)
	sort.Ints(enemyEnergies)
	n := len(enemyEnergies)
	energy := int64(currentEnergy)
	var points int64

	// First defeat all enemies we can defeat, accumulating energy
	i := 0
	for i < n && int64(enemyEnergies[i]) <= energy {
		energy -= int64(enemyEnergies[i])
		points++
		i++
	}

	if points == 0 {
		return 0
	}

	// Now we can use the smallest enemy to get energy
	smallest := enemyEnergies[0]
	for energy >= int64(smallest) {
		cnt := energy / int64(smallest)
		points += cnt
		energy %= int64(smallest)
		energy += int64(smallest) // Keep 1 point by defeating smallest again
		points--
	}
	return points
}

func main() {
	fmt.Println(maximumPoints([]int{3, 5, 6}, 3))  // Expected: 1
	fmt.Println(maximumPoints([]int{1, 2, 4, 8}, 3)) // Expected: 6
}
```
