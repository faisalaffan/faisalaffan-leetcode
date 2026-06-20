# 0213 — House Robber Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func rob(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #213: House Robber II
// https://leetcode.com/problems/house-robber-ii/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func rob(nums []int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	robLinear := func(arr []int) int {
		prev, curr := 0, 0
		for _, num := range arr {
			prev, curr = curr, max(curr, prev+num)
		}
		return curr
	}

	return max(robLinear(nums[1:]), robLinear(nums[:len(nums)-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{1, 2, 3}))
}
```
