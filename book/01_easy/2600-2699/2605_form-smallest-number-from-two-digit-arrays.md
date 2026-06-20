# 2605 — Form Smallest Number From Two Digit Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FormSmallestNumberFromTwoDigitArrays(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2605: Form Smallest Number From Two Digit Arrays
// https://leetcode.com/problems/form-smallest-number-from-two-digit-arrays/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(FormSmallestNumberFromTwoDigitArrays([]int{4, 1, 3}, []int{5, 7}))       // 15
	fmt.Println(FormSmallestNumberFromTwoDigitArrays([]int{3, 5, 2, 6}, []int{3, 1, 7})) // 3
}

func FormSmallestNumberFromTwoDigitArrays(nums1 []int, nums2 []int) int {
	seen := [10]bool{}
	for _, n := range nums1 {
		seen[n] = true
	}

	common := 10
	for _, n := range nums2 {
		if seen[n] && n < common {
			common = n
		}
	}
	if common < 10 {
		return common
	}

	min1, min2 := 10, 10
	for _, n := range nums1 {
		if n < min1 {
			min1 = n
		}
	}
	for _, n := range nums2 {
		if n < min2 {
			min2 = n
		}
	}
	if min1 < min2 {
		return min1*10 + min2
	}
	return min2*10 + min1
}
```
