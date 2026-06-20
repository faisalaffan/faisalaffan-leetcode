# 0280 — Wiggle Sort

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func wiggleSort(nums []int) `

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #280: Wiggle Sort
// https://leetcode.com/problems/wiggle-sort/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(1)

import "fmt"

func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if i%2 == 1 {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
}

func main() {
	nums1 := []int{3, 5, 2, 1, 6, 4}
	wiggleSort(nums1)
	fmt.Println(nums1)

	nums2 := []int{1, 2, 3, 4}
	wiggleSort(nums2)
	fmt.Println(nums2)

	nums3 := []int{1}
	wiggleSort(nums3)
	fmt.Println(nums3)
}
```
