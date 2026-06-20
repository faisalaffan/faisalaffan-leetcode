# 2563 — Count The Number Of Fair Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countFairPairs(nums []int, lower int, upper int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2563: Count the Number of Fair Pairs
// https://leetcode.com/problems/count-the-number-of-fair-pairs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
  // Sort O(n log n)
	sort.Ints(nums)
	var ans int64
	n := len(nums)

	for i := 0; i < n; i++ {
		// Find lower bound for j > i such that nums[j] >= lower - nums[i]
		left := sort.Search(n, func(j int) bool {
			return j > i && nums[j] >= lower-nums[i]
		})
		// Find upper bound for j > i such that nums[j] <= upper - nums[i]
		right := sort.Search(n, func(j int) bool {
			return j > i && nums[j] > upper-nums[i]
		})
		if right > left {
			ans += int64(right - left)
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", countFairPairs([]int{0, 0, 0, 0, 0, 0}, 0, 0))
	// Expected: 15
}
```
