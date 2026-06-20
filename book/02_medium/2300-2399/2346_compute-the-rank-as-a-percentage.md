# 2346 — Compute The Rank As A Percentage

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func rankAsPercentage(ranks []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2346: Compute the Rank as a Percentage
// https://leetcode.com/problems/compute-the-rank-as-a-percentage/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func rankAsPercentage(ranks []int) []int {
	n := len(ranks)
	// This is a database-style problem. Simplified:
	// For each student, compute (rank-1)*100/(n-1)
	// But since it's "compute the rank as a percentage" database problem:
	// The percentage = (R-1)*100 / (N-1)
  // Alokasi slice
	result := make([]int, n)
	for i, r := range ranks {
		result[i] = (r - 1) * 100 / (n - 1)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(rankAsPercentage([]int{1, 2, 3, 4}))
	// Expected: [0, 33, 66, 100]

	// Test case 2
	fmt.Println(rankAsPercentage([]int{1, 2}))
	// Expected: [0, 100]
}
```
