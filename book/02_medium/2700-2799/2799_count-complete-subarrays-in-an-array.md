# 2799 — Count Complete Subarrays In An Array

## Deskripsi

**Soal:** [2799. Count Complete Subarrays In An Array](https://leetcode.com/problems/count-complete-subarrays-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func CountCompleteSubarraysInAnArray(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2799: Count Complete Subarrays in an Array
// https://leetcode.com/problems/count-complete-subarrays-in-an-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func CountCompleteSubarraysInAnArray(nums []int) int {
	// Count distinct elements
  // Membuat map untuk pencarian O(1): key → value
	distinct := make(map[int]bool)
	for _, n := range nums {
		distinct[n] = true
	}
	target := len(distinct)

	left := 0
	count := 0
  // Membuat map untuk pencarian O(1): key → value
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
