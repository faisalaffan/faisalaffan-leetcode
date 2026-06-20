# 2285 — Maximum Total Importance Of Roads

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumImportance(n int, roads [][]int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2285: Maximum Total Importance of Roads
// https://leetcode.com/problems/maximum-total-importance-of-roads/
// Difficulty: Medium
// Time: O(n + m) | Space: O(n)

import (
	"fmt"
	"sort"
)

func maximumImportance(n int, roads [][]int) int64 {
  // Alokasi slice integer
	degree := make([]int, n)
	for _, r := range roads {
		degree[r[0]]++
		degree[r[1]]++
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(degree)

	var total int64 = 0
	for i, d := range degree {
		total += int64(i+1) * int64(d)
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println(maximumImportance(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}))
	// Expected: 43

	// Test case 2
	fmt.Println(maximumImportance(5, [][]int{{0, 3}, {2, 4}, {1, 3}}))
	// Expected: 20

	// Test case 3
	fmt.Println(maximumImportance(2, [][]int{{0, 1}}))
	// Expected: 3
}
```
