# 3647 — Maximum Weight In Two Bags

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumWeightInTwoBags(weight []int, bag1 int, bag2 int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3647: Maximum Weight in Two Bags
// https://leetcode.com/problems/maximum-weight-in-two-bags/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumWeightInTwoBags(weight []int, bag1 int, bag2 int) int {
  // Custom sort
	sort.Slice(weight, func(i, j int) bool {
		return weight[i] > weight[j]
	})

	// Try to find best combination: one item for bag1, one for bag2
	// Since items are sorted descending, we try all pairs
	n := len(weight)
	best := 0

	// Single item approach: pick up to bag1 items, then fill bag2 with remaining
	maxBag1 := 0
	used := -1
	for i := 0; i < n && weight[i] <= bag1; i++ {
		if weight[i] > maxBag1 {
			maxBag1 = weight[i]
			used = i
		}
	}

	if used != -1 {
		for i := 0; i < n; i++ {
			if i != used && weight[i] <= bag2 {
				if maxBag1+weight[i] > best {
					best = maxBag1 + weight[i]
				}
				break
			}
		}
	}

	maxBag2 := 0
	used2 := -1
	for i := 0; i < n && weight[i] <= bag2; i++ {
		if weight[i] > maxBag2 {
			maxBag2 = weight[i]
			used2 = i
		}
	}

	if used2 != -1 {
		for i := 0; i < n; i++ {
			if i != used2 && weight[i] <= bag1 {
				if maxBag2+weight[i] > best {
					best = maxBag2 + weight[i]
				}
				break
			}
		}
	}

	return best
}

func main() {
	fmt.Println(maximumWeightInTwoBags([]int{10, 20, 30, 40}, 30, 40))
	fmt.Println(maximumWeightInTwoBags([]int{5, 5, 10}, 10, 5))
	fmt.Println(maximumWeightInTwoBags([]int{2, 3, 5, 7}, 5, 5))
}
```
