# 3878 — Count Good Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countGoodSubarrays(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Sliding Window

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3878: Count Good Subarrays
// https://leetcode.com/problems/count-good-subarrays/
// Difficulty: Hard
//
// Count subarrays where all elements are distinct (no duplicates).
//
// Approach: Sliding window with frequency map. For each right index,
// maintain window with all distinct elements. Count subarrays ending
// at right with all distinct elements.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countGoodSubarrays([]int{1, 2, 3}))
	// Example 2
	fmt.Println(countGoodSubarrays([]int{1, 2, 1, 3}))
	// Edge: all same
	fmt.Println(countGoodSubarrays([]int{1, 1, 1}))
	// Edge: empty
	fmt.Println(countGoodSubarrays([]int{}))
}

func countGoodSubarrays(nums []int) int64 {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

  // HashMap: O(1) lookup
	freq := make(map[int]int)
	var ans int64
	left := 0

	for right := 0; right < n; right++ {
		freq[nums[right]]++

		for freq[nums[right]] > 1 {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				delete(freq, nums[left])
			}
			left++
		}

		ans += int64(right - left + 1)
	}

	return ans
}
```
