# 3948 — Lexicographically Maximum Mex Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan graf. Tugasmu menjelajahi atau menganalisis konektivitas graf.

**Cara berpikir:** Adjacency list `map[int][]int`. Gunakan BFS (queue) atau DFS (rekursif) dengan visited set untuk hindari siklus.

**Fungsi Solusi:** `func maximumMEXArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3948: Lexicographically Maximum MEX Array
// https://leetcode.com/problems/lexicographically-maximum-mex-array/
// Difficulty: Hard
//
// Given array nums, construct a lexicographically maximum array
// result where result[i] is the MEX of a subsequence of nums
// ending at position i (or the MEX after certain operations).
//
// Approach: Track frequency of each value. For each i, find the
// MEX by checking the smallest non-negative integer not in the
// current multiset.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maximumMEXArray([]int{0, 1, 2, 3}))
	// Example 2
	fmt.Println(maximumMEXArray([]int{0, 0, 1, 1}))
	// Edge: empty
	fmt.Println(maximumMEXArray([]int{}))
}

func maximumMEXArray(nums []int) []int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return []int{}
	}

  // HashMap: O(1) lookup
	freq := make(map[int]int)
  // Alokasi slice
	result := make([]int, n)
	mex := 0

	for i, v := range nums {
		freq[v]++

		// Update mex: find smallest non-negative not in freq
		for freq[mex] > 0 {
			mex++
		}
		result[i] = mex
	}

	return result
}
```
