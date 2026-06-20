# 0846 — Hand Of Straights

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func HandOfStraights(hand []int, groupSize int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #846: Hand of Straights
// https://leetcode.com/problems/hand-of-straights/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(HandOfStraights([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3))
	fmt.Println(HandOfStraights([]int{1, 2, 3, 4, 5}, 4))
	fmt.Println(HandOfStraights([]int{2, 1}, 2))
}

// Time: O(n log n) | Space: O(n)
func HandOfStraights(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

  // HashMap: O(1) lookup
	count := make(map[int]int)
	for _, card := range hand {
		count[card]++
	}

  // Alokasi slice
	unique := make([]int, 0, len(count))
	for card := range count {
		unique = append(unique, card)
	}
  // Sort O(n log n)
	sort.Ints(unique)

	for _, card := range unique {
		if count[card] > 0 {
			freq := count[card]
			for i := 0; i < groupSize; i++ {
				if count[card+i] < freq {
					return false
				}
				count[card+i] -= freq
			}
		}
	}

	return true
}
```
