# 0992 — Subarrays With K Different Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func subarraysWithKDistinct(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Sliding Window

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #992: Subarrays with K Different Integers
// https://leetcode.com/problems/subarrays-with-k-different-integers/
// Difficulty: Hard
//
// Approach: atMostK trick.
//   subarraysWithKDistinct(nums, k) = atMostK(nums, k) - atMostK(nums, k-1)
//   atMostK counts subarrays with <= K distinct integers using a sliding window.

import "fmt"

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2)) // 7
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3)) // 3
}

func subarraysWithKDistinct(nums []int, k int) int {
	return atMostK(nums, k) - atMostK(nums, k-1)
}

func atMostK(nums []int, k int) int {
	if k == 0 {
		return 0
	}
  // HashMap: O(1) lookup
	count := make(map[int]int)
	left, result := 0, 0
	for right := 0; right < len(nums); right++ {
		count[nums[right]]++
		for len(count) > k {
			count[nums[left]]--
			if count[nums[left]] == 0 {
				delete(count, nums[left])
			}
			left++
		}
		result += right - left + 1
	}
	return result
}
```
