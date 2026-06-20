# 0491 — Non Decreasing Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func NonDecreasingSubsequences(nums []int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Backtracking

**Waktu:** O(2^n * n) worst case  |  **Ruang:** O(2^n * n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #491: Non-decreasing Subsequences
// https://leetcode.com/problems/non-decreasing-subsequences/
// Difficulty: Medium
// Time: O(2^n * n) worst case
// Space: O(2^n * n)

import "fmt"

func main() {
	fmt.Println(NonDecreasingSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(NonDecreasingSubsequences([]int{4, 4, 3, 2, 1}))
}

func NonDecreasingSubsequences(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) >= 2 {
  // Alokasi slice
			cp := make([]int, len(path))
			copy(cp, path)
			result = append(result, cp)
		}
  // HashMap: O(1) lookup
		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				backtrack(i+1, append(path, nums[i]))
			}
		}
	}
	backtrack(0, []int{})
	return result
}
```
