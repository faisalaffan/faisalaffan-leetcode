# 2191 — Sort The Jumbled Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func sortJumbled(mapping []int, nums []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n + n * d)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2191: Sort the Jumbled Numbers
// https://leetcode.com/problems/sort-the-jumbled-numbers/
// Difficulty: Medium
// Time: O(n log n + n * d) | Space: O(n)

import (
	"fmt"
	"sort"
)

func sortJumbled(mapping []int, nums []int) []int {
	type pair struct {
		val    int
		mapped int
		idx    int
	}

	pairs := make([]pair, len(nums))
	for i, num := range nums {
		mapped := 0
		if num == 0 {
			mapped = mapping[0]
		} else {
			digits := []int{}
			n := num
			for n > 0 {
				digits = append(digits, n%10)
				n /= 10
			}
			for j := len(digits) - 1; j >= 0; j-- {
				mapped = mapped*10 + mapping[digits[j]]
			}
		}
		pairs[i] = pair{num, mapped, i}
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].mapped != pairs[j].mapped {
			return pairs[i].mapped < pairs[j].mapped
		}
		return pairs[i].idx < pairs[j].idx
	})

  // Alokasi slice integer
	result := make([]int, len(nums))
	for i, p := range pairs {
		result[i] = p.val
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(sortJumbled([]int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}))
	// Expected: [338, 38, 991]

	// Test case 2
	fmt.Println(sortJumbled([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}))
	// Expected: [123, 456, 789]
}
```
