# 0455 — Assign Cookies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func AssignCookies(g, s []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n + m log m), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #455: Assign Cookies
// https://leetcode.com/problems/assign-cookies/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n + m log m), Space: O(1)
func AssignCookies(g, s []int) int {
  // Sort O(n log n)
	sort.Ints(g)
  // Sort O(n log n)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
		}
		j++
	}
	return i
}

func main() {
	fmt.Println(AssignCookies([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(AssignCookies([]int{1, 2}, []int{1, 2, 3}))
}
```
