# 1822 — Sign Of The Product Of An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ArraySign(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1822: Sign of the Product of an Array
// https://leetcode.com/problems/sign-of-the-product-of-an-array/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func ArraySign(nums []int) int {
	negCount := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			negCount++
		}
	}
	if negCount%2 == 0 {
		return 1
	}
	return -1
}

func main() {
	fmt.Println(ArraySign([]int{-1, -2, -3, -4, 3, 2, 1}))
	fmt.Println(ArraySign([]int{1, 5, 0, 2, -3}))
	fmt.Println(ArraySign([]int{-1, 1, -1, 1, -1}))
}
```
