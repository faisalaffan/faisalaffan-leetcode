# 0039 — Combination Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func combinationSum(candidates []int, target int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Backtracking

**Kompleksitas Waktu:** O(n^(target/min))  
**Kompleksitas Ruang:** O(target/min)

> **Untuk fresh graduate:** Kuasai dulu teknik **Backtracking** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
