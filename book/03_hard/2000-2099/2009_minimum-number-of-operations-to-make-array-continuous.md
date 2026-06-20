# 2009 — Minimum Number Of Operations To Make Array Continuous

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func minOperations(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sliding Window, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2009: Minimum Number of Operations to Make Array Continuous
// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/
// Difficulty: Hard
// Approach: Sort + sliding window on unique elements.
// A continuous array has max-min < n and distinct elements.
// For each element as left bound, find rightmost element with value < left+n.
// Answer = n - max window size.

import (
	"fmt"
	"sort"
)

func minOperations(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// Sort and deduplicate
  // Sort O(n log n)
	sort.Ints(nums)
  // Alokasi slice
	uniq := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		uniq = append(uniq, nums[i])
	}

	maxWindow := 0
	right := 0
	for left := 0; left < len(uniq); left++ {
		for right < len(uniq) && uniq[right] < uniq[left]+n {
			right++
		}
		windowSize := right - left
		if windowSize > maxWindow {
			maxWindow = windowSize
		}
	}

	return n - maxWindow
}

func main() {
	// Example: [4,2,5,3] -> 0 (already continuous)
	fmt.Println(minOperations([]int{4, 2, 5, 3}))

	// Additional tests
	fmt.Println(minOperations([]int{1, 2, 3, 5, 6})) // [1,2,3,5,6] -> need to make [2,3,4,5,6] or [1,2,3,4,5], ops = 1
	fmt.Println(minOperations([]int{1, 10, 100, 1000}))
	fmt.Println(minOperations([]int{8, 5, 9, 9, 5, 7, 2, 3}))
}
```
