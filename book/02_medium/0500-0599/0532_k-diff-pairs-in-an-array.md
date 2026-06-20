# 0532 — K Diff Pairs In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindPairs(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1) (ignoring sorting overhead)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #532: K-diff Pairs in an Array
// https://leetcode.com/problems/k-diff-pairs-in-an-array/
// Difficulty: Medium
// Time: O(n log n)
// Space: O(1) (ignoring sorting overhead)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(FindPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(FindPairs([]int{1, 3, 1, 5, 4}, 0))
}

func FindPairs(nums []int, k int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
	count := 0
	i, j := 0, 1

	for i < n && j < n {
		if i == j || nums[j]-nums[i] < k {
			j++
		} else if nums[j]-nums[i] > k {
			i++
		} else {
			count++
			i++
			j++
			// Skip duplicates
			for j < n && nums[j] == nums[j-1] {
				j++
			}
		}
	}

	return count
}
```
