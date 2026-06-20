# 2540 — Minimum Common Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumCommonValue(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2540: Minimum Common Value
// https://leetcode.com/problems/minimum-common-value/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumCommonValue([]int{1, 2, 3}, []int{2, 4}))       // 2
	fmt.Println(MinimumCommonValue([]int{1, 2, 3, 6}, []int{2, 3, 4, 5})) // 2
}

func MinimumCommonValue(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			return nums1[i]
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}
```
