# 1502 — Can Make Arithmetic Progression From Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func canMakeArithmeticProgression(arr []int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1502: Can Make Arithmetic Progression From Sequence
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/
// Difficulty: Easy
//
// LeetCode submission: func canMakeArithmeticProgression(arr []int) bool

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{3, 5, 1})) // true
	fmt.Println(CanMakeArithmeticProgressionFromSequence([]int{1, 2, 4})) // false
}

// Time: O(n log n), Space: O(1)
func CanMakeArithmeticProgressionFromSequence(arr []int) bool {
  // Sort O(n log n)
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
```
