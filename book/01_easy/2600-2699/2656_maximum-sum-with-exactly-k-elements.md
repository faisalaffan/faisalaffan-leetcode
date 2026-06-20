# 2656 — Maximum Sum With Exactly K Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumSumWithExactlyKElements(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2656: Maximum Sum With Exactly K Elements
// https://leetcode.com/problems/maximum-sum-with-exactly-k-elements/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(MaximumSumWithExactlyKElements([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(MaximumSumWithExactlyKElements([]int{5, 5, 5}, 2))
}

func MaximumSumWithExactlyKElements(nums []int, k int) int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	// Sum of arithmetic series: maxVal + (maxVal+1) + ... + (maxVal+k-1)
	return k * (2*maxVal + k - 1) / 2
}
```
