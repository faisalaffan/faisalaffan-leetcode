# 1046 — Last Stone Weight

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func lastStoneWeight(stones []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1046: Last Stone Weight
// https://leetcode.com/problems/last-stone-weight/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1})) // 1
	fmt.Println(lastStoneWeight([]int{1}))                 // 1
	fmt.Println(lastStoneWeight([]int{2, 2}))              // 0
}

// LeetCode submission: lastStoneWeight
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
  // Sort O(n log n)
		sort.Ints(stones)
		n := len(stones)
		if stones[n-1] == stones[n-2] {
			stones = stones[:n-2]
		} else {
			stones[n-2] = stones[n-1] - stones[n-2]
			stones = stones[:n-1]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
```
