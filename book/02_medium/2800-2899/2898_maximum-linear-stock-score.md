# 2898 — Maximum Linear Stock Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumLinearStockScore(prices []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2898: Maximum Linear Stock Score
// https://leetcode.com/problems/maximum-linear-stock-score/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumLinearStockScore(prices []int) int64 {
	// For each stock, score = sum of prices where prices[i] - i is same
  // HashMap: O(1) lookup
	score := make(map[int]int64)
	var best int64

	for i, p := range prices {
		key := p - i
		score[key] += int64(p)
		if score[key] > best {
			best = score[key]
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumLinearStockScore([]int{1, 2, 3, 4}))
	fmt.Println(MaximumLinearStockScore([]int{2, 1, 3}))
}
```
