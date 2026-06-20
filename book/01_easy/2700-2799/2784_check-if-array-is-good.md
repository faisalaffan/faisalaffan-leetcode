# 2784 — Check If Array Is Good

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CheckIfArrayIsGood(nums []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2784: Check if Array is Good
// https://leetcode.com/problems/check-if-array-is-good/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(CheckIfArrayIsGood([]int{2, 1, 3}))
	fmt.Println(CheckIfArrayIsGood([]int{3, 4, 4, 1, 2, 1}))
}

func CheckIfArrayIsGood(nums []int) bool {
	n := len(nums) - 1
  // Alokasi slice
	counts := make([]int, n+1)
	for _, v := range nums {
		if v > n {
			return false
		}
		counts[v]++
	}
	if counts[n] != 2 {
		return false
	}
	for i := 1; i < n; i++ {
		if counts[i] != 1 {
			return false
		}
	}
	return true
}
```
