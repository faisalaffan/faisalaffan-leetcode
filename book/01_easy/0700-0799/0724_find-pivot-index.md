# 0724 — Find Pivot Index

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func pivotIndex(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #724: Find Pivot Index
// https://leetcode.com/problems/find-pivot-index/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))   // 3
	fmt.Println(pivotIndex([]int{1, 2, 3}))             // -1
	fmt.Println(pivotIndex([]int{2, 1, -1}))            // 0
}

// pivotIndex finds the index where sum of left elements equals sum of right elements.
// Time: O(n). Space: O(1).
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	leftSum := 0
	for i, v := range nums {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
```
