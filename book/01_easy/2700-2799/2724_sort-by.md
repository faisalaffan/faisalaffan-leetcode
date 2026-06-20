# 2724 — Sort By

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func SortBy(arr []int, fn func(int) int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2724: Sort By
// https://leetcode.com/problems/sort-by/
// Difficulty: Easy
// Time: O(n log n) | Space: O(1)
// Note: JavaScript problem, adapted to Go.

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fn := func(n int) int { return n % 2 }
	fmt.Println(SortBy(arr, fn))

	arr2 := []int{1, 2, 3, 4, 5}
	fn2 := func(n int) int { return n }
	fmt.Println(SortBy(arr2, fn2))
}

func SortBy(arr []int, fn func(int) int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return fn(arr[i]) < fn(arr[j])
	})
	return arr
}
```
