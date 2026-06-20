# 3641 — Longest Semi Repeating Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestSemiRepeatingSubarray(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3641: Longest Semi-Repeating Subarray
// https://leetcode.com/problems/longest-semi-repeating-subarray/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func longestSemiRepeatingSubarray(nums []int, k int) int {
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	dupCount := 0
	maxLen := 0
	l := 0

	for r, x := range nums {
		freq[x]++
		if freq[x] == 2 {
			dupCount++
		}

		for dupCount > k {
			left := nums[l]
			freq[left]--
			if freq[left] == 1 {
				dupCount--
			}
			l++
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return maxLen
}

func main() {
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 2, 3, 1, 2, 3, 4}, 2))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 4))
	fmt.Println(longestSemiRepeatingSubarray([]int{1, 1, 1, 1, 1}, 0))
}
```
