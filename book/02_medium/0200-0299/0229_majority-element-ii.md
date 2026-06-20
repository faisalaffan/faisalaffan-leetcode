# 0229 — Majority Element Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func majorityElement(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #229: Majority Element II
// https://leetcode.com/problems/majority-element-ii/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func majorityElement(nums []int) []int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return nil
	}

	candidate1, candidate2 := 0, 0
	count1, count2 := 0, 0

	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 == 0 {
			candidate1 = num
			count1 = 1
		} else if count2 == 0 {
			candidate2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}

	result := []int{}
	threshold := len(nums) / 3
	if count1 > threshold {
		result = append(result, candidate1)
	}
	if count2 > threshold {
		result = append(result, candidate2)
	}

	return result
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1}))
	fmt.Println(majorityElement([]int{1, 2}))
}
```
