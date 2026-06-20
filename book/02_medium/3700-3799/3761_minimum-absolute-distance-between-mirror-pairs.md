# 3761 — Minimum Absolute Distance Between Mirror Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumAbsoluteDistanceBetweenMirrorPairs(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3761: Minimum Absolute Distance Between Mirror Pairs
// https://leetcode.com/problems/minimum-absolute-distance-between-mirror-pairs/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumAbsoluteDistanceBetweenMirrorPairs(nums []int) int {
  // HashMap: O(1) lookup
	prev := make(map[int]int)
	ans := -1

	for j, v := range nums {
		if pos, ok := prev[v]; ok {
			dist := j - pos
			if ans == -1 || dist < ans {
				ans = dist
			}
		}
		// Store reversed number
		rev := 0
		for x := v; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		prev[rev] = j
	}

	return ans
}

func main() {
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{12, 21, 45, 33, 54}))
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{1, 2, 3, 4}))
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{11, 22, 11}))
}
```
