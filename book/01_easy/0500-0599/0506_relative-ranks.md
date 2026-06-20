# 0506 — Relative Ranks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func RelativeRanks(score []int) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #506: Relative Ranks
// https://leetcode.com/problems/relative-ranks/
// Difficulty: Easy

import (
	"fmt"
	"sort"
	"strconv"
)

// Time: O(n log n), Space: O(n)
func RelativeRanks(score []int) []string {
  // Alokasi slice
	sorted := make([]int, len(score))
	copy(sorted, score)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
  // HashMap: O(1) lookup
	rank := make(map[int]string)
	for i, s := range sorted {
		switch i {
		case 0:
			rank[s] = "Gold Medal"
		case 1:
			rank[s] = "Silver Medal"
		case 2:
			rank[s] = "Bronze Medal"
		default:
			rank[s] = strconv.Itoa(i + 1)
		}
	}
	result := make([]string, len(score))
	for i, s := range score {
		result[i] = rank[s]
	}
	return result
}

func main() {
	fmt.Println(RelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(RelativeRanks([]int{10, 3, 8, 9, 4}))
}
```
