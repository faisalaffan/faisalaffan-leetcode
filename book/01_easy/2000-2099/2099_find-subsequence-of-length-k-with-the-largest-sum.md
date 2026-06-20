# 2099 — Find Subsequence Of Length K With The Largest Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindSubsequenceOfLengthKWithTheLargestSum(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2099: Find Subsequence of Length K With the Largest Sum
// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{2, 1, 3, 3}, 2))       // [3 3]
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{-1, -2, 3, 4}, 3))    // [-1 3 4]
	fmt.Println(FindSubsequenceOfLengthKWithTheLargestSum([]int{3, 4, 3, 3}, 2))      // [3 4]
}

// Time: O(n log n), Space: O(n)
func FindSubsequenceOfLengthKWithTheLargestSum(nums []int, k int) []int {
	type pair struct {
		val int
		idx int
	}

	pairs := make([]pair, len(nums))
	for i, v := range nums {
		pairs[i] = pair{v, i}
	}

	// Sort by value descending
  // Custom sort dengan comparator
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val > pairs[j].val
	})

	// Take top k
	selected := pairs[:k]

	// Sort by original index to preserve order
  // Custom sort dengan comparator
	sort.Slice(selected, func(i, j int) bool {
		return selected[i].idx < selected[j].idx
	})

  // Alokasi slice integer
	result := make([]int, k)
	for i, p := range selected {
		result[i] = p.val
	}
	return result
}
```
