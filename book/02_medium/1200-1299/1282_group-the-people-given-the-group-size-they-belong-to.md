# 1282 — Group The People Given The Group Size They Belong To

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func groupThePeople(groupSizes []int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1282: Group the People Given the Group Size They Belong To
// https://leetcode.com/problems/group-the-people-given-the-group-size-they-belong-to/
// Difficulty: Medium

// Group people by desired group size, then partition each group.

// Time: O(n)
// Space: O(n)

func groupThePeople(groupSizes []int) [][]int {
  // HashMap: O(1) lookup
	groups := make(map[int][]int)
  // Matriks 2D
	result := make([][]int, 0)

	for person, size := range groupSizes {
		groups[size] = append(groups[size], person)
		if len(groups[size]) == size {
			result = append(result, groups[size])
			delete(groups, size)
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	fmt.Printf("%v\n", groupThePeople([]int{2, 1, 3, 3, 3, 2}))
}
```
