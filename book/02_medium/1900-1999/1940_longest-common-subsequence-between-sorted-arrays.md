# 1940 — Longest Common Subsequence Between Sorted Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func LongestCommonSubsequenceBetweenSortedArrays(arrs [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(total elements), Space: O(unique elements)  |  **Ruang:** O(unique elements)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1940: Longest Common Subsequence Between Sorted Arrays
// https://leetcode.com/problems/longest-common-subsequence-between-sorted-arrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{1, 3, 4}, {1, 4, 7, 9}}))
	fmt.Println(LongestCommonSubsequenceBetweenSortedArrays([][]int{{2, 3, 6, 8}, {1, 2, 3, 5, 6, 7, 10}, {2, 3, 4, 6, 9}}))
}

// Time: O(total elements), Space: O(unique elements)
func LongestCommonSubsequenceBetweenSortedArrays(arrs [][]int) []int {
	// Since arrays are sorted, use frequency counting
	// Numbers appearing in ALL arrays are the answer
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, arr := range arrs {
		for _, v := range arr {
			freq[v]++
		}
	}

	n := len(arrs)
  // Alokasi slice
	result := make([]int, 0)
	for _, v := range arrs[0] {
		if freq[v] == n {
			result = append(result, v)
		}
	}
	return result
}
```
