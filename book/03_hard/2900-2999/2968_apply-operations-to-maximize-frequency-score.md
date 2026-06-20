# 2968 — Apply Operations To Maximize Frequency Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxFrequencyScore(nums []int, k int64) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window, Binary Search, Sorting, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2968: Apply Operations to Maximize Frequency Score
// https://leetcode.com/problems/apply-operations-to-maximize-frequency-score/
// Difficulty: Hard
//
// Sort the array. For each window [l, r], compute cost to make all elements
// equal to the median (nums[mid]) within k operations. Use sliding window
// with prefix sums to check feasibility efficiently.
//
// Cost to make all elements in [l, r] equal to nums[mid]:
//   leftCost  = nums[mid]*(mid-l) - sum(l..mid-1)
//   rightCost = sum(mid+1..r) - nums[mid]*(r-mid)

import (
	"fmt"
	"sort"
)

func maxFrequencyScore(nums []int, k int64) int {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
  // Alokasi slice
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	ans := 1
	left := 0
	for right := 0; right < n; right++ {
		// Shrink window from left if cost > k
  // Two-pointer loop
		for left < right {
			mid := (left + right) / 2
			leftCost := int64(nums[mid])*int64(mid-left) - (prefix[mid] - prefix[left])
			rightCost := (prefix[right+1] - prefix[mid+1]) - int64(nums[mid])*int64(right-mid)
			if leftCost+rightCost <= k {
				break
			}
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}

func main() {
	// Example
	fmt.Println(maxFrequencyScore([]int{1, 2, 3, 4, 5, 6}, 10))
	fmt.Println(maxFrequencyScore([]int{1, 4, 4, 2, 7}, 6))

	// Edge cases
	fmt.Println(maxFrequencyScore([]int{1, 1, 1}, 0))
	fmt.Println(maxFrequencyScore([]int{1, 100}, 50))
}
```
