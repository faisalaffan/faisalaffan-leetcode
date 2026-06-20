# 1464 — Maximum Product Of Two Elements In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxProduct(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1464: Maximum Product of Two Elements in an Array
// https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/
// Difficulty: Easy
//
// LeetCode submission: func maxProduct(nums []int) int

import "fmt"

func main() {
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 4, 5, 2})) // 12
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{1, 5, 4, 5})) // 16
	fmt.Println(MaximumProductOfTwoElementsInAnArray([]int{3, 7}))       // 12
}

// Time: O(n), Space: O(1)
func MaximumProductOfTwoElementsInAnArray(nums []int) int {
	first, second := 0, 0
	for _, v := range nums {
		if v > first {
			second = first
			first = v
		} else if v > second {
			second = v
		}
	}
	return (first - 1) * (second - 1)
}
```
