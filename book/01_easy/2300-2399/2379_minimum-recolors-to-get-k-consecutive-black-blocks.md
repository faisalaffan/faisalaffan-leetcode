# 2379 — Minimum Recolors To Get K Consecutive Black Blocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumRecolorsToGetKConsecutiveBlackBlocks(blocks string, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #2379: Minimum Recolors to Get K Consecutive Black Blocks
// https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(MinimumRecolorsToGetKConsecutiveBlackBlocks("WBBWWBBWBW", 7)) // 3
	fmt.Println(MinimumRecolorsToGetKConsecutiveBlackBlocks("WBWBBBW", 2))    // 0
}

func MinimumRecolorsToGetKConsecutiveBlackBlocks(blocks string, k int) int {
	wCount := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			wCount++
		}
	}
	minOps := wCount
	for i := k; i < len(blocks); i++ {
		if blocks[i-k] == 'W' {
			wCount--
		}
		if blocks[i] == 'W' {
			wCount++
		}
		if wCount < minOps {
			minOps = wCount
		}
	}
	return minOps
}
```
