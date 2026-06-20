# 2389 — Longest Subsequence With Limited Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LongestSubsequenceWithLimitedSum(nums []int, queries []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2389: Longest Subsequence With Limited Sum
// https://leetcode.com/problems/longest-subsequence-with-limited-sum/
// Difficulty: Easy
// Time O((n+m) log n) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(LongestSubsequenceWithLimitedSum([]int{4, 5, 2, 1}, []int{3, 10, 21})) // [2,3,4]
	fmt.Println(LongestSubsequenceWithLimitedSum([]int{2, 3, 4, 5}, []int{1}))          // [0]
}

func LongestSubsequenceWithLimitedSum(nums []int, queries []int) []int {
  // Sort O(n log n)
	sort.Ints(nums)
  // Alokasi slice
	prefix := make([]int, len(nums))
	sum := 0
	for i, n := range nums {
		sum += n
		prefix[i] = sum
	}

  // Alokasi slice
	res := make([]int, len(queries))
	for i, q := range queries {
		// Binary search for rightmost index where prefix <= q
		res[i] = sort.SearchInts(prefix, q+1)
	}
	return res
}
```
