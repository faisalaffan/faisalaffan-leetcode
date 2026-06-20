# 3206 — Alternating Groups I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func AlternatingGroupsI(colors []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3206: Alternating Groups I
// https://leetcode.com/problems/alternating-groups-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AlternatingGroupsI([]int{1, 1, 1}))
	fmt.Println(AlternatingGroupsI([]int{0, 1, 0, 0, 1}))
}

// AlternatingGroupsI counts the number of groups of 3 adjacent elements where all three are alternating.
// Time: O(n). Space: O(1).
func AlternatingGroupsI(colors []int) int {
	n := len(colors)
	count := 0
	for i := 0; i < n; i++ {
		if colors[i] == colors[(i+2)%n] && colors[i] != colors[(i+1)%n] {
			count++
		}
	}
	return count
}
```
