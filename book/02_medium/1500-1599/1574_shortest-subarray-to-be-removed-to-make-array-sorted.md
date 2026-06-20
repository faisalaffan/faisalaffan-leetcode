# 1574 — Shortest Subarray To Be Removed To Make Array Sorted

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindLengthOfShortestSubarray(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Prefix Sum

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1574: Shortest Subarray to be Removed to Make Array Sorted
// https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FindLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
	fmt.Println(FindLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))
	fmt.Println(FindLengthOfShortestSubarray([]int{1, 2, 3}))
}

func FindLengthOfShortestSubarray(arr []int) int {
	// Time: O(N), Space: O(1)
	n := len(arr)

	// Find longest non-decreasing prefix
	left := 0
	for left < n-1 && arr[left] <= arr[left+1] {
		left++
	}

	if left == n-1 {
		return 0 // already sorted
	}

	// Find longest non-decreasing suffix
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}

	// Remove everything between left and right (minimum so far)
	result := minInt(n-left-1, right)

	// Try to merge prefix and suffix
	i, j := 0, right
	for i <= left && j < n {
		if arr[i] <= arr[j] {
			result = minInt(result, j-i-1)
			i++
		} else {
			j++
		}
	}

	return result
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
