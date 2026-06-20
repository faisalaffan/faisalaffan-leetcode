# 1053 — Previous Permutation With One Swap

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func prevPermOpt1(arr []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1053: Previous Permutation With One Swap
// https://leetcode.com/problems/previous-permutation-with-one-swap/
// Difficulty: Medium
//
// Approach: Find rightmost pair where arr[i] > arr[i+1].
//           Swap with the largest smaller element to the right.
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{3, 2, 1}))    // [3,1,2]
	fmt.Println(prevPermOpt1([]int{1, 1, 5}))    // [1,1,5]
	fmt.Println(prevPermOpt1([]int{1, 9, 4, 6, 7})) // [1,7,4,6,9]
}

func prevPermOpt1(arr []int) []int {
	i := len(arr) - 2
	for i >= 0 && arr[i] <= arr[i+1] {
		i--
	}

	if i < 0 {
		return arr
	}

	// Find rightmost smaller than arr[i]
	j := len(arr) - 1
	for arr[j] >= arr[i] {
		j--
	}
	// Skip duplicates
	for arr[j] == arr[j-1] {
		j--
	}

	arr[i], arr[j] = arr[j], arr[i]
	return arr
}
```
