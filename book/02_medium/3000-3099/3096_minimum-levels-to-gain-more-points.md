# 3096 — Minimum Levels To Gain More Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumLevels(possible []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3096: Minimum Levels to Gain More Points
// https://leetcode.com/problems/minimum-levels-to-gain-more-points/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumLevels(possible []int) int {
	n := len(possible)
  // Alokasi slice
	score := make([]int, n)
	for i, v := range possible {
		if v == 0 {
			score[i] = -1
		} else {
			score[i] = 1
		}
	}

	total := 0
	for _, v := range score {
		total += v
	}

	prefix := 0
	for i := 0; i < n-1; i++ {
		prefix += score[i]
		if prefix > total-prefix {
			return i + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(minimumLevels([]int{1, 0, 1, 0}))       // Expected: 1
	fmt.Println(minimumLevels([]int{1, 1, 1, 1, 1}))    // Expected: 3
	fmt.Println(minimumLevels([]int{0, 0}))              // Expected: -1
}
```
