# 3740 — Minimum Distance Between Three Equal Elements I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumDistanceBetweenThreeEqualElementsI(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3740: Minimum Distance Between Three Equal Elements I
// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-i/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1, 2, 1, 1, 3}))
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1, 1, 2, 3, 2, 1, 2}))
	fmt.Println(MinimumDistanceBetweenThreeEqualElementsI([]int{1}))
}

// Time: O(n)
// Space: O(n)
func MinimumDistanceBetweenThreeEqualElementsI(nums []int) int {
  // HashMap: O(1) lookup
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	ans := math.MaxInt
	for _, indices := range pos {
		if len(indices) < 3 {
			continue
		}
		for i := 0; i <= len(indices)-3; i++ {
			dist := 2 * (indices[i+2] - indices[i])
			if dist < ans {
				ans = dist
			}
		}
	}

	if ans == math.MaxInt {
		return -1
	}
	return ans
}
```
