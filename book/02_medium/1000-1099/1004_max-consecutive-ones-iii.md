# 1004 — Max Consecutive Ones Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func longestOnes(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1004: Max Consecutive Ones III
// https://leetcode.com/problems/max-consecutive-ones-iii/
// Difficulty: Medium
//
// Approach: Sliding window (expand right, shrink left when zeros > k)
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)) // 6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) // 10
	fmt.Println(longestOnes([]int{0, 0, 0, 0}, 0)) // 0
}

func longestOnes(nums []int, k int) int {
	left := 0
	zeros := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		for zeros > k {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		if right-left+1 > result {
			result = right - left + 1
		}
	}

	return result
}
```
