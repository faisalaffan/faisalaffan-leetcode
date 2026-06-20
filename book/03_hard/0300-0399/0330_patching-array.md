# 0330 — Patching Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minPatches(nums []int, n int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #330: Patching Array
// https://leetcode.com/problems/patching-array/
// Difficulty: Hard

import "fmt"

func minPatches(nums []int, n int) int {
	patches := 0
	miss := int64(1) // smallest sum we cannot form
	i := 0

	for miss <= int64(n) {
		if i < len(nums) && int64(nums[i]) <= miss {
			miss += int64(nums[i])
			i++
		} else {
			// Patch with miss itself
			miss += miss
			patches++
		}
	}
	return patches
}

func main() {
	// Example 1
	fmt.Println(minPatches([]int{1, 3}, 6))
	// 1

	// Example 2
	fmt.Println(minPatches([]int{1, 5, 10}, 20))
	// 2

	// Example 3
	fmt.Println(minPatches([]int{1, 2, 2}, 5))
	// 0

	// Example 4
	fmt.Println(minPatches([]int{1, 2, 31, 33}, 2147483647))
	// 28
}
```
