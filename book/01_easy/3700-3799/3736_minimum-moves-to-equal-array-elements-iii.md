# 3736 — Minimum Moves To Equal Array Elements Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumMovesToEqualArrayElementsIii(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3736: Minimum Moves to Equal Array Elements III
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-iii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumMovesToEqualArrayElementsIii([]int{2, 1, 3}))
	fmt.Println(MinimumMovesToEqualArrayElementsIii([]int{4, 4, 5}))
}

// Time: O(n)
// Space: O(1)
func MinimumMovesToEqualArrayElementsIii(nums []int) int {
	maxVal := nums[0]
	sum := 0
	for _, v := range nums {
		sum += v
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal*len(nums) - sum
}
```
