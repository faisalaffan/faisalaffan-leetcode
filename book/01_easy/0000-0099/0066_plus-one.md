# 0066 — Plus One

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func PlusOne(digits []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1) (excluding output)


## 💻 Solusi Go

```go
package main

// LeetCode #66: Plus One
// https://leetcode.com/problems/plus-one/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1) (excluding output)
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(PlusOne([]int{1, 2, 3}))
	fmt.Println(PlusOne([]int{4, 3, 2, 1}))
	fmt.Println(PlusOne([]int{9}))
}
```
