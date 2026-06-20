# 3644 — Maximum K To Sort A Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumKToSortAPermutation(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3644: Maximum K to Sort a Permutation
// https://leetcode.com/problems/maximum-k-to-sort-a-permutation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumKToSortAPermutation(nums []int) int {
	ans := -1 // all bits set to 1 (identity for AND)
	for i, x := range nums {
		if i != x {
			ans &= x
		}
	}
	if ans < 0 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(maximumKToSortAPermutation([]int{0, 3, 2, 1}))
	fmt.Println(maximumKToSortAPermutation([]int{3, 2, 1, 0}))
	fmt.Println(maximumKToSortAPermutation([]int{0, 1, 2, 3}))
}
```
