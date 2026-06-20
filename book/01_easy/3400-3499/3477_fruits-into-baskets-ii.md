# 3477 — Fruits Into Baskets Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FruitsIntoBasketsIi(fruits []int, baskets []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * m). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3477: Fruits Into Baskets II
// https://leetcode.com/problems/fruits-into-baskets-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FruitsIntoBasketsIi([]int{4, 2, 5}, []int{3, 5, 4}))
	fmt.Println(FruitsIntoBasketsIi([]int{3, 6, 1}, []int{6, 4, 7}))
}

// FruitsIntoBasketsIi counts fruits that cannot be placed into baskets.
// Each fruit i can go into basket j if fruits[i] <= baskets[j].
// Time: O(n * m). Space: O(1).
func FruitsIntoBasketsIi(fruits []int, baskets []int) int {
	used := make([]bool, len(baskets))
	unplaced := 0
	for _, f := range fruits {
		placed := false
		for j, b := range baskets {
			if !used[j] && f <= b {
				used[j] = true
				placed = true
				break
			}
		}
		if !placed {
			unplaced++
		}
	}
	return unplaced
}
```
