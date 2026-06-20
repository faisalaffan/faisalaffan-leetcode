# 2418 — Sort The People

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func SortThePeople(names []string, heights []int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Alokasi slice integer
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
  // Custom sort dengan comparator
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
