# 3741 — Minimum Distance Between Three Equal Elements Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumDistanceBetweenThreeEqualElementsIi(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3741: Minimum Distance Between Three Equal Elements II
// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumDistanceBetweenThreeEqualElementsIi(nums []int) int {
	ans := -1
  // HashMap: O(1) lookup
	prev1 := make(map[int]int)
  // HashMap: O(1) lookup
	prev2 := make(map[int]int)

	// prev1[v] = last index where v appeared
	// prev2[v] = second-last index where v appeared

	for i, v := range nums {
		if p2, ok := prev2[v]; ok {
			dist := 2 * (i - p2)
			if ans == -1 || dist < ans {
				ans = dist
			}
		}
		// Shift: prev2 gets prev1, prev1 gets current
		prev2[v] = prev1[v]
		prev1[v] = i
	}

	return ans
}

func main() {
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 3, 1, 1, 2, 1}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 2, 3, 4}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 1, 1}))
}
```
