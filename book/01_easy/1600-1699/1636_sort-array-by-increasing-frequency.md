# 1636 — Sort Array By Increasing Frequency

## Deskripsi

**Soal:** [1636. Sort Array By Increasing Frequency](https://leetcode.com/problems/sort-array-by-increasing-frequency/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func FrequencySort(nums []int) []int`

## Solusi Go

```go
package main

// LeetCode #1636: Sort Array by Increasing Frequency
// https://leetcode.com/problems/sort-array-by-increasing-frequency/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(n)
func FrequencySort(nums []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}

func main() {
	fmt.Println(FrequencySort([]int{1, 1, 2, 2, 2, 3}))
	fmt.Println(FrequencySort([]int{2, 3, 1, 3, 2}))
	fmt.Println(FrequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
```
