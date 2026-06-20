# 3948 — Lexicographically Maximum Mex Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah graf — kumpulan node (simpul) yang terhubung oleh edge (sisi). Tugasmu adalah menjelajahi graf, mencari jalur terpendek, atau menganalisis konektivitas.

Ibarat peta jalan: kota adalah node, jalan adalah edge. Kamu perlu mencari rute terpendek dari kota A ke kota B. Graf direpresentasikan dengan adjacency list (`map[int][]int` atau `[][]int`).

**Konsep kunci:** node, edge, directed/undirected, weighted/unweighted, BFS (level-order), DFS (depth-first), cycle detection.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumMEXArray(nums []int) []int
```

> **💡 Hint:** Track frequency of each value. For each i, find the

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Edge case: input kosong — langsung return
	if n == 0 {
		return []int{}
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
  // Alokasi slice integer
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
