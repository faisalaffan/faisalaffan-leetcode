# 2670 — Find The Distinct Difference Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindTheDistinctDifferenceArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2670: Find the Distinct Difference Array
// https://leetcode.com/problems/find-the-distinct-difference-array/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(FindTheDistinctDifferenceArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(FindTheDistinctDifferenceArray([]int{3, 2, 3, 4, 2}))
}

func FindTheDistinctDifferenceArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice
	suffixDistinct := make([]int, n+1)
	seen := map[int]bool{}

	for i := n - 1; i >= 0; i-- {
		suffixDistinct[i] = suffixDistinct[i+1]
		if !seen[nums[i]] {
			seen[nums[i]] = true
			suffixDistinct[i]++
		}
	}

	seen = map[int]bool{}
	prefixDistinct := 0
  // Alokasi slice
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if !seen[nums[i]] {
			seen[nums[i]] = true
			prefixDistinct++
		}
		ans[i] = prefixDistinct - suffixDistinct[i+1]
	}

	return ans
}
```
