# 0922 — Sort Array By Parity Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sortArrayByParityII(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #922: Sort Array By Parity II
// https://leetcode.com/problems/sort-array-by-parity-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(sortArrayByParityII([]int{4, 2, 5, 7})) // [4,5,2,7] or [4,7,2,5]
	fmt.Println(sortArrayByParityII([]int{2, 3}))        // [2,3]
}

// sortArrayByParityII puts even numbers at even indices, odd numbers at odd indices.
// Time: O(n). Space: O(1).
func sortArrayByParityII(nums []int) []int {
	j := 1 // odd pointer
  // Linear scan O(n)
	for i := 0; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
```
