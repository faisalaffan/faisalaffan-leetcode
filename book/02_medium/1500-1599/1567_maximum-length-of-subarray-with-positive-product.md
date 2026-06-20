# 1567 — Maximum Length Of Subarray With Positive Product

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func GetMaxLen(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1567: Maximum Length of Subarray With Positive Product
// https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(GetMaxLen([]int{1, -2, -3, 4}))
	fmt.Println(GetMaxLen([]int{0, 1, -2, -3, -4}))
	fmt.Println(GetMaxLen([]int{-1, -2, -3, 0, 1}))
}

func GetMaxLen(nums []int) int {
	// Time: O(N), Space: O(1)
	// Track first occurrence of positive and negative prefix products
	maxLen := 0
	firstPos := -1
	firstNeg := -1
	prefix := 1 // 1 = positive, -1 = negative

	for i, num := range nums {
		if num > 0 {
			prefix = prefix // sign unchanged
		} else if num < 0 {
			prefix = -prefix
		} else {
			// Reset at zero
			prefix = 1
			firstPos = -1
			firstNeg = -1
			continue
		}

		if prefix == 1 {
			if firstPos == -1 {
				firstPos = i
			}
			if i-firstPos+1 > maxLen {
				maxLen = i - firstPos + 1
			}
		} else { // prefix == -1
			if firstNeg == -1 {
				firstNeg = i
			}
			if i-firstNeg+1 > maxLen {
				maxLen = i - firstNeg + 1
			}
		}
	}

	return maxLen
}
```
