# 2219 — Maximum Sum Score Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumSumScore(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2219: Maximum Sum Score of Array
// https://leetcode.com/problems/maximum-sum-score-of-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumSumScore(nums []int) int64 {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	maxScore := int64(-1 << 63)
	var prefix int64 = 0
	for i := 0; i < n; i++ {
		prefix += int64(nums[i])
		suffix := total - prefix + int64(nums[i])
		score := prefix
		if suffix > score {
			score = suffix
		}
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	// Test case 1
	fmt.Println(maximumSumScore([]int{1, 2, 3, 4, 5}))
	// Expected: 15

	// Test case 2
	fmt.Println(maximumSumScore([]int{-5, 10, -3, 4}))
	// Expected: 11

	// Test case 3
	fmt.Println(maximumSumScore([]int{-1, -2, -3}))
	// Expected: -1
}
```
