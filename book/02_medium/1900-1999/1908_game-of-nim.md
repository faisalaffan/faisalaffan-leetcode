# 1908 — Game Of Nim

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func NimGame(piles []int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1908: Game of Nim
// https://leetcode.com/problems/game-of-nim/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(NimGame([]int{1, 2, 3}))
	fmt.Println(NimGame([]int{1, 1, 1}))
	fmt.Println(NimGame([]int{1, 2}))
}

// Time: O(n), Space: O(1)
func NimGame(piles []int) bool {
	xor := 0
	for _, p := range piles {
		xor ^= p
	}
	return xor != 0
}
```
