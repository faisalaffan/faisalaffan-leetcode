# 3159 — Find Occurrences Of An Element In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func occurrencesOfElement(nums []int, queries []int, x int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n + q)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3159: Find Occurrences of an Element in an Array
// https://leetcode.com/problems/find-occurrences-of-an-element-in-an-array/
// Difficulty: Medium
// Time: O(n + q) | Space: O(n)

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
  // Alokasi slice
	pos := make([]int, 0)
	for i, v := range nums {
		if v == x {
			pos = append(pos, i)
		}
	}

  // Alokasi slice
	ans := make([]int, len(queries))
	for i, q := range queries {
		if q-1 < len(pos) {
			ans[i] = pos[q-1]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	fmt.Println(occurrencesOfElement([]int{1, 3, 1, 7}, []int{1, 3, 2, 4}, 1)) // Expected: [0, -1, 2, -1]
	fmt.Println(occurrencesOfElement([]int{1, 2, 3}, []int{10}, 5))             // Expected: [-1]
}
```
