# 2831 — Find The Longest Equal Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindTheLongestEqualSubarray(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2831: Find the Longest Equal Subarray
// https://leetcode.com/problems/find-the-longest-equal-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheLongestEqualSubarray(nums []int, k int) int {
  // HashMap: O(1) lookup
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := 0
	for _, indices := range pos {
		left := 0
		for right := 0; right < len(indices); right++ {
			// Elements between indices[left] and indices[right] that need to be removed
			for indices[right]-indices[left]-(right-left) > k {
				left++
			}
			if right-left+1 > best {
				best = right - left + 1
			}
		}
	}

	return best
}

func main() {
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))
}
```
