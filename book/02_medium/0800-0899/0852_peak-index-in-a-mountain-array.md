# 0852 — Peak Index In A Mountain Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func PeakIndexInAMountainArray(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #852: Peak Index in a Mountain Array
// https://leetcode.com/problems/peak-index-in-a-mountain-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PeakIndexInAMountainArray([]int{0, 1, 0}))
	fmt.Println(PeakIndexInAMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(PeakIndexInAMountainArray([]int{0, 10, 5, 2}))
}

// Time: O(log n) | Space: O(1)
func PeakIndexInAMountainArray(arr []int) int {
	left, right := 1, len(arr)-2
  // Two-pointer loop
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```
