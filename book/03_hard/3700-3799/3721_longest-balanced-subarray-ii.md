# 3721 — Longest Balanced Subarray Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestBalanced(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sliding Window

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3721: Longest Balanced Subarray II
// https://leetcode.com/problems/longest-balanced-subarray-ii/
// Difficulty: Hard
//
// Find longest subarray where number of distinct even numbers equals
// number of distinct odd numbers.
//
// Approach: Sliding window with frequency maps. For each window,
// track distinct even and odd counts using hash maps.

import "fmt"

func main() {
	// Example 1
	fmt.Println(longestBalanced([]int{2, 5, 4, 3}))
	// Example 2
	fmt.Println(longestBalanced([]int{1, 3, 5, 2, 4}))
	// Edge: single element
	fmt.Println(longestBalanced([]int{1}))
	// Edge: all even
	fmt.Println(longestBalanced([]int{2, 4, 6}))
}

func longestBalanced(nums []int) int {
	n := len(nums)
	result := 0

	// For each starting position, expand window
	for i := 0; i < n; i++ {
  // HashMap: O(1) lookup
		evenSet := make(map[int]bool)
  // HashMap: O(1) lookup
		oddSet := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				if !evenSet[nums[j]] {
					evenSet[nums[j]] = true
					evenCount++
				}
			} else {
				if !oddSet[nums[j]] {
					oddSet[nums[j]] = true
					oddCount++
				}
			}
			if evenCount == oddCount {
				length := j - i + 1
				if length > result {
					result = length
				}
			}
		}
	}

	return result
}
```
