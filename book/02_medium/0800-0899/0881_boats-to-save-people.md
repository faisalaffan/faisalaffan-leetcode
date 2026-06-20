# 0881 — Boats To Save People

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BoatsToSavePeople(people []int, limit int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #881: Boats to Save People
// https://leetcode.com/problems/boats-to-save-people/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BoatsToSavePeople([]int{1, 2}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 2, 2, 1}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 5, 3, 4}, 5))
}

// Time: O(n log n) | Space: O(log n)
func BoatsToSavePeople(people []int, limit int) int {
  // Sort O(n log n)
	sort.Ints(people)
	left, right := 0, len(people)-1
	ans := 0

  // Binary search loop
	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		ans++
	}

	return ans
}
```
