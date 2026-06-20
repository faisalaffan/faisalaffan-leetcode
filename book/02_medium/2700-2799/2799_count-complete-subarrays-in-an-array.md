# 2799 — Count Complete Subarrays In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func CountCompleteSubarraysInAnArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2799: Count Complete Subarrays in an Array
// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountCompleteSubarraysInAnArray(nums []int) int {
	// Count distinct elements
  // HashMap: O(1) lookup
	distinct := make(map[int]bool)
	for _, n := range nums {
		distinct[n] = true
	}
	target := len(distinct)

	left := 0
	count := 0
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	unique := 0

	for right := 0; right < len(nums); right++ {
		freq[nums[right]]++
		if freq[nums[right]] == 1 {
			unique++
		}

		for unique == target {
			// All subarrays from left to right-end are valid
			count += len(nums) - right
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				unique--
			}
			left++
		}
	}

	return count
}

func main() {
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 3, 1, 2, 2}))
	fmt.Println(CountCompleteSubarraysInAnArray([]int{1, 1}))
}
```
