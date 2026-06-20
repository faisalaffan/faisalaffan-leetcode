# 0898 — Bitwise Ors Of Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BitwiseOrsOfSubarrays(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n * log(max))  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #898: Bitwise ORs of Subarrays
// https://leetcode.com/problems/bitwise-ors-of-subarrays/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(BitwiseOrsOfSubarrays([]int{0}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 1, 2}))
	fmt.Println(BitwiseOrsOfSubarrays([]int{1, 2, 4}))
}

// Time: O(n * log(max)) | Space: O(n)
func BitwiseOrsOfSubarrays(arr []int) int {
  // HashMap: O(1) lookup
	set := make(map[int]bool)

  // Linear scan O(n)
	for i := 0; i < len(arr); i++ {
		set[arr[i]] = true
		for j := i - 1; j >= 0; j-- {
			if arr[i]|arr[j] == arr[j] {
				break
			}
			arr[j] |= arr[i]
			set[arr[j]] = true
		}
	}

	return len(set)
}
```
