# 2418 — Sort The People

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func SortThePeople(names []string, heights []int) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2418: Sort the People
// https://leetcode.com/problems/sort-the-people/
// Difficulty: Easy
// Time O(n log n) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortThePeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})) // ["Mary","Emma","John"]
	fmt.Println(SortThePeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))   // ["Bob","Alice","Bob"]
}

func SortThePeople(names []string, heights []int) []string {
	n := len(names)
  // Alokasi slice
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
  // Custom sort
	sort.Slice(idx, func(i, j int) bool {
		return heights[idx[i]] > heights[idx[j]]
	})
	res := make([]string, n)
	for i, id := range idx {
		res[i] = names[id]
	}
	return res
}
```
