# 0658 — Find K Closest Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findClosestElements(arr []int, k int, x int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n + k)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #658: Find K Closest Elements
// https://leetcode.com/problems/find-k-closest-elements/
// Difficulty: Medium
// Time: O(log n + k)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{1, 1, 1, 10, 10, 10}, 1, 9))
}

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

  // Two-pointer loop
	for left < right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}
```
