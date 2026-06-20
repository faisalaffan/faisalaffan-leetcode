# 3423 — Maximum Difference Between Adjacent Elements In A Circular Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumDifferenceBetweenAdjacentElementsInACircularArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3423: Maximum Difference Between Adjacent Elements in a Circular Array
// https://leetcode.com/problems/maximum-difference-between-adjacent-elements-in-a-circular-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenAdjacentElementsInACircularArray([]int{1, 2, 4}))
	fmt.Println(MaximumDifferenceBetweenAdjacentElementsInACircularArray([]int{-5, -1, -3}))
}

// MaximumDifferenceBetweenAdjacentElementsInACircularArray returns the max absolute diff between adjacent elements (circular).
// Time: O(n). Space: O(1).
func MaximumDifferenceBetweenAdjacentElementsInACircularArray(nums []int) int {
	n := len(nums)
	maxDiff := 0
	for i := 0; i < n; i++ {
		diff := nums[i] - nums[(i+1)%n]
		if diff < 0 {
			diff = -diff
		}
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}
```
