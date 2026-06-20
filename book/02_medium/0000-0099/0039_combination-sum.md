# 0039 — Combination Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func combinationSum(candidates []int, target int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Backtracking, Sorting

**Waktu:** O(n^(target/min))  |  **Ruang:** O(target/min)

> 🎓 **Fresh Grad Tips:** Kuasai **Backtracking** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #39: Combination Sum
// https://leetcode.com/problems/combination-sum/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
  // Sort O(n log n)
	sort.Ints(candidates)
	result := [][]int{}
	var backtrack func(start int, target int, path []int)
	backtrack = func(start int, target int, path []int) {
		if target == 0 {
  // Alokasi slice
			comb := make([]int, len(path))
			copy(comb, path)
			result = append(result, comb)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			backtrack(i, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, target, []int{})
	return result
}

func main() {
	// Test case 1
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2 2 3] [7]]

	// Test case 2
	fmt.Println(combinationSum([]int{2, 3, 5}, 8)) // [[2 2 2 2] [2 3 3] [3 5]]

	// Test case 3
	fmt.Println(combinationSum([]int{2}, 1)) // []
}

// Time: O(n^(target/min)) | Space: O(target/min)
```
