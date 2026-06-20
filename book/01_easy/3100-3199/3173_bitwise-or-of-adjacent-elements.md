# 3173 — Bitwise Or Of Adjacent Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BitwiseOrOfAdjacentElements(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3173: Bitwise OR of Adjacent Elements
// https://leetcode.com/problems/bitwise-or-of-adjacent-elements/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	// LeetCode name: orArray
	fmt.Println(BitwiseOrOfAdjacentElements([]int{1, 2, 3, 4})) // [3, 3, 7]
	fmt.Println(BitwiseOrOfAdjacentElements([]int{5, 1, 6}))     // [5, 7]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: orArray
func BitwiseOrOfAdjacentElements(nums []int) []int {
  // Alokasi slice
	result := make([]int, len(nums)-1)
  // Linear scan O(n)
	for i := 0; i < len(nums)-1; i++ {
		result[i] = nums[i] | nums[i+1]
	}
	return result
}
```
