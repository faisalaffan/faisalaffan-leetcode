# 3795 — Minimum Subarray Length With Distinct Sum At Least K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumSubarrayLengthWithDistinctSumAtLeastK(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3795: Minimum Subarray Length With Distinct Sum At Least K
// https://leetcode.com/problems/minimum-subarray-length-with-distinct-sum-at-least-k/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func minimumSubarrayLengthWithDistinctSumAtLeastK(nums []int, k int) int {
	n := len(nums)
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	left := 0
	distinctSum := 0
	ans := math.MaxInt32

	for right := 0; right < n; right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			distinctSum += nums[right]
		}

		for distinctSum >= k {
			length := right - left + 1
			if length < ans {
				ans = length
			}
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinctSum -= nums[left]
			}
			left++
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

func main() {
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{2, 2, 3, 1}, 4))
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{3, 2, 3, 4}, 5))
	fmt.Println(minimumSubarrayLengthWithDistinctSumAtLeastK([]int{5, 5, 4}, 5))
}
```
