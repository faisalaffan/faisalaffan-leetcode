# 2551 — Put Marbles In Bags

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func putMarbles(weights []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2551: Put Marbles in Bags
// https://leetcode.com/problems/put-marbles-in-bags/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// putMarbles computes the difference between max and min scores.
// Key insight: each cut between i and i+1 contributes weights[i]+weights[i+1]
// to the total score. We need exactly k-1 cuts. Sort adjacent sums,
// take k-1 largest minus k-1 smallest.
//
// Complexity: O(n log n) time, O(n) space
func putMarbles(weights []int, k int) int64 {
	n := len(weights)
	if k <= 1 || k >= n {
		return 0
	}

  // Alokasi slice
	sums := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		sums[i] = weights[i] + weights[i+1]
	}
  // Sort O(n log n)
	sort.Ints(sums)

	var minScore, maxScore int64
	for i := 0; i < k-1; i++ {
		minScore += int64(sums[i])
		maxScore += int64(sums[n-2-i])
	}
	return maxScore - minScore
}

func main() {
	// Example from LeetCode
	fmt.Println("Test 1: [1,3,5,1], k=2 ->", putMarbles([]int{1, 3, 5, 1}, 2)) // 4

	// Additional test cases
	fmt.Println("Test 2: [1,3,5,1], k=1 ->", putMarbles([]int{1, 3, 5, 1}, 1)) // 0
	fmt.Println("Test 3: [1,4,2,5,3], k=3 ->", putMarbles([]int{1, 4, 2, 5, 3}, 3))
	fmt.Println("Test 4: [1,2], k=2 ->", putMarbles([]int{1, 2}, 2)) // 0
}
```
