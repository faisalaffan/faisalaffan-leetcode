# 1929 — Concatenation Of Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConcatenationOfArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1929: Concatenation of Array
// https://leetcode.com/problems/concatenation-of-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConcatenationOfArray([]int{1, 2, 1}))       // [1,2,1,1,2,1]
	fmt.Println(ConcatenationOfArray([]int{1, 3, 2, 1}))    // [1,3,2,1,1,3,2,1]
}

// Time: O(n), Space: O(n)
func ConcatenationOfArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = v
	}
	return ans
}
```
