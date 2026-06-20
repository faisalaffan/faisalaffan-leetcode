# 0781 — Rabbits In Forest

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func numRabbits(answers []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #781: Rabbits in Forest
// https://leetcode.com/problems/rabbits-in-forest/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
}

func numRabbits(answers []int) int {
  // HashMap: O(1) lookup
	count := make(map[int]int)
	for _, a := range answers {
		count[a]++
	}

	result := 0
	for k, v := range count {
		groupSize := k + 1
		groups := (v + groupSize - 1) / groupSize
		result += groups * groupSize
	}

	return result
}
```
