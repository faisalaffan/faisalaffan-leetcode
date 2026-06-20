# 0877 — Stone Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func StoneGame(piles []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #877: Stone Game
// https://leetcode.com/problems/stone-game/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(StoneGame([]int{5, 3, 4, 5}))
	fmt.Println(StoneGame([]int{3, 7, 2, 5}))
	fmt.Println(StoneGame([]int{1, 100, 3, 2}))
}

// Time: O(1) | Space: O(1)
// Alex always wins because there are an even number of piles
// and total stones is odd (no ties), with Alex going first.
func StoneGame(piles []int) bool {
	return true
}
```
