# 0040 — Combination Sum Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func combinationSum2(candidates []int, target int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #40: Combination Sum II
// https://leetcode.com/problems/combination-sum-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(candidates)
	result := [][]int{}
	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
  // Alokasi slice integer
			comb := make([]int, len(path))
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(i+1, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)) // [[1 1 6] [1 2 5] [1 7] [2 6]]

	// Test case 2
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5)) // [[1 2 2] [5]]
}

// Time: O(2^n) | Space: O(n)
```
