# 0384 — Shuffle An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func Constructor(nums []int) Solution`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n) per shuffle  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #384: Shuffle an Array
// https://leetcode.com/problems/shuffle-an-array/
// Difficulty: Medium
// Time: O(n) per shuffle | Space: O(n)

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int
}

func Constructor(nums []int) Solution {
  // Alokasi slice
	orig := make([]int, len(nums))
	copy(orig, nums)
	return Solution{original: orig}
}

func (s *Solution) Reset() []int {
  // Alokasi slice
	result := make([]int, len(s.original))
	copy(result, s.original)
	return result
}

func (s *Solution) Shuffle() []int {
  // Alokasi slice
	result := make([]int, len(s.original))
	copy(result, s.original)
	// Fisher-Yates
	for i := len(result) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func main() {
	sol := Constructor([]int{1, 2, 3})
	fmt.Println("Reset:", sol.Reset())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Shuffle:", sol.Shuffle())
	fmt.Println("Reset:", sol.Reset())
}
```
