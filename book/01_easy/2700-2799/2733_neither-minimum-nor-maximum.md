# 2733 — Neither Minimum Nor Maximum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func NeitherMinimumNorMaximum(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2733: Neither Minimum nor Maximum
// https://leetcode.com/problems/neither-minimum-nor-maximum/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(NeitherMinimumNorMaximum([]int{3, 2, 1, 4}))
	fmt.Println(NeitherMinimumNorMaximum([]int{1, 2}))
}

func NeitherMinimumNorMaximum(nums []int) int {
	if len(nums) < 3 {
		return -1
	}

	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	for _, v := range nums {
		if v != minVal && v != maxVal {
			return v
		}
	}

	return -1
}
```
