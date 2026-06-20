# 2248 — Intersection Of Multiple Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func IntersectionOfMultipleArrays(nums [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n * m), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2248: Intersection of Multiple Arrays
// https://leetcode.com/problems/intersection-of-multiple-arrays/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(IntersectionOfMultipleArrays([][]int{{3, 1, 2, 4, 5}, {1, 2, 3, 4}, {3, 4, 5, 6}})) // [3 4]
	fmt.Println(IntersectionOfMultipleArrays([][]int{{1, 2, 3}, {4, 5, 6}}))                        // []
}

// Time: O(n * m), Space: O(n)
func IntersectionOfMultipleArrays(nums [][]int) []int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return []int{}
	}

  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, v := range nums[0] {
		freq[v] = 1
	}

	for i := 1; i < len(nums); i++ {
  // HashMap: O(1) lookup
		seen := make(map[int]bool)
		for _, v := range nums[i] {
			if !seen[v] {
				freq[v]++
				seen[v] = true
			}
		}
	}

	var result []int
	for v, c := range freq {
		if c == len(nums) {
			result = append(result, v)
		}
	}
  // Sort O(n log n)
	sort.Ints(result)
	return result
}
```
