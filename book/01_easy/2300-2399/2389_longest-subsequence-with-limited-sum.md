# 2389 — Longest Subsequence With Limited Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestSubsequenceWithLimitedSum(nums []int, queries []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
  // Alokasi slice integer
	prefix := make([]int, len(nums))
	sum := 0
	for i, n := range nums {
		sum += n
		prefix[i] = sum
	}

  // Alokasi slice integer
	res := make([]int, len(queries))
	for i, q := range queries {
		// Binary search for rightmost index where prefix <= q
		res[i] = sort.SearchInts(prefix, q+1)
	}
	return res
}
```
