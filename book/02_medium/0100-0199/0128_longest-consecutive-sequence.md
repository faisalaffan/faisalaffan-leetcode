# 0128 — Longest Consecutive Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestConsecutive(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #128: Longest Consecutive Sequence
// https://leetcode.com/problems/longest-consecutive-sequence/
// Difficulty: Medium

import "fmt"

func longestConsecutive(nums []int) int {
  // HashMap: O(1) lookup
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0
	for num := range numSet {
		if !numSet[num-1] {
			curr := num
			length := 1
			for numSet[curr+1] {
				curr++
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}

	return maxLen
}

func main() {
	// Test case 1
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // 4

	// Test case 2
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9

	// Test case 3
	fmt.Println(longestConsecutive([]int{})) // 0
}

// Time: O(n) | Space: O(n)
```
