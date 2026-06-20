# 0575 — Distribute Candies

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func DistributeCandies(candyType []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #575: Distribute Candies
// https://leetcode.com/problems/distribute-candies/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(n)
func DistributeCandies(candyType []int) int {
  // HashMap: O(1) lookup
	types := make(map[int]bool)
	for _, c := range candyType {
		types[c] = true
	}
	maxAllowed := len(candyType) / 2
	if len(types) < maxAllowed {
		return len(types)
	}
	return maxAllowed
}

func main() {
	fmt.Println(DistributeCandies([]int{1, 1, 2, 2, 3, 3}))
	fmt.Println(DistributeCandies([]int{1, 1, 2, 3}))
	fmt.Println(DistributeCandies([]int{6, 6, 6, 6}))
}
```
