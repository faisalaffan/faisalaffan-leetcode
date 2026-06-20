# 1979 — Find Greatest Common Divisor Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindGreatestCommonDivisorOfArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1979: Find Greatest Common Divisor of Array
// https://leetcode.com/problems/find-greatest-common-divisor-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{2, 5, 6, 9, 10})) // 2
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{7, 5, 6, 8, 3}))  // 1
	fmt.Println(FindGreatestCommonDivisorOfArray([]int{3, 3}))            // 3
}

// Time: O(n), Space: O(1)
func FindGreatestCommonDivisorOfArray(nums []int) int {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	a, b := min, max
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
