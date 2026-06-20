# 0384 — Shuffle An Array

## Deskripsi

**Soal:** [0384. Shuffle An Array](https://leetcode.com/problems/shuffle-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) per shuffle  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func Constructor(nums []int) Solution`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	orig := make([]int, len(nums))
	copy(orig, nums)
	return Solution{original: orig}
}

func (s *Solution) Reset() []int {
  // Membuat slice untuk menyimpan hasil
	result := make([]int, len(s.original))
	copy(result, s.original)
	return result
}

func (s *Solution) Shuffle() []int {
  // Membuat slice untuk menyimpan hasil
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
