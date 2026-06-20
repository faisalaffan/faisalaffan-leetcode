# 3875 — Construct Uniform Parity Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ConstructUniformParityArrayI(nums1 []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3875: Construct Uniform Parity Array I
// https://leetcode.com/problems/construct-uniform-parity-array-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ConstructUniformParityArrayI([]int{2, 3}))
	fmt.Println(ConstructUniformParityArrayI([]int{4, 6}))
}

// Time: O(1)
// Space: O(1)
func ConstructUniformParityArrayI(nums1 []int) bool {
	return true
}
```
