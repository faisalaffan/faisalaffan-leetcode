# 1636 — Sort Array By Increasing Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FrequencySort(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
  // Custom sort
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
