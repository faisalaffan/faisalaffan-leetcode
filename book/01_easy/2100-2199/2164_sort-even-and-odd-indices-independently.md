# 2164 — Sort Even And Odd Indices Independently

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func SortEvenAndOddIndicesIndependently(nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2164: Sort Even and Odd Indices Independently
// https://leetcode.com/problems/sort-even-and-odd-indices-independently/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortEvenAndOddIndicesIndependently([]int{4, 1, 2, 3})) // [2 3 4 1]
	fmt.Println(SortEvenAndOddIndicesIndependently([]int{2, 1}))       // [2 1]
}

// Time: O(n log n), Space: O(n)
func SortEvenAndOddIndicesIndependently(nums []int) []int {
	n := len(nums)
  // Alokasi slice integer
	even := make([]int, 0, (n+1)/2)
  // Alokasi slice integer
	odd := make([]int, 0, n/2)

	for i, v := range nums {
		if i%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

  // Urutkan secara ascending — O(n log n)
	sort.Ints(even)
	sort.Sort(sort.Reverse(sort.IntSlice(odd)))

  // Alokasi slice integer
	result := make([]int, n)
	ei, oi := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = even[ei]
			ei++
		} else {
			result[i] = odd[oi]
			oi++
		}
	}
	return result
}
```
